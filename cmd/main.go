package main

import (
	"log"

	"github.com/ilovepitsa/youtube_tumbnail/internal/app"
)

const defaultConfigPath = "config/config.yaml"

func main() {
	if err := app.Run(defaultConfigPath); err != nil {
		log.Fatal(err)
	}
}
