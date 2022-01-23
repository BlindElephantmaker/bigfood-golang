POSTGRES_DB = bigfood
POSTGRES_USER = bigfood
POSTGRES_PASSWORD = bigfood

migrate-create: # usage: make migration-create name=migration_name
	migrate create -ext sql -dir ./migrations -seq $(name)

migrate-up:
	migrate -path ./migrations -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5431/$(POSTGRES_DB)?sslmode=disable' up

migrate-down: # todo: pass parameter or 1 if not passed
	migrate -path ./migrations -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5431/$(POSTGRES_DB)?sslmode=disable' down 1

server-run:
	go run cmd/app/main.go

swag-init: # todo: warning: failed to get package name in dir: ./, error: execute go list command, exit status 1, stdout:, stderr:no Go files in
	swag init --generalInfo ./cmd/app/main.go --output ./swagger

swag-fmt:
	swag fmt --generalInfo ./cmd/app/main.go

swag:
	make swag-fmt && make swag-init
