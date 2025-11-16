package producer

import (
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
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare("alerts", false, false, false, false, nil)
	if err != nil {
		conn.Close()
		ch.Close()
		return nil, err
	}

	return &Producer{
		log:  log,
		conn: conn,
		ch:   ch,
		q:    q.Name,
	}, nil
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
