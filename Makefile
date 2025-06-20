# --- Variables ---
SERVICE_NAME := domain-driven-golang

# PostgreSQL Dependency (for local 'make run')
DB_CONTAINER_NAME := dev-postgres
DB_IMAGE          := postgres:17-alpine
DB_PORT           := 5432
DB_VOLUME         := dev-postgres-data
DB_NAME           := tasks_db
DB_USER           := postgres
DB_PASSWORD       := 12345

# Environment file
ENV_FILE          := .env

# --- Cross-Platform Helpers ---
# Set the correct 'remove directory' command based on the operating system
ifeq ($(OS),Windows_NT)
	RM = rmdir /s /q
else
	RM = rm -rf
endif

# ====================================================================================
# HELPERS
# ====================================================================================

.PHONY: help
## help: Show this help message.
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "--------------------------"
	@echo "  Development & Building "
	@echo "--------------------------"
	@echo "  run           Run the Go application locally."
	@echo "  tidy          Ensure Go modules are tidy."
	@echo "  build         Build the Go application."
	@echo "  test          Run all Go tests."
	@echo "  clean-build   Remove build artifacts."
	@echo ""
	@echo "--------------------------"
	@echo "  Local Dependencies     "
	@echo "--------------------------"
	@echo "  postgres      Start the local PostgreSQL container for development."
	@echo "  postgres-stop Stop and remove the local PostgreSQL container."
	@echo "  postgres-logs Follow the logs of the local PostgreSQL container."
	@echo "  postgres-shell Open a 'psql' shell in the local PostgreSQL container."
	@echo ""
	@echo "--------------------------"
	@echo "  Housekeeping           "
	@echo "--------------------------"
	@echo "  clean         Stop the PostgreSQL container and remove build artifacts."
	@echo "  help          Show this help message."

# ====================================================================================
# DEVELOPMENT & BUILDING (Go)
# ====================================================================================

.PHONY: run build test clean-build

## run: Run the Go application.
run:
	@echo "==> Starting Go service '$(SERVICE_NAME)'..."
	@echo "==> Note: Your service must be configured to connect to localhost:$(DB_PORT)."
	@go run main.go

## tidy: Ensure Go modules are tidy.
tidy:
	@echo "==> Ensuring Go modules are tidy..."
	@go mod tidy
	@echo "==> Go modules are tidy."

## build: Build the Go application.
build:
	@echo "==> Building application '$(SERVICE_NAME)'..."
	@go build -o bin/$(SERVICE_NAME)
	@echo "==> Build complete: bin/$(SERVICE_NAME)"

## test: Run all Go tests.
test:
	@echo "==> Running Go tests..."
	@go test ./...

## clean-build: Remove build artifacts.
clean-build:
	@echo "==> Cleaning up build artifacts..."
	@go clean
	@$(RM) bin

# ====================================================================================
# LOCAL DEPENDENCIES (PostgreSQL for 'make run')
# ====================================================================================

.PHONY: postgres postgres-stop postgres-logs postgres-shell

## postgres: Start the local PostgreSQL container if stopped, or create it if it doesn’t exist.
postgres:
	@echo "==> Ensuring local PostgreSQL container '$(DB_CONTAINER_NAME)' is running..."
	@docker run -d \
		--name $(DB_CONTAINER_NAME) \
		-p $(DB_PORT):5432 \
		-e POSTGRES_DB=$(DB_NAME) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-v $(DB_VOLUME):/var/lib/postgresql/data \
		$(DB_IMAGE)
	@echo "==> Local PostgreSQL is running on localhost:$(DB_PORT)"

## postgres-stop: Stop and remove the local PostgreSQL container.
postgres-stop:
	@echo "==> Stopping and removing local PostgreSQL container '$(DB_CONTAINER_NAME)'..."
	@docker stop $(DB_CONTAINER_NAME)
	@docker rm $(DB_CONTAINER_NAME)
	@echo "==> Container stopped and removed."

## postgres-logs: Tail the logs of the running local PostgreSQL container.
postgres-logs:
	@echo "==> Tailing logs for '$(DB_CONTAINER_NAME)'. Press Ctrl+C to exit."
	@docker logs -f $(DB_CONTAINER_NAME)

## postgres-shell: Open an interactive psql shell to the local PostgreSQL container.
postgres-shell:
	@echo "==> Connecting to PostgreSQL shell in '$(DB_CONTAINER_NAME)'..."
	@docker exec -it $(DB_CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)

# ====================================================================================
# HOUSEKEEPING
# ====================================================================================

.PHONY: clean

## clean: Stop the PostgreSQL container and remove build artifacts.
clean: postgres-stop clean-build
	@echo "==> Full cleanup complete."