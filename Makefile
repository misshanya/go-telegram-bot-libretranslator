BINARY_NAME=bot
BIN_DIR=./bin
CMD_DIR=./cmd/bot
DB_NAME=bot.db

.PHONY: all build clean run

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(CMD_DIR)

clean:
	@echo "Cleaning..."
	rm -f $(BIN_DIR)/$(BINARY_NAME)

run: build
	@echo "Running $(BINARY_NAME)..."
	$(BIN_DIR)/$(BINARY_NAME)

migrate-up:
	@echo "Applying migrations..."
	goose -dir ./sql/migrations sqlite3 ${DB_NAME} up

migrate-down:
	@echo "Reverting migrations..."
	goose -dir ./sql/migrations sqlite3 ${DB_NAME} down