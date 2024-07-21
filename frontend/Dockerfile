FROM node:20.15.1-alpine

WORKDIR /app

COPY package.json package-lock.json ./
RUN yarn install
COPY . .

CMD ["yarn", "run", "dev"]