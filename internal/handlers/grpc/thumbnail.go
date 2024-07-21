package grpc

import (
	"context"
	pb "youtube-tumbnail-grpc/internal/handlers/grpc/thumbnail"
	"youtube-tumbnail-grpc/internal/service"
)

type ThumbnailGRPSServer struct {
	pb.UnimplementedThumbnailsServer

	thService service.ThumbnailsService
}

func NewServer(thService service.ThumbnailsService) *ThumbnailGRPSServer {
	return &ThumbnailGRPSServer{
		thService: thService,
	}
}

func (ts *ThumbnailGRPSServer) GetThumbnail(ctx context.Context, r pb.ThumbnailRequest) (*pb.ThumbnailResponce, error) {

}
