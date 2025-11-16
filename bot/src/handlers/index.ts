import { HearManager } from "@puregram/hear";
import { MessageContext, Telegram } from "puregram";
import { messagesHandler } from "./messages";
import { callbacksHandler } from "./callbacks";

export function registerHandlers(hm: HearManager<MessageContext>, telegram: Telegram) {
  hm.hear("/start", (ctx) => messagesHandler.start(ctx));

  telegram.updates.on("callback_query", (ctx) => callbacksHandler.subsribe(ctx));

  messagesHandler.alert(telegram);
}
