package handlers

import (
	"fmt"
	"net/http"

	"ieais.com/m/funcs"
	"ieais.com/m/http/client"
)

// метод обработки изменения данных аккаунта
func AccountEditHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie!")
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		return
	}

	var _client client.RequestUpdateClient
	_client = client.RequestUpdateClient{
		Client_Login:      r.FormValue("login"),
		Client_Name:       r.FormValue("name"),
		Client_Firstname:  r.FormValue("lastname"),
		Client_Patronymic: r.FormValue("patronymic"),
		Client_Email:      r.FormValue("EMail"),
	}

	fmt.Println(_client)

	funcs.EditAccount(cookie, _client)
	http.Redirect(w, r, "/user", http.StatusMovedPermanently)
}
