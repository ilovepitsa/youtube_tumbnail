package service

type (
	MediaSource interface {
		GetThumbnailURL(id string) string
	}

	ThumbnailsService interface {
		GetThumbnail(id string) ([]byte, error)
	}
)
