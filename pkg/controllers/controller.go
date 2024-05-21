package controllers

import (
	"encoding/json"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/config"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/modelMongo"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/models"
	"net/http"
)

var r interface{}

func Read_entries(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if config.GetConfig() {
		Output := models.ReadDatabse()
		json.NewEncoder(res).Encode(Output)
	} else {

		Output := modelMongo.ReadMongo("MongoNew", "entries", r, r)
		json.NewEncoder(res).Encode(Output)
	}
}

func Add_entry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if config.GetConfig() {
		var New_entry models.Entry
		json.NewDecoder(req.Body).Decode(&New_entry)
		New_entry.AddToDatabase()
		json.NewEncoder(res).Encode(New_entry)
	} else {
		var New_entry modelMongo.BsonEntry
		json.NewDecoder(req.Body).Decode(&New_entry)
		New_entry.WriteMongo("MongoNew", "entries")
		json.NewEncoder(res).Encode(modelMongo.ReadMongo("MongoNew", "entries", r, r))
	}
}

func Update_entry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if config.GetConfig() {
		var New_entry models.Entry
		json.NewDecoder(req.Body).Decode(&New_entry)
		New_entry.UpdateDatabase()
		json.NewEncoder(res).Encode(models.ReadDatabse())
	} else {
		var New_entry modelMongo.BsonEntry
		json.NewDecoder(req.Body).Decode(&New_entry)
		New_entry.UpdateMongo("MongoNew", "entries")
		json.NewEncoder(res).Encode(modelMongo.ReadMongo("MongoNew", "entries", r, r))
	}
}
func Delete_entry(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if config.GetConfig() {
		var New_entry models.Entry
		json.NewDecoder(req.Body).Decode(&New_entry)
		New_entry.RemoveFromDatabase()
		json.NewEncoder(res).Encode(models.ReadDatabse())
	} else {
		var New_entry modelMongo.BsonEntry
		json.NewDecoder(req.Body).Decode(&New_entry)
		New_entry.DeleteMongo("MongoNew", "entries")
		json.NewEncoder(res).Encode(modelMongo.ReadMongo("MongoNew", "entries", r, r))
	}
}

func Join_Query(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if config.GetConfig() {
		var v models.TableList
		json.NewDecoder(req.Body).Decode(&v)
		json.NewEncoder(res).Encode(v.JoinTable())
	}
}
