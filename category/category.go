package category

type Category struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
}
