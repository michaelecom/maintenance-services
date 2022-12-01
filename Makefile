ifeq ($(OS), Windows_NT)
SHELL := powershell.exe
.SHELLFLAGS := -Command
endif

run:
	go run cmd/main.go

build:
	$$env:GOOS = "linux"; go build -o build cmd/main.go

docker-build:
	docker build -t go-app:alpine .

build-image: build docker-build

docker-compose-up:
	docker-compose --project-name maintenance-services up -d --build

docker-compose-down:
	docker-compose --project-name maintenance-services down

migrate:
	migrate -path ./schema -database 'postgres://postgres:pass@localhost:5432/postgres?sslmode=disable' up

.PHONY: build

.DEFAULT_GOAL := run
