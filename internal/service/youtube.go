package service

import (
	"context"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/yaml.v3"
)

type Youtube struct {
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
	err = yaml.Unmarshal(file, &apiKkey)
	if err != nil {
		log.Info("cant unmarshal key")
		return credentialAPI{}
	}
	return apiKkey
}

func New(pathString string) *Youtube {

	if pathString == "" {
		return &Youtube{
			isAvailable: false,
		}
	}

	apiKey := readSecretFile(pathString)
	if apiKey.APIKey == "" {
		return &Youtube{
			isAvailable: false,
		}
	}

	youtubeService, err := youtube.NewService(
		context.Background(),
		option.WithAPIKey(apiKey.APIKey),
	)
	if err != nil {
		return &Youtube{
			isAvailable: false,
		}
	}

	return &Youtube{
		client:      youtubeService,
		isAvailable: true,
	}

}

func (y *Youtube) GetThumbnailURL(id string) string {
	if !y.isAvailable || id == "" {
		return ""
	}
	log.Debug("Getting video list")
	call := y.client.Videos.List([]string{"snippet"}).Id(id)

	log.Debug(fmt.Sprintf("Call do with id %v", id))
	responce, err := call.Do()
	if err != nil {
		log.Error(err)
		return ""
	}

	return responce.Items[0].Snippet.Thumbnails.Standard.Url
}
