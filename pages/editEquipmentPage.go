package pages

import "ieais.com/m/models"

//структура страницы изменения данных оборудования
type EditEquipmentPage struct {
	Title                    string
	Equip                    models.Equipment
	TypeEquipments           []models.Type_Equipment
	EquipmentStatuses        []models.Equipment_Status
	Equipment_Code_Exception string
	Equipment_Size_Exception string
}
