import { InlineKeyboard, MessageContext } from "puregram";
import { logCommand } from "../utils/logs";

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
}

export const messagesHandler = new MessagesHandler();
