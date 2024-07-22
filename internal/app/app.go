package app

import (
	"fmt"
	"net"
	"youtube-tumbnail-grpc/config"
	grpcTh "youtube-tumbnail-grpc/internal/handlers/grpc"
	pb "youtube-tumbnail-grpc/internal/handlers/grpc/thumbnail"
	"youtube-tumbnail-grpc/internal/service"
	"youtube-tumbnail-grpc/pkg/repo/redis"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

func Run(configPath string) error {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return err
	}

	SetLogrus(cfg.Log)

	log.Info("Initialize redis.....")

	r, err := redis.New(cfg.R)
	if err != nil {
		return err
	}
	defer r.Close()

	log.Info("Initializing youtube service.....")
	youtubeService := service.New(cfg.YoutubeAPI.YoutubeConfigFilePath)

	log.Info("Initializing thumbnail service.....")
	thumb := service.NewThumbnails(r, youtubeService)

	log.Info("Initializing thumbnail grpc.....")
	grpcThumbnail := grpcTh.NewServerThumb(thumb, cfg.Application.MaxWorkerPoolSize)

	log.Info("Initializing thumbnail grpc.....")
	grpcServer := grpc.NewServer()

	log.Info("Starting listing tcp ", fmt.Sprintf("%v:%v", cfg.Net.Host, cfg.Net.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", cfg.Net.Host, cfg.Net.Port))
	if err != nil {
		return err
	}
	pb.RegisterThumbnailsServer(grpcServer, grpcThumbnail)

	err = grpcServer.Serve(lis)
	return err
}
