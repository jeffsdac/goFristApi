package handler

import (
	"fmt"
	"net/http"
)

func testPing(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		fmt.Fprintf(writer, "Ping realizado com sucesso! method GET")
	}
	fmt.Fprint(writer, "Realize seu ping com o method GET!")
}
