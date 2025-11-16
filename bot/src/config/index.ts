import { readFileSync } from "fs";
import { Config } from "../types/config";
import * as toml from "toml";

const raw = readFileSync("config.toml", "utf8");

export const config: Config = toml.parse(raw) as Config;
