import { HearManager } from "@puregram/hear";
import { MessageContext, Telegram } from "puregram";
import { messagesHandler } from "./messages";

export function registerHandlers(hm: HearManager<MessageContext>, telegram: Telegram) {
  hm.hear("/start", (ctx) => messagesHandler.start(ctx));
}
