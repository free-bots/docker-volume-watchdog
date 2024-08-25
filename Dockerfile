FROM golang:1.19 AS builder

WORKDIR /go/src/github.com/free-bots/docker-volume-watchdog/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -tags netgo -a -v -o ./bin/watch-dog main.go

FROM alpine:3.20.2

ARG DOCKER_VOLUME_WATCHDOG_DISCORD_WEBHOOK
ARG DOCKER_VOLUME_WATCHDOG_INTERVAL_VALUE

ENV DOCKER_VOLUME_WATCHDOG_DISCORD_WEBHOOK $DOCKER_VOLUME_WATCHDOG_DISCORD_WEBHOOK
ENV DOCKER_VOLUME_WATCHDOG_INTERVAL_VALUE $DOCKER_VOLUME_WATCHDOG_INTERVAL_VALUE

RUN mkdir /watch-dog-mount # creating mount point

WORKDIR watch-dog
COPY --from=builder /go/src/github.com/free-bots/docker-volume-watchdog/bin/ ./

CMD ./watch-dog
