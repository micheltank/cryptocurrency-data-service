# Stage 1
FROM golang:1.16.4-alpine3.13 AS builder

ENV GOPATH="$HOME/go"

WORKDIR $GOPATH/src/github.com/micheltank/cryptocurrency-data-service

COPY . $GOPATH/src/github.com/micheltank/cryptocurrency-data-service

RUN apk update && apk add curl git && apk add gcc libc-dev

RUN go mod download
RUN go build -o cryptocurrency-data-service main.go

# Stage 2
FROM alpine:3.13

ENV GOPATH="$HOME/go"

WORKDIR /app

RUN apk update && apk add ca-certificates && apk --no-cache add tzdata

COPY --from=builder $GOPATH/src/github.com/micheltank/cryptocurrency-data-service .

ENTRYPOINT ["./cryptocurrency-data-service"]