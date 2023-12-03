package handlers

import (
	"net/http"
	"text/template"

	"ieais.com/m/pages"
)

// метод обработки страницы index
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	p := pages.IndexPage{Title: "IEAIS", Logo: "./assets/img/Logo.svg"}
	t, _ := template.ParseFiles("./assets/html/index.html")

	t.Execute(w, p)
}
