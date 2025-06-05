package api

import (
	"net/http"

	"myproject/api/handlers"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.CreateUserHandler(w, r)
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
