package controllers

import (
	"encoding/json"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/models"
	"net/http"
)

func Read_entries(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	Output := models.ReadDatabse()
	json.NewEncoder(res).Encode(Output)
}

func Add_entry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var New_entry models.Entry
	json.NewDecoder(req.Body).Decode(&New_entry)
	New_entry.AddToDatabase()
	json.NewEncoder(res).Encode(New_entry)
}

func Update_entry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var New_entry models.Entry
	json.NewDecoder(req.Body).Decode(&New_entry)
	New_entry.UpdateDatabase()
	json.NewEncoder(res).Encode(models.ReadDatabse())
}

func Delete_entry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var New_entry models.Entry
	json.NewDecoder(req.Body).Decode(&New_entry)
	New_entry.RemoveFromDatabase()
	json.NewEncoder(res).Encode(models.ReadDatabse())
}
