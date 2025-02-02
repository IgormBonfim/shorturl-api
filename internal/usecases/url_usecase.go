package usecases

import (
	"crypto/md5"
	"log"
	"math/big"
	"strings"

	"github.com/igormbonfim/shorturl-api/internal/domain/entities"
	"github.com/igormbonfim/shorturl-api/internal/domain/repositories"
)

type UrlUsecase struct {
	repository repositories.UrlRepository
}

func NewUrlUsecase(repository repositories.UrlRepository) *UrlUsecase {
	return &UrlUsecase{
		repository: repository,
	}
}

func (u *UrlUsecase) CreateUrl(url string) (*entities.URL, error) {
	shortcode := generateShortCode(url)
	urlObject, err := u.repository.GetByShortCode(shortcode)
	if err != nil {
		return nil, err
	}

	if urlObject != nil {
		return urlObject, nil
	}

	urlObject = entities.NewURL(url, shortcode)

	go func() {
		id, err := u.repository.Save(urlObject)
		if err != nil {
			log.Println("Erro ao salvar URL:", err)
			return
		}
		urlObject.ID = id
	}()

	return urlObject, nil
}

func (u *UrlUsecase) GetUrlByShortCode(shortcode string) (*entities.URL, error) {
	urlObject, err := u.repository.GetByShortCode(shortcode)
	if err != nil {
		return nil, err
	}

	return urlObject, nil
}

func generateShortCode(url string) string {
	hash := md5.Sum([]byte(url))
	return toBase62(hash[:])
}

func toBase62(hash []byte) string {
	const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	num := new(big.Int).SetBytes(hash)

	var result strings.Builder
	base := big.NewInt(62)

	for num.Sign() > 0 {
		mod := new(big.Int)
		num.DivMod(num, base, mod)
		result.WriteByte(base62Chars[mod.Int64()])
	}

	return result.String()[:6]
}
