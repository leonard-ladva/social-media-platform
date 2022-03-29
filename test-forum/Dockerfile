# syntax=docker/dockerfile:1

FROM golang:1.16-alpine AS go_build

WORKDIR /app
COPY . ./
RUN go mod download
RUN apk add --update gcc musl-dev
RUN go build -o /forum

FROM alpine:latest
COPY assets /assets
COPY public /public
COPY --from=go_build /forum /forum 
ENV SERVER_PORT=":8001"
ENV SERVER_HTTPS_PORT=":8000"
ENV SQLITE_USERNAME="admin"
ENV SQLITE_PASSWORD="admin321"
ENV SQLITE_DB_PATH="./assets/database.db"
CMD [ "/forum" ]
