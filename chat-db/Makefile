include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

gen-sql:
	$(LOCAL_BIN)/sqlc generate

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-add:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create $(name) sql

migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v