package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"ieais.com/m/funcs"
	"ieais.com/m/http/equipment"
	"ieais.com/m/models"
	"ieais.com/m/pages"
)

// метод обработки страницы изменения данных оборудования
func EditEquipmentHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./assets/html/editEquipment.html", "./assets/html/header.html", "./assets/html/footer.html")

	cookie, _ := r.Cookie("Authorization")

	var _equipment models.Equipment
	_equipment = funcs.GetEquipment(cookie, r.FormValue("editEquip"))

	var type_equipments []models.Type_Equipment
	type_equipments = funcs.GetTypeEquipments()

	var equipment_statuses []models.Equipment_Status
	equipment_statuses = funcs.GetEquipmentStatus()

	fmt.Println(_equipment)

	p := pages.EditEquipmentPage{
		Title:             "Редактирование",
		Equip:             _equipment,
		TypeEquipments:    type_equipments,
		EquipmentStatuses: equipment_statuses,
	}

	if r.Method == "POST" {
		es, _ := strconv.Atoi(r.FormValue("equipment_size"))

		var _equipment equipment.RequestCreateEquipment
		_equipment = equipment.RequestCreateEquipment{
			Equipment_Code: r.FormValue("number"),
			Type_Name:      r.FormValue("type_name"),
			Equipment_Size: es,
			Status_Name:    r.FormValue("status_name"),
		}

		funcs.EditEquipment(cookie, _equipment, _equipment.Equipment_Code)
		http.Redirect(w, r, "/main", http.StatusMovedPermanently)
	}

	t.Execute(w, p)
}
