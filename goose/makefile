GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING = postgres://postgres:123456@localhost:5433/user-db?sslmode=disable
GOOSE_MIGRATION_DIR ?= ./sql


export GOOSE_DRIVER
export GOOSE_DBSTRING
export GOOSE_MIGRATION_DIR
export GOOSE_TABLE
 
up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
