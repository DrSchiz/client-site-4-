package handlers

import (
	"fmt"
	"net/http"

	"ieais.com/m/funcs"
)

// метод обработки удаления оборудования
func DeleteEquipmentHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("Authorization")
	if err != nil {
		fmt.Println("No cookie!")
		http.Redirect(w, r, "/authorization", http.StatusMovedPermanently)
		return
	}

	funcs.DeleteEquipment(cookie, r.FormValue("deleteEquip"))
	http.Redirect(w, r, "/main", http.StatusMovedPermanently)
}
