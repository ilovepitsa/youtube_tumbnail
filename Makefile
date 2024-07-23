include .env

.PHONY: compose-up compose-down restart test
compose-up:
	docker compose up --build -d

compose-down:
	docker compose down --remove-orphans

restart:
	docker compose down --remove-orphans
	docker compose up --build -d

test:
	go test -v ./...