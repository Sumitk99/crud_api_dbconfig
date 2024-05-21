package models

import (
	"database/sql"
	"fmt"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/config"
	"log"
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

type Table struct {
	TableName  string `json:"table_name"`
	LocalKey   string `json:"local_key"`
	ForeignKey string `json:"foreign_key"`
}
type TableList struct {
	RootTable  string  `json:"root_table"`
	TableArray []Table `json:"table_array"`
}
type Pair struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func (v *TableList) CreateQuery() string {
	query := "SELECT * FROM " + v.RootTable
	leftTable := v.RootTable
	foreignKey := leftTable + "." + v.TableArray[0].ForeignKey

	for i := 0; i < len(v.TableArray); i++ {
		if i > 0 {
			leftTable = v.TableArray[i-1].TableName
			foreignKey = leftTable + "." + v.TableArray[i-1].ForeignKey
		}
		rightTable := v.TableArray[i].TableName
		localKey := rightTable + "." + v.TableArray[i].LocalKey

		query = query + " INNER JOIN " + rightTable + " ON " + localKey + " = " + foreignKey
		fmt.Println(query)
	}
	return query
}

func (v *TableList) JoinTable() [][]Pair {

	query := v.CreateQuery()
	rows, err := db_sql.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	var results [][]Pair
	for rows.Next() {
		rowValues := make([]string, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range rowValues {
			valuePtrs[i] = &rowValues[i]
		}
		err = rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal(err)
		}
		var rowData []Pair
		for i, col := range columns {
			ele := rowValues[i]
			rowData = append(rowData, Pair{Field: col, Value: ele})
		}
		results = append(results, rowData)
	}
	return results
}
