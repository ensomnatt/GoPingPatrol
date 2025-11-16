// src/db.ts
import { Pool } from "pg";
import { logger } from "../../utils/logger";

class DB {
  private pool: Pool;

  constructor() {
    this.pool = new Pool({
      host: "bot-db",
      port: 5432,
      user: "postgres",
      password: "postgres",
      database: "postgres",
    });

    this.init();
  }

  private async init() {
    await this.pool.query(`
      CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        tgid TEXT NOT NULL UNIQUE
      )
    `);
    logger.info("Initialized DB");
  }

  async addUser(tgid: string): Promise<void> {
    try {
      await this.pool.query(
        "INSERT INTO users (tgid) VALUES ($1) ON CONFLICT (tgid) DO NOTHING",
        [tgid]
      );
      logger.info({ tgid }, "Added user");
    } catch (err) {
      logger.error(err, "Failed to add user");
    }
  }

  async getAllUsers(): Promise<string[]> {
    const res = await this.pool.query<{ tgid: string }>(
      "SELECT tgid FROM users"
    );
    logger.info("Got all users");
    return res.rows.map(row => row.tgid);
  }

  async checkIfExists(tgid: string): Promise<boolean> {
    const res = await this.pool.query<{ tgid: string }>(
      "SELECT 1 FROM users WHERE tgid = $1 LIMIT 1",
      [tgid]
    );
    return typeof res.rowCount === "number" && res.rowCount > 0;
  }
}

export const db = new DB();
