import { InlineKeyboard, MessageContext, Telegram } from "puregram";
import { logCommand } from "../utils/logs";
import { Consumer } from "../services/consumer";
import { db } from "../services/db";
import { logger } from "../utils/logger";

class MessagesHandler {
  async start(ctx: MessageContext) {
    logCommand("start", ctx);

    await ctx.send("Do you wanna subsribe to alerts?", {
      reply_markup: InlineKeyboard.keyboard([
        InlineKeyboard.textButton({
          text: "Subscribe",
          payload: "SUBSCRIBE"
        })
      ])
    });
  }

  async alert(telegram: Telegram) {
    const consumer = new Consumer("amqp://rabbitmq:5672", "alerts");
    await consumer.connect();

    const tgids = await db.getAllUsers();

    await consumer.consume(async (url: string) => {
      for (const tgid of tgids) {
        telegram.api.sendMessage({
          chat_id: tgid,
          text: `${url} is down!`
        });
      }
      logger.info("Sent alert");
    });
  }
}

export const messagesHandler = new MessagesHandler();
