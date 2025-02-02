package services

import (
	"github.com/igormbonfim/shorturl-api/internal/api/controllers"
	"github.com/igormbonfim/shorturl-api/internal/domain/repositories"
	"github.com/igormbonfim/shorturl-api/internal/infra/database"
	dbRepositories "github.com/igormbonfim/shorturl-api/internal/infra/repositories"
	"github.com/igormbonfim/shorturl-api/internal/usecases"
)

var (
	UrlController *controllers.UrlController
	UrlUsecase    *usecases.UrlUsecase
	UrlRepository repositories.UrlRepository
)

func RegisterServices() {
	urlRepo := dbRepositories.NewUrlRepository(database.DB)
	UrlRepository = &urlRepo
	UrlUsecase = usecases.NewUrlUsecase(UrlRepository)
	UrlController = controllers.NewUserController(UrlUsecase)
}
