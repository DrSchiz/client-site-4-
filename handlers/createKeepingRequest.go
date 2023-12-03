package handlers

import (
	"fmt"
	"net/http"

	"ieais.com/m/funcs"
)

// метод обработки создания запроса на хранение оборудования
func CreateKeepingRequestHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie!")
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		return
	}

	funcs.CreateKeepingRequest(cookie, r.FormValue("createRequest"))
	http.Redirect(w, r, "/main", http.StatusMovedPermanently)
}
