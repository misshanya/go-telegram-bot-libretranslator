FROM golang:1.23.6 AS builder

WORKDIR /app

COPY . .

# Download requirements
RUN go mod download

# Build
RUN go build -o ./bot ./cmd/bot

# Debian as runner
FROM debian:stable-slim AS runner

WORKDIR /app
COPY --from=builder /app/sql/migrations/ ./migrations
COPY --from=builder /app/bot .

# Install libsqlite3-dev and wget
RUN apt-get update && apt-get install -y libsqlite3-dev wget

# Install goose for db migrations
RUN wget -O /usr/local/bin/goose https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_x86_64
RUN chmod +x /usr/local/bin/goose

# Migrate db and run bot
CMD ["sh", "-c", "goose -dir ./migrations sqlite3 $DB_PATH up &&./bot"]