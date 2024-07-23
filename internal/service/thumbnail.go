package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"youtube_tumbnail/pkg/repo"
)

type Thumbnails struct {
	cacheRepo repo.CacheRepo
	source    MediaSource
}

func NewThumbnails(cacheRepo repo.CacheRepo, source MediaSource) *Thumbnails {
	return &Thumbnails{
		cacheRepo: cacheRepo,
		source:    source,
	}
}

func (t *Thumbnails) readThumbnail(url string) ([]byte, error) {
	buf := new(bytes.Buffer)
	responce, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer responce.Body.Close()

	if _, err := io.Copy(buf, responce.Body); err != nil {
		return nil, fmt.Errorf("cant read image: %w", err)
	}
	return buf.Bytes(), nil
}

func (t *Thumbnails) GetThumbnail(id string) ([]byte, error) {
	res, err := t.cacheRepo.GetCache(id)

	if err == nil {
		return res, nil
	}
	if err != repo.ErrNotFound {
		return nil, fmt.Errorf("get cache err: %w", err)
	}

	url := t.source.GetThumbnailURL(id)

	body, err := t.readThumbnail(url)
	if err != nil {
		return nil, err
	}
	go t.cacheRepo.Cache(id, body)

	return body, nil
}
