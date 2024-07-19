RUN_GO=docker compose exec go
DB_CONFIG=-config=internal/dataacess/database/migrations/dbconfig.yml

.PHONY: wire
wire:
	$(RUN_GO) wire internal/wiring/wire.go

.PHONY: start
start:
	$(RUN_GO) go run cmd/*.go

.PHONY: migrate-new
migrate-new:
	@read -p  "What is the name of file?" NAME; \
	${RUN_GO} sql-migrate new ${DB_CONFIG} $$NAME

.PHONY: migrate-up
migrate-up:
	$(RUN_GO) sql-migrate up ${DB_CONFIG}

.PHONY: migrate-down
migrate-down:
	@read -p  "Enter your step num you want to migrate down: " NUM; \
	$(RUN_GO) sql-migrate down ${DB_CONFIG} -limit=$$NUM

