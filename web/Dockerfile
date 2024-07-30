FROM node:20.15.1-alpine AS builder

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install

COPY . .

FROM node:20.15.1-alpine AS runner

WORKDIR /app

COPY --from=builder /app .

CMD ["yarn", "run", "dev"]