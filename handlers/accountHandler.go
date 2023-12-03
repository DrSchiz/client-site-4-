package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"ieais.com/m/funcs"
	"ieais.com/m/models"
	"ieais.com/m/pages"
)

// метод обработки страницы изменения данных клиента
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./assets/html/account.html", "./assets/html/header.html", "./assets/html/footer.html")
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie!")
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		return
	}

	var client models.Client

	client, err = funcs.Validation(cookie)
	if err != nil {
		fmt.Println("Ошибка получения клиента!")
		return
	}

	p := pages.AccountPage{}

	p.ClientModel = client

	t.Execute(w, p)
}
