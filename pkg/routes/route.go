package routes

import (
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/controllers"
	"github.com/gorilla/mux"
)

var Operation_routes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Read_entries).Methods("GET")
	//router.HandleFunc("/", controllers.Add_entry).Methods("POST")
	//router.HandleFunc("/", controllers.Update_entry).Methods("PUT")
	//router.HandleFunc("/", controllers.Delete_entry).Methods("DELETE")
}
