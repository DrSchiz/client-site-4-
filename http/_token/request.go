package _token

//структура токена
type RequestToken struct {
	Token_string string `json:"token"`
}

//структура данных токена
type TokenData struct {
	Sub int `json:"sub"`
}
