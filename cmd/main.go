package main

import (
	"log"
	"youtube_tumbnail/internal/app"
)

const defaultConfigPath = "config/config.yaml"

func main() {
	if err := app.Run(defaultConfigPath); err != nil {
		log.Fatal(err)
	}
}
