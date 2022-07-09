FROM golang:1.17 AS build-stage

WORKDIR /app

COPY . .

RUN echo $(ls .)

RUN cd cmd/docker/grpc && \
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o main .

FROM alpine:3.11

LABEL maintainer="daman"
LABEL Description="soccer-manager-internal"

RUN apk update && \
    adduser -S -D -H -h /app app_user && \
    apk add strace && \
    apk add curl && \
    apk add tcpdump

COPY --from=build-stage /app/cmd/docker/grpc/main /app/

RUN mkdir -p /app/config
COPY ./cmd/docker/grpc/config/* /app/config/

WORKDIR /app

CMD ["./main"]
