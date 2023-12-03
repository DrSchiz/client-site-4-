package pages

import "ieais.com/m/models"

//структура главной страницы
type MainPage struct {
	Title                    string
	Equipments               []models.Equipment
	TypeEquipments           []models.Type_Equipment
	EquipmentStatuses        []models.Equipment_Status
	Equipment_Code_Exception string
	Equipment_Size_Exception string
}
