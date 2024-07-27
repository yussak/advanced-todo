FROM golang:1.22

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

CMD ["air", "-c", ".air.toml"]