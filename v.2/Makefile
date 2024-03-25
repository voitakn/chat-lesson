LOCAL_BIN:=$(CURDIR)/bin

DB_HOST:=0.0.0.0
DB_PORT:=5401
DB_USER:=postgres
DB_PASSWORD:=qwerty12345
DB_DATABASE:=chat_lesson

MIGRATION_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_DATABASE} user=${DB_USER} password=${DB_PASSWORD} sslmode=disable"

MIGRATION_DIR=./migrations

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

run:
	DB_DSN="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" \
	go run .

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

# make migration-add name=init_tables
migration-add:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create $(name) sql

migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v