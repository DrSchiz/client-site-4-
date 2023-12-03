package handlers

import (
	"net/http"
	"strings"
	"text/template"

	"github.com/go-playground/validator/v10"
	"ieais.com/m/funcs"
	"ieais.com/m/http/client"
	"ieais.com/m/pages"
)

// функция обработки страницы регистрации
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	p := pages.RegistrationPage{}
	p.Title = "Регистрация"

	t, _ := template.ParseFiles("./assets/html/registration.html")

	if r.Method == "POST" {
		client := client.RequestCreateClient{
			Client_Login:           r.FormValue("login"),
			Client_Password:        r.FormValue("password"),
			Client_Repeat_Password: r.FormValue("password-repeat"),
			Client_Firstname:       r.FormValue("firstname"),
			Client_Name:            r.FormValue("name"),
			Client_Patronymic:      r.FormValue("patronymic"),
			Client_Phone_Number:    r.FormValue("phone-number"),
			Client_Email:           r.FormValue("email"),
		}

		validate := validator.New()
		err := validate.Struct(client)
		if err != nil {
			if strings.Contains(err.Error(), "Client_Login") {
				p.Login_Exception = "Логин должен содержать от 3 до 20 символов!"
			} else {
				p.Login_Exception = ""
			}
			if strings.Contains(err.Error(), "Client_Name") {
				p.Name_Exception = "Имя не должно быть пустым!"
			} else {
				p.Name_Exception = ""
			}
			if strings.Contains(err.Error(), "Client_Firstname") {
				p.Firstname_Exception = "Фамилия не должна быть пустой!"
			} else {
				p.Firstname_Exception = ""
			}
			if strings.Contains(err.Error(), "Client_Phone_Number") {
				p.Phone_Number_Exception = "Неверный формат телефона!"
			} else {
				p.Phone_Number_Exception = ""
			}
			if strings.Contains(err.Error(), "Client_Email") {
				p.Email_Exception = "Неверный формат EMail!"
			} else {
				p.Email_Exception = ""
			}
			if strings.Contains(err.Error(), "Client_Password") {
				p.Password_Exception = "Пароль должен содержать от 3 до 50 символов!"
			} else {
				p.Password_Exception = ""
			}
		}

		if client.Client_Password != client.Client_Repeat_Password {
			p.Repeat_Password_Exception = "Пароли не совпадают!"
		} else {
			p.Repeat_Password_Exception = ""
		}

		createClient := funcs.Registration(client)
		if createClient {
			http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		}
	}

	t.Execute(w, p)
}
