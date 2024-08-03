FROM golang:1.22 AS builder

WORKDIR /app

RUN apt-get update && \
    apt-get install -y \
        git \
        build-essential

# goファイルでimportして使用するものはimportに追加してgo mod tidyでいいが、それ以外のものはDockerfileでインストールする
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

FROM golang:1.22-alpine AS runner

WORKDIR /app

# airを使用するためにlibc6-compatが必要
RUN apk add --no-cache libc6-compat

COPY --from=builder /app .
# airの使用に必要
COPY --from=builder /go/bin/air /usr/local/bin/air

CMD ["air", "-c", ".air.toml"]