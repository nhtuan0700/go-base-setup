# RUN_GO=docker compose exec go
DB_CONFIG=-config=internal/dataacess/database/migrations/dbconfig.yml

.PHONY: wire
wire:
	wire internal/wiring/wire.go

.PHONY: start
start:
	go run cmd/*.go start

.PHONY: migrate-new
migrate-new:
	@read -p  "What is the name of file?" NAME; \
	${RUN_GO} sql-migrate new ${DB_CONFIG} $$NAME

.PHONY: migrate-up
migrate-up:
	sql-migrate up ${DB_CONFIG}

.PHONY: migrate-down
migrate-down:
	@read -p  "Enter your step num you want to migrate down: " NUM; \
	sql-migrate down ${DB_CONFIG} -limit=$$NUM

