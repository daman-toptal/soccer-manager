FROM golang:1.17 AS build-stage

WORKDIR /app

COPY . .

RUN cd cmd/docker/external && \
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o main .

FROM alpine:3.11

LABEL maintainer="daman"
LABEL Description="soccer-manager-external"

RUN apk update && \
    adduser -S -D -H -h /app app_user && \
    apk add strace && \
    apk add curl && \
    apk add tcpdump

COPY --from=build-stage /app/cmd/docker/external/main /app/

RUN mkdir -p /app/config

COPY cmd/docker/external/config/* /app/config/

RUN chown -R app_user: /app/*

WORKDIR /app

CMD ["./main"]
