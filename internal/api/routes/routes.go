package routes

import (
	"net/http"

	"github.com/igormbonfim/shorturl-api/internal/api/services"
)

func RegisterRoutes() {
	http.HandleFunc("/api/url", corsMiddleware(handlePost))
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

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := []string{"https://www.goshort.tech", "http://localhost:3000"}
		origin := r.Header.Get("Origin")

		for _, o := range allowedOrigins {
			if origin == o {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
