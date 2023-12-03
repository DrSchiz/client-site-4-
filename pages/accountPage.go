package pages

import "ieais.com/m/models"

//структура страницы личного кабинета
type AccountPage struct {
	ClientModel               models.Client
	Login_Exception           string
	Firstname_Exception       string
	Name_Exception            string
	Email_Exception           string
	Password_Exception        string
	Repeat_Password_Exception string
}
