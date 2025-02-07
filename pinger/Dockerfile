FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

ENV GOOS=linux

WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY ./pinger/cmd ./pinger/cmd
COPY ./pinger/internal ./pinger/internal

RUN go build -ldflags="-s -w" -o /app/service ./pinger/cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/service /app/service

CMD ["./service"]