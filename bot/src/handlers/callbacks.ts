import { CallbackQueryContext } from "puregram";
import { logCbQuery } from "../utils/logs";
import { db } from "../services/db";

class CallbacksHandler {
  async subsribe(ctx: CallbackQueryContext) {
    logCbQuery("subscribe", ctx);

    if (await db.checkIfExists(ctx.from.id.toString())) {
      await ctx.message?.send("You already subscribed");
      return;
    }

    await db.addUser(ctx.from.id.toString());

    await ctx.message?.send("Success");
    await ctx.answer();
  }
}

export const callbacksHandler = new CallbacksHandler();
