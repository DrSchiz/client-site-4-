package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"ieais.com/m/funcs"
	"ieais.com/m/http/equipment"
	"ieais.com/m/models"
	"ieais.com/m/pages"
)

// функция обработки главной страницы
func MainHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./assets/html/main.html", "./assets/html/header.html", "./assets/html/footer.html")
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie!")
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		return
	}
	log.Println(cookie)

	var client models.Client

	client, err = funcs.Validation(cookie)
	if err != nil {
		fmt.Println("Ошибка получения клиента!")
		return
	}

	var equipments []models.Equipment
	equipments = funcs.GetEquipments(cookie)

	var type_equipments []models.Type_Equipment
	type_equipments = funcs.GetTypeEquipments()

	var equipment_statuses []models.Equipment_Status
	equipment_statuses = funcs.GetEquipmentStatus()

	p := pages.MainPage{
		Title:             client.Client_Login,
		Equipments:        equipments,
		TypeEquipments:    type_equipments,
		EquipmentStatuses: equipment_statuses,
	}

	if r.Method == "POST" {

		var newEquipment equipment.RequestCreateEquipment

		es, _ := strconv.Atoi(r.FormValue("equipment_size"))

		newEquipment = equipment.RequestCreateEquipment{
			Equipment_Code: r.FormValue("number"),
			Type_Name:      r.FormValue("type_name"),
			Equipment_Size: es,
			Status_Name:    r.FormValue("status_name"),
		}

		if r.FormValue("number") == "" {
			p.Equipment_Code_Exception = "Поле Номер оборудования не должно быть пустым!"
		} else {
			p.Equipment_Code_Exception = ""
		}

		for _, i := range equipments {
			if i.Equipment_Code == r.FormValue("number") {
				p.Equipment_Code_Exception = "Данный номер оборудования уже используется!"
			}
		}

		if es == 0 {
			p.Equipment_Size_Exception = "Поле Размер оборудования не должно быть пустым!"
		} else {
			p.Equipment_Size_Exception = ""
		}

		if p.Equipment_Code_Exception == p.Equipment_Size_Exception {
			funcs.CreateEquipment(cookie, newEquipment)
		}
		http.Redirect(w, r, "/main", http.StatusMovedPermanently)
	}

	t.Execute(w, p)
}
