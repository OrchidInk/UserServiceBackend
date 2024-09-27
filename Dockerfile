FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/backend main.go

FROM alpine:3.18

RUN apk add --no-cache bash postgresql-client

WORKDIR /app

COPY --from=builder /app/backend /app/backend

COPY conf /app/conf
COPY db/migration /app/db/migration
COPY utils/secure /app/utils/secure

EXPOSE 8000

CMD ["/app/backend", "conf/conf_development.yml"]
