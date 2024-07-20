package service

import (
	"context"
	"os"
	"youtube-tumbnail-grpc/pkg/repo"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/yaml.v3"
)

type Youtube struct {
	cacheRepo   repo.CacheRepo
	client      *youtube.Service
	isAvailable bool
}

type credentialAPI struct {
	APIKey string `yaml:"api_key"`
}

func readSecretFile(path string) credentialAPI {

	file, err := os.ReadFile(path)
	if err != nil {
		log.Info("cant read secret file")
		return credentialAPI{}
	}

	apiKkey := credentialAPI{}
	err = yaml.Unmarshal(file, apiKkey)
	if err != nil {
		log.Info("cant unmarshal key")
		return credentialAPI{}
	}
	return apiKkey
}

func New(pathString string, cacheRepo repo.CacheRepo) *Youtube {

	if pathString == "" {
		return &Youtube{
			isAvailable: false,
			cacheRepo:   cacheRepo,
		}
	}

	apiKey := readSecretFile(pathString)
	if apiKey.APIKey != "" {
		return &Youtube{
			isAvailable: false,
			cacheRepo:   cacheRepo,
		}
	}

	youtubeService, err := youtube.NewService(
		context.Background(),
		option.WithAPIKey(apiKey.APIKey),
	)
	if err != nil {
		return &Youtube{
			isAvailable: false,
			cacheRepo:   cacheRepo,
		}
	}

	return &Youtube{
		client:      youtubeService,
		isAvailable: true,
		cacheRepo:   cacheRepo,
	}

}

func (y *Youtube) GetInfo(id string) string {

}
