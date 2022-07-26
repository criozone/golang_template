FROM golang:1.18.4-alpine3.15 as chunk1

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go install github.com/githubnemo/CompileDaemon@latest

FROM chunk1 as build

WORKDIR /app
COPY . /app
COPY ./docker/go/wss-entry-point.sh /usr/local/bin/

ARG APP_WSS_PORT=8081
EXPOSE $APP_WSS_PORT

CMD [ "wss-entry-point.sh" ]
