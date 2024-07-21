package service

type (
	MediaSource interface {
		GetThumbnailURL(id string) string
	}
)
