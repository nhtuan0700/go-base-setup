RUN_GO=docker compose exec backend
DB_CONFIG=-config=internal/dataacess/database/migrations/dbconfig.yml

.PHONY: wire
wire:
	${RUN_GO} wire internal/wiring/wire.go

.PHONY: start
start:
	${RUN_GO} go run cmd/*.go start

.PHONY: migrate-new
migrate-new:
	@read -p  "What is the name of file?" NAME; \
	${RUN_GO} sql-migrate new ${DB_CONFIG} $$NAME

.PHONY: migrate-up
migrate-up:
	${RUN_GO} sql-migrate up ${DB_CONFIG}

.PHONY: migrate-down
migrate-down:
	@read -p  "Enter your step num you want to migrate down: " NUM; \
	${RUN_GO} sql-migrate down ${DB_CONFIG} -limit=$$NUM

.PHONY: debug
debug:
	${RUN_GO} go build -gcflags "all=-N -l" -buildvcs=false -o ./cmd/main-debug ./cmd/*.go
	${RUN_GO} dlv --listen=:4000 --headless=true --api-version=2 exec ./cmd/main-debug start
	rm ./cmd/main-debug

.PHONY: generate-swagger
generate-swagger:
	swag init -g internal/handler/http/handler.go -o ./docs
