package app

import (
	"youtube-tumbnail-grpc/config"
	"youtube-tumbnail-grpc/internal/service"
	"youtube-tumbnail-grpc/pkg/repo/redis"

	log "github.com/sirupsen/logrus"
)

func Run(configPath string) error {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return err
	}

	SetLogrus(cfg.Log)

	log.Info("Initialize redis")

	r, err := redis.New(cfg.R)
	if err != nil {
		return err
	}
	defer r.Close()

	log.Info("Initializeng youtube service")
	youtubeService := service.New(cfg.YoutubeAPI.YoutubeConfigFilePath)

	log.Info("Getting thumbnail")
	log.Info(youtubeService.GetThumbnailURL("NA7NDdW7Lvw"))
	return nil
}
