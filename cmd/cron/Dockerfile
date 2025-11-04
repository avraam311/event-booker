FROM golang:alpine AS build_base

WORKDIR /cron

COPY . .

RUN go mod download

RUN go build -o cron ./cmd/cron/main.go

FROM alpine AS runner

COPY --from=build_base ./cron .
COPY ./config ./config
COPY ./.env .

CMD ["./cron"]