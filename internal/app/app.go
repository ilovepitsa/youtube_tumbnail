package app

import (
	"youtube-tumbnail-grpc/config"
	"youtube-tumbnail-grpc/pkg/repo/redis"

	log "github.com/sirupsen/logrus"
)

func Run(configPath string) error {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return err
	}

	SetLogrus(cfg.Log)

	r, err := redis.New(cfg.R)
	if err != nil {
		return err
	}

	defer r.Close()

	log.Info("Initialize redis")
	return nil
}
