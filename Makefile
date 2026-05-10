.PHONY: generate dev dev-local

GO ?= go
COMPOSE ?= docker compose

# Regenerate internal/di/wire_gen.go from internal/di/wire.go (Wire).
generate:
	$(GO) generate ./internal/di

# Run the API server locally.
dev:
	$(GO) run ./cmd/api

# Start Postgres (docker compose db), ensure app database exists, then run the API on the host using .env.local.
# Use `make dev-local` (GNU Make cannot use `:` inside target names without escaping).
dev-local:
	set -a; . ./.env.local; set +a; \
	$(COMPOSE) up -d db --wait; \
	$(COMPOSE) exec -T db psql -U "$${POSTGRES_USER}" -d postgres -c "CREATE DATABASE \"$${APP_DB_NAME}\";" >/dev/null 2>&1 || true; \
	ENV_PATH=.env.local $(GO) run ./cmd/api
