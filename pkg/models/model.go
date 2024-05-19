package models

import (
	"database/sql"
	"fmt"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/config"
)

var db_sql *sql.DB

type Entry struct {
	Date  string `json:"date"`
	Day   string `json:"day"`
	Tasks string `json:"tasks"`
}

func init() {
	config.ConnectSQL()
	db_sql = config.GetSQL()
}

func (E *Entry) AddToDatabase() {
	temp, err := db_sql.Prepare("INSERT INTO entry_table VALUES(?,?,?)")
	defer temp.Close()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", E.Tasks)
	temp.Exec(E.Date, E.Day, E.Tasks)
	temp.Close()
}

func ReadDatabse() []Entry {
	var Ouput []Entry
	result, err := db_sql.Query("SELECT * FROM entry_table")
	defer result.Close()
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var a, b, c string
		result.Scan(&a, &b, &c)
		fmt.Printf("%s %s %s\n", a, b, c)
		Ouput = append(Ouput, Entry{a, b, c})
	}
	return Ouput
}

func (e *Entry) RemoveFromDatabase() {
	temp, err := db_sql.Prepare("DELETE FROM entry_table WHERE date=?")
	defer temp.Close()
	if err != nil {
		panic(err.Error())
	}
	temp.Exec(e.Date)
}

func (e *Entry) UpdateDatabase() {
	temp, err := db_sql.Prepare("UPDATE entry_table SET task=?,day=? WHERE date=?")

	defer temp.Close()
	if err != nil {
		panic(err.Error())
	}
	temp.Exec(e.Tasks, e.Day, e.Date)
}
