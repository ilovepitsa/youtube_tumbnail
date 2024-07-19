include .env

.PHONY: compose-up compose-down test
compose-up:
	docker compose up --build -d

compose-down:
	docker compose down --remove-orphans

test:
	go test -v ./...