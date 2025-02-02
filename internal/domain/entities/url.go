package entities

import "time"

type URL struct {
	ID          int       `json:"-"`
	ShortCode   string    `json:"short_code"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"-"`
}

func NewURL(originalUrl string, shortCode string) *URL {
	return &URL{
		ShortCode:   shortCode,
		OriginalUrl: originalUrl,
		CreatedAt:   time.Now(),
	}
}
