FROM golang:1.18-alpine AS builder
RUN apk update && apk add git && apk add build-base
RUN mkdir /build
ADD ./backend/ /build
WORKDIR /build

RUN go get github.com/gorilla/websocket
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/satori/go.uuid
RUN go get golang.org/x/crypto

RUN go build -o real-time-forum

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/real-time-forum /build/database.db /app/
WORKDIR /app

CMD ["./real-time-forum"]