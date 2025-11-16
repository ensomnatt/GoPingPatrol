import { Telegram } from "puregram";
import { config } from "./config";
import { logger } from "./utils/logger";
import { HearManager } from "@puregram/hear";
import { registerHandlers } from "./handlers";

async function bootstrap() {
  try {
    const telegram = new Telegram({ token: config.token });

    logger.info("Initialized bot");

    const hearManager = new HearManager();
    telegram.updates.on("message", hearManager.middleware);

    registerHandlers(hearManager, telegram);
    logger.info("Registered handlers");

    telegram.updates.startPolling().then(async () => {
      logger.info(`Started polling @${telegram.bot.username}`)
    }).catch(logger.error);
  } catch (err) {
    logger.error(`Error while starting bot: ${err}`);
  }
}

bootstrap()
