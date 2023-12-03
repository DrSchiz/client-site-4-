package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"ieais.com/m/http/_token"
	"ieais.com/m/http/client"
	"ieais.com/m/models"
)

// функция авторизации
func Authorization(login string, password string) (bool, string) {
	authData := client.RequestLoginClient{}
	authData.Client_Login = login
	authData.Client_Password = password

	data, _ := json.Marshal(authData)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://"+os.Getenv("IP_API")+":8080/login", "application/json", r)
	if err != nil {
		fmt.Println("No response from request")
		return false, ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	tokenStr := _token.RequestToken{}

	json.Unmarshal(body, &tokenStr)
	if tokenStr.Token_string == "" {
		return false, ""
	}

	return true, tokenStr.Token_string
}

// функция получения авторизированного пользователя
func Validation(cookie *http.Cookie) (models.Client, error) {
	url := "http://" + os.Getenv("IP_API") + ":8080/main/validate"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return models.Client{}, err
	}

	req.Header.Add("Cookie", "Authorization="+cookie.Value)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.Client{}, err
	}

	var _client models.Client
	json.NewDecoder(res.Body).Decode(&_client)

	defer res.Body.Close()

	return _client, nil
}

// функция регистрации
func Registration(client client.RequestCreateClient) bool {
	data, err := json.Marshal(client)
	if err != nil {
		fmt.Println(err)
		return false
	}
	r := bytes.NewReader(data)
	resp, err := http.Post("http://"+os.Getenv("IP_API")+":8080/create-account", "application/json", r)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.Body)
		return false
	}
	return true
}

// функция редактирования данных клиента
func EditAccount(cookie *http.Cookie, _client client.RequestUpdateClient) {
	url := "http://" + os.Getenv("IP_API") + ":8080/main/client"
	method := "PUT"

	jsonValue, err := json.Marshal(_client)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", "Authorization="+cookie.Value)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
}

// функция смены пароля
func ChangePassword(cookie *http.Cookie, _client client.RequestChangePassword) {
	url := "http://" + os.Getenv("IP_API") + ":8080/main/client/password"
	method := "PUT"

	jsonValue, err := json.Marshal(_client)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Cookie", "Authorization="+cookie.Value)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

}
