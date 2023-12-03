package funcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"ieais.com/m/http/equipment"
	"ieais.com/m/models"
)

// функция получения оборудования по его номеру
func GetEquipment(cookie *http.Cookie, equipment_code string) models.Equipment {
	url := "http://" + os.Getenv("IP_API") + ":8080/main/equipment/" + equipment_code
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return models.Equipment{}
	}

	req.Header.Add("Cookie", "Authorization="+cookie.Value)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return models.Equipment{}
	}

	var equipment models.Equipment

	json.NewDecoder(res.Body).Decode(&equipment)

	defer res.Body.Close()

	return equipment
}

// функция получения всех оборудований авторизированного клиента
func GetEquipments(cookie *http.Cookie) []models.Equipment {

	url := "http://" + os.Getenv("IP_API") + ":8080/main/equipment"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("Cookie", "Authorization="+cookie.Value)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var equipments []models.Equipment

	json.NewDecoder(res.Body).Decode(&equipments)

	defer res.Body.Close()

	return equipments
}

// функция получения видов оборудования
func GetTypeEquipments() []models.Type_Equipment {
	resp, err := http.Get("http://" + os.Getenv("IP_API") + ":8080/type_equipments")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var type_equipments []models.Type_Equipment

	json.NewDecoder(resp.Body).Decode(&type_equipments)

	defer resp.Body.Close()

	return type_equipments
}

// функция получения статусов целостности оборудования
func GetEquipmentStatus() []models.Equipment_Status {
	resp, err := http.Get("http://" + os.Getenv("IP_API") + ":8080/equipment_statuses")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var equipment_statuses []models.Equipment_Status

	json.NewDecoder(resp.Body).Decode(&equipment_statuses)

	defer resp.Body.Close()

	return equipment_statuses
}

// функция создания оборудования
func CreateEquipment(cookie *http.Cookie, newEquipment equipment.RequestCreateEquipment) bool {
	apiurl := "http://" + os.Getenv("IP_API") + ":8080/main/equipment/create-equipment"
	method := "POST"

	jsonValue, err := json.Marshal(newEquipment)
	if err != nil {
		fmt.Println(err)
		return false
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, apiurl, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return false
	}

	req.Header.Add("Cookie", "Authorization="+cookie.Value)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer res.Body.Close()

	return true
}

// функция удаления оборудования
func DeleteEquipment(cookie *http.Cookie, equipmentCode string) {
	url := "http://" + os.Getenv("IP_API") + ":8080/main/equipment/" + equipmentCode + "/delete"
	method := "DELETE"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
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

	var equipments []models.Equipment

	json.NewDecoder(res.Body).Decode(&equipments)

	defer res.Body.Close()
}

// функция создания заявки на хранение оборудования
func CreateKeepingRequest(cookie *http.Cookie, equipmentCode string) {
	apiurl := "http://" + os.Getenv("IP_API") + ":8080/main/equipment/" + equipmentCode + "/create-request"
	method := "POST"

	client := &http.Client{}

	req, err := http.NewRequest(method, apiurl, nil)
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

// функция изменения данных оборудования
func EditEquipment(cookie *http.Cookie, _equipment equipment.RequestCreateEquipment, equipment_code string) {
	apiurl := "http://" + os.Getenv("IP_API") + ":8080/main/equipment/" + equipment_code + "/edit"
	method := "PUT"

	fmt.Println(apiurl)
	fmt.Println(_equipment)

	jsonValue, err := json.Marshal(_equipment)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, apiurl, bytes.NewBuffer(jsonValue))
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
