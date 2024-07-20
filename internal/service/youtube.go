package service

import (
	"youtube-tumbnail-grpc/pkg/repo"

	"google.golang.org/api/youtube/v3"
)

type Youtube struct {
	cacheRepo repo.CacheRepo
	client    youtube.Service
}
