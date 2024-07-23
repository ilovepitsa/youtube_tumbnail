package grpc

import (
	"context"
	"io"
	"sync"

	"github.com/ilovepitsa/youtube_tumbnail/internal/service"
	pb "github.com/ilovepitsa/youtube_tumbnail/pkg/handlers/grpc/thumbnail"
	"github.com/ilovepitsa/youtube_tumbnail/pkg/repo"
	log "github.com/sirupsen/logrus"
)

type ThumbnailGRPSServer struct {
	pb.UnimplementedThumbnailsServer
	maxWorkerPool int
	thService     service.ThumbnailsService
}

func NewServerThumb(thService service.ThumbnailsService, maxWorkerPool int) *ThumbnailGRPSServer {
	return &ThumbnailGRPSServer{
		thService:     thService,
		maxWorkerPool: maxWorkerPool,
	}
}

func (th *ThumbnailGRPSServer) worker(in chan string, out chan repo.ThumbnailObject, errCh chan error, wg *sync.WaitGroup, ctx context.Context) {

	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP

		case id := <-in:
			res, err := th.thService.GetThumbnail(id)
			if err != nil {
				return
			}
			out <- repo.ThumbnailObject{Data: res}
		}
	}

}

func (ts *ThumbnailGRPSServer) GetThumbnail(ctx context.Context, r *pb.ThumbnailRequest) (*pb.ThumbnailResponce, error) {

	log.Info("new call")
	thumbnail, err := ts.thService.GetThumbnail(r.Id)
	if err != nil {
		return &pb.ThumbnailResponce{Data: nil}, err
	}

	return &pb.ThumbnailResponce{Data: thumbnail}, err
}

func (ts *ThumbnailGRPSServer) GetThumbnailsAsync(stream pb.Thumbnails_GetThumbnailsAsyncServer) error {

	jobs := make(chan string, ts.maxWorkerPool/2+1)
	result := make(chan repo.ThumbnailObject, ts.maxWorkerPool/2+1)
	errCh := make(chan error, 1)
	wg := &sync.WaitGroup{}
	ctx, finish := context.WithCancel(context.Background())
	wg.Add(ts.maxWorkerPool + 1)
	for i := 0; i < ts.maxWorkerPool; i++ {
		go ts.worker(jobs, result, errCh, wg, ctx)
	}

	waitc := make(chan struct{})
	go func() {
		for data := range result {
			if err := stream.Send(&pb.ThumbnailResponce{Data: data.Data}); err != nil {
				errCh <- err
				return
			}
		}
		close(waitc)
	}()
	stopAll := func() {
		finish()
		if _, ok := <-jobs; ok {
			close(jobs)
		}
		if _, ok := <-errCh; ok {
			close(errCh)
		}
		if _, ok := <-result; ok {
			close(jobs)
		}
		wg.Wait()
	}

	defer stopAll()

	// LOOP_MAIN:
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		select {
		case err := <-errCh:
			return err
		default:
			jobs <- in.Id

		}

	}
	close(jobs)
	close(errCh)
	wg.Wait()
	close(result)
	<-waitc
	log.Error("here end")
	return nil
}
