package repo

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type ThumbnailObject struct {
	Data []byte
}

type CacheRepo interface {
	Cache(key string, data []byte) error
	GetCache(key string) ([]byte, error)
}
