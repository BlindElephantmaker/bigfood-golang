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

swag-init:
	swag init --generalInfo ./cmd/app/main.go --output ./swagger

swag-fmt:
	swag fmt --generalInfo ./cmd/app/main.go

swag:
	make swag-fmt && make swag-init

migrate-up-test:
	migrate -path ./migrations -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@185.246.64.66:32075/$(POSTGRES_DB)?sslmode=disable' up

#build: # todo: build for deploy
#    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bigfood ./cmd/app/
