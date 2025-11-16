package producer

import (
	"time"

	"github.com/ensomnatt/gopingpatrol/checker/internal/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	log  *logger.Logger
	conn *amqp.Connection
	ch   *amqp.Channel
	q    string
}

func New(log *logger.Logger) (*Producer, error) {
	retries := 5
	delay := 2 * time.Second

	var err error
	for i := 0; i < retries; i++ {
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
		if err == nil {
			ch, err := conn.Channel()
			if err == nil {
				q, err := ch.QueueDeclare("alerts", false, false, false, false, nil)
				if err == nil {
					// Успешное подключение
					return &Producer{
						log:  log,
						conn: conn,
						ch:   ch,
						q:    q.Name,
					}, nil
				}
				ch.Close()
			}
			conn.Close()
		}

		log.Infof("Unable to connect to rabbitmq: %v. Next try %d/%d after %s...", err, i+1, retries, delay)
		time.Sleep(delay)
	}

	return nil, err
}

func (p *Producer) Close() {
	if p.ch != nil {
		p.ch.Close()
	}

	if p.conn != nil {
		p.conn.Close()
	}
}

func (p *Producer) Publish(body string) error {
	return p.ch.Publish("", p.q, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(body),
		DeliveryMode: amqp.Persistent,
	})
}
