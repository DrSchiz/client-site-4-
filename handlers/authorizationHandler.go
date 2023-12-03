package handlers

import (
	"log"
	"net/http"
	"text/template"

	"ieais.com/m/funcs"
	"ieais.com/m/pages"
)

// метод обработки страницы авторизации
func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	p := pages.AuthorizationPage{}
	p.Title = "Авторизация"

	_, err := r.Cookie("Authorization")
	if err == nil {
		http.Redirect(w, r, "/main", http.StatusMovedPermanently)
		return
	}
	t, _ := template.ParseFiles("./assets/html/authorization.html")

	if r.Method == "POST" {
		login := r.FormValue("login")
		password := r.FormValue("password")

		checkAuthData, tokenStr := funcs.Authorization(login, password)

		if checkAuthData {
			cookie, err := r.Cookie("Authorization")
			if err != nil {
				cookie = &http.Cookie{
					Name:     "Authorization",
					Value:    tokenStr,
					Path:     "/",
					MaxAge:   3600,
					HttpOnly: true,
					Secure:   false,
					SameSite: http.SameSiteLaxMode,
				}
				log.Println("Создал куки")
				log.Println(cookie)
			}

			http.SetCookie(w, cookie)
			log.Println("Установил куки")
			http.Redirect(w, r, "/main", http.StatusMovedPermanently)
		} else {
			p.Exception = "Неверный логин или пароль!"
		}
	}

	t.Execute(w, p)
}
