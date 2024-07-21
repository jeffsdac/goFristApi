package main

import (
	"encoding/json"
	"net/http"
)

type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int    `json:"idade"`
}

func handleGetUsuario(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		usuario := Usuario{
			ID:    1,
			Nome:  "Jefferson",
			Email: "email@email.com",
			Idade: 15,
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(usuario); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func main() {
	http.HandleFunc("/test", handleGetUsuario)

	port := ":80"

	http.ListenAndServe(port, nil)
}
