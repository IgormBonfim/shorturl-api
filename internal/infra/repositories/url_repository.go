package repositories

import (
	"database/sql"
	"fmt"

	"github.com/igormbonfim/shorturl-api/internal/domain/entities"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(connection *sql.DB) UrlRepository {
	return UrlRepository{
		db: connection,
	}
}

func (repo *UrlRepository) Save(url *entities.URL) (int, error) {

	result, err := repo.db.Exec(
		"INSERT INTO urls (short_code, original_url, created_at) VALUES (?, ?, ?)",
		url.ShortCode, url.OriginalUrl, url.CreatedAt,
	)
	if err != nil {
		fmt.Println("Erro ao inserir url no banco de dados:", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao obter o ID inserido:", err)
		return 0, err
	}

	return int(id), nil
}

func (repo *UrlRepository) GetByShortCode(shortCode string) (*entities.URL, error) {
	queryString := "SELECT id, short_code, original_url, created_at FROM urls WHERE short_code = ?"
	query, err := repo.db.Prepare(queryString)
	if err != nil {
		return nil, err
	}

	var url entities.URL

	err = query.QueryRow(shortCode).Scan(
		&url.ID,
		&url.ShortCode,
		&url.OriginalUrl,
		&url.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &url, nil
}
