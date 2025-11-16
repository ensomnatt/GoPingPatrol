FROM node:24-slim AS base

RUN npm install -g pnpm

WORKDIR /app

COPY package.json .

RUN pnpm install

FROM base AS dev

WORKDIR /app

COPY . .

CMD ["npm", "run", "dev"]
