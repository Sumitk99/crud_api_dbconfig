package controllers

import (
	"encoding/json"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/config"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/modelMongo"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/models"
	"net/http"
)

func Read_entries(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if config.GetConfig() {
		Output := models.ReadDatabse()
		json.NewEncoder(res).Encode(Output)
	} else {
		var r interface{}
		Output := modelMongo.ReadMongo("MongoNew", "entries", r, r)
		json.NewEncoder(res).Encode(Output)
	}
}

//func Add_entry(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("Content-Type", "application/json")
//	if config.GetConfig() {
//		var New_entry models.Entry
//		json.NewDecoder(req.Body).Decode(&New_entry)
//		New_entry.AddToDatabase()
//		json.NewEncoder(res).Encode(New_entry)
//	} else {
//		var New_entry modelMongo.BsonEntry
//		json.NewDecoder(req.Body).Decode(&New_entry)
//		New_entry.WriteMongo()
//		json.NewEncoder(res).Encode(modelMongo.ReadMongo())
//	}
//}
//
//func Update_entry(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("Content-Type", "application/json")
//	if config.GetConfig() {
//		var New_entry models.Entry
//		json.NewDecoder(req.Body).Decode(&New_entry)
//		New_entry.UpdateDatabase()
//		json.NewEncoder(res).Encode(models.ReadDatabse())
//	} else {
//		var New_entry modelMongo.BsonEntry
//		json.NewDecoder(req.Body).Decode(&New_entry)
//		New_entry.UpdateMongo()
//		json.NewEncoder(res).Encode(modelMongo.ReadMongo())
//	}
//}
//
//func Delete_entry(res http.ResponseWriter, req *http.Request) {
//	res.Header().Set("Content-Type", "application/json")
//	if config.GetConfig() {
//		var New_entry models.Entry
//		json.NewDecoder(req.Body).Decode(&New_entry)
//		New_entry.RemoveFromDatabase()
//		json.NewEncoder(res).Encode(models.ReadDatabse())
//	} else {
//		var New_entry modelMongo.BsonEntry
//		json.NewDecoder(req.Body).Decode(&New_entry)
//		New_entry.DeleteMongo()
//		json.NewEncoder(res).Encode(modelMongo.ReadMongo())
//	}
//}
