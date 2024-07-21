package redis

import (
	"context"
	"fmt"
	"sync"
	"time"
	"youtube-tumbnail-grpc/config"
	"youtube-tumbnail-grpc/pkg/repo"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

const (
	defaultConnAttempts = 5
	defaultConnTimeout  = time.Second
)

type Redis struct {
	mtx          sync.RWMutex
	client       *redis.Client
	cache        *cache.Cache
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

	r.cache = cache.New(&cache.Options{
		Redis: r.client,
	})

	return r, nil
}

func (r *Redis) Cache(key string, data []byte) error {

	r.mtx.Lock()
	defer r.mtx.Unlock()
	obj := &repo.ThumbnailObject{}

	copy(obj.Data, data)
	err := r.cache.Set(&cache.Item{
		Key:   key,
		Value: obj,
		TTL:   time.Hour * 2,
	})

	return err
}

func (r *Redis) GetCache(key string) ([]byte, error) {

	var res repo.ThumbnailObject
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if err := r.cache.Get(context.TODO(), key, &res); err != nil {
		switch err {
		case cache.ErrCacheMiss:
			return nil, repo.ErrNotFound
		default:
			return nil, err
		}
	}

	return res.Data, nil
}

func (r *Redis) Close() {
	if r.client != nil {
		r.client.Close()
	}
}
