package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/igormbonfim/shorturl-api/internal/usecases"
)

type UrlController struct {
	urlUsecase *usecases.UrlUsecase
}

func NewUserController(usecase *usecases.UrlUsecase) *UrlController {
	return &UrlController{
		urlUsecase: usecase,
	}
}

func (u *UrlController) CreateUrl(w http.ResponseWriter, r *http.Request) {
	var request struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.URL == "" {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	urlValid := isValidURL(request.URL)
	if !urlValid {
		response := map[string]string{
			"message": "Invalid URL",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	defer r.Body.Close()

	url, err := u.urlUsecase.CreateUrl(request.URL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(url)
}

func (u *UrlController) GetUrl(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	url, err := u.urlUsecase.GetUrlByShortCode(shortCode)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	if url == nil {
		http.Error(w, "URL not found", http.StatusBadRequest)
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
