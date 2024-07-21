package service

type (
	MediaSource interface {
		GetThumbnailURL(id string) string
	}

	ThumbnailsService interface {
		GetThumbnail(string) ([]byte, error)
	}
)
