package main

import (
	"net/http"
	"os"

	"ieais.com/m/handlers"
)

// запуск приложения и указание маршрутов для работы с ним
func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/authorization", handlers.AuthorizationHandler)
	http.HandleFunc("/registration", handlers.RegistrationHandler)
	http.HandleFunc("/main", handlers.MainHandler)
	http.HandleFunc("/main/delete_equipment", handlers.DeleteEquipmentHandler)
	http.HandleFunc("/main/create_keeping_request", handlers.CreateKeepingRequestHandler)
	http.HandleFunc("/equipment/edit", handlers.EditEquipmentHandler)
	http.HandleFunc("/user", handlers.AccountHandler)
	http.HandleFunc("/user/edit", handlers.AccountEditHandler)
	http.HandleFunc("/user/password/edit", handlers.AccountPasswordEditHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(os.Getenv("IP")+":8090", nil)
}
