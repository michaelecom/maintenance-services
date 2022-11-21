.PHONY: build

run:
	go run cmd/main.go

build:
	go build -o build cmd/main.go

migrate:
	migrate -path ./schema -database 'postgres://postgres:pass@localhost:5432/postgres?sslmode=disable' up

docker-build:
	docker build -t go-app:alpine .

docker-compose-up:
	docker-compose --project-name maintenance-services up -d --build

docker-compose-down:
	docker-compose --project-name maintenance-services down

.DEFAULT_GOAL := run
