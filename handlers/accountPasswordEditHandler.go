package handlers

import (
	"fmt"
	"net/http"

	"ieais.com/m/funcs"
	"ieais.com/m/http/client"
)

// метод обработки изменения пароля авторизированного клиента
func AccountPasswordEditHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie!")
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		return
	}

	var _client client.RequestChangePassword
	_client = client.RequestChangePassword{
		Client_Old_Password:    r.FormValue("lastpassword"),
		Client_New_Password:    r.FormValue("newpassword"),
		Client_Repeat_Password: r.FormValue("acceptpassword"),
	}

	fmt.Println(_client)

	funcs.ChangePassword(cookie, _client)

	// cookie = &http.Cookie{
	// 	Name:     "Authorization",
	// 	Value:    "",
	// 	Path:     "/",
	// 	MaxAge:   -1,
	// 	HttpOnly: true,
	// }

	// http.SetCookie(w, cookie)
	http.Redirect(w, r, "/user", http.StatusMovedPermanently)
}
