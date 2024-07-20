package redis

import (
	"context"
	"fmt"
	"time"
	"youtube-tumbnail-grpc/config"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

const (
	defaultConnAttempts = 5
	defaultConnTimeout  = time.Second
)

type Redis struct {
	client       *redis.Client
	connAttempts int
	connTimeout  time.Duration
}

func New(config config.Redis) (*Redis, error) {
	r := &Redis{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	r.client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
	var err error
	for r.connAttempts > 0 {

		if err = r.client.Ping(context.Background()).Err(); err == nil {
			break
		}

		log.Info("Attempting to connect redis. Attempts left: ", r.connAttempts)
		time.Sleep(r.connTimeout)
		r.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("redis new %w", err)
	}

	return r, nil
}

func (r *Redis) Cache(key string, data []byte) error {

	return nil
}

func (r *Redis) GetCache(key string) ([]byte, error) {

	return nil, nil
}

func (r *Redis) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
