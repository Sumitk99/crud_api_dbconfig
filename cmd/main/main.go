package main

import (
	"fmt"
	_ "github.com/Sumitk99/crud_api_dbconfig.git/pkg/modelMongo"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/models"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	routes.Operation_routes(router)
	TestEntry := models.Entry{}
	TestEntry.AddToDatabase()

	http.Handle("/", router)
	fmt.Println("Starting server on port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
