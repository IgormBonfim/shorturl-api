package repositories

import "github.com/igormbonfim/shorturl-api/internal/domain/entities"

type UrlRepository interface {
	Save(url *entities.URL) (int, error)
	GetByShortCode(shortCode string) (*entities.URL, error)
}
