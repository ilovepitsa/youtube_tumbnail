include .env

.PHONY: compose-up compose-down restart test cover-html cover
compose-up:
	docker compose up --build -d

compose-down:
	docker compose down --remove-orphans

restart:
	docker compose down --remove-orphans
	docker compose up --build -d

test:
	go test -v ./...

cover-html: 
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o cover.html
	open cover.html
	rm coverage.out
	
cover: ### run test with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

