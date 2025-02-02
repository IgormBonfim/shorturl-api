package routes

import (
	"net/http"

	"github.com/igormbonfim/shorturl-api/internal/api/services"
)

func RegisterRoutes() {
	http.HandleFunc("/api/url", handlePost)
	http.HandleFunc("/", handleGet)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
	}

	services.UrlController.CreateUrl(w, r)

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
	}

	services.UrlController.GetUrl(w, r)
}
