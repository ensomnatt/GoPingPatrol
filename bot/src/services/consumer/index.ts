import amqp, { Channel, ConsumeMessage } from "amqplib";
import { logger } from "../../utils/logger";

export class Consumer {
  private connection: any;
  private channel!: Channel;

  constructor(private url: string, private queue: string) { }

  async connect(retries = 5, delay = 2000): Promise<void> {
    for (let attempt = 1; attempt <= retries; attempt++) {
      try {
        this.connection = await amqp.connect(this.url);
        this.channel = await this.connection.createChannel();
        await this.channel.assertQueue(this.queue, { durable: false });
        logger.info(`Connected to rabbitmq queue: ${this.queue}`);
        return;
      } catch (err) {
        logger.error(`Failed to connect to rabbitmq (attempt ${attempt}/${retries}): ${err}`);
        if (attempt < retries) {
          const waitTime = delay * attempt;
          logger.info(`Retrying in ${waitTime}ms...`);
          await new Promise(res => setTimeout(res, waitTime));
        } else {
          throw err;
        }
      }
    }
  }

  async consume(callback: (msgContent: string) => Promise<void> | void) {
    if (!this.channel) {
      throw new Error("Channel is not initialized");
    }

    this.channel.consume(this.queue, async (msg: ConsumeMessage | null) => {
      if (msg) {
        try {
          const content = msg.content.toString();
          await callback(content);
          this.channel.ack(msg);
        } catch (err) {
          logger.error(`Failed to process message: ${err}`);
          this.channel.nack(msg, false, true);
        }
      }
    }, { noAck: false });
  }

  async close() {
    await this.channel.close();
    await this.connection.close();
    logger.info("Rabbitmq connection closed");
  }
}
