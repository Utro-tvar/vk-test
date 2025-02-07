FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

WORKDIR /build

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY ./backend/cmd ./cmd
COPY ./backend/internal ./internal

RUN go build -ldflags="-s -w" -o /app/service ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/service /app/service

EXPOSE 80

CMD ["./service"]