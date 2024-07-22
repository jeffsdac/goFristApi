package handler

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandlerLogin(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Não foi possível decodificar o JSON", http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Login recebido com sucesso!"))
}
