# Build stage
FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o incrementor ./cmd/incrementor/main.go && go build -o incrementor-cli ./cmd/incrementor-cli/main.go

# Run stage
FROM alpine:3.18
WORKDIR /app
ENV PG_MIGRATIONS_PATH=file://migrations
COPY ./internal/database/migrations /app/migrations
COPY --from=builder /app/incrementor .
COPY --from=builder /app/incrementor-cli .

EXPOSE 9000
CMD [ "/app/incrementor" ]