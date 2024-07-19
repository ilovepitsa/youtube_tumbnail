package main

import (
	"youtube-tumbnail-grpc/internal/app"
	"log"
)

const defaultConfigPath = "config/config.yaml"

func main() {
	if err := app.Run(defaultConfigPath); err != nil {
		log.Fatal(err)
	}
}