FROM golang:1.25.5-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gentestdata ./scripts/gentestdata

FROM alpine:latest

RUN apk --no-cache add ca-certificates netcat-openbsd bash

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/gentestdata .

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8000

ENTRYPOINT ["./entrypoint.sh"]
