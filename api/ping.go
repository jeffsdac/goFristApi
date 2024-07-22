package handler

import (
	"fmt"
	"net/http"
)

func testPing(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprint(writer, "Realize seu ping com o method GET!")
}
