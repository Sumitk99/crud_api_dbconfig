package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db_sql *sql.DB

func Connect() {
	d, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbnew")

	if err != nil {
		panic(err)
	}
	err = d.Ping()
	db_sql = d
	//	_, err = db_sql.Exec("CREATE TABLE entry_table(date VARCHAR(20) PRIMARY KEY , day VARCHAR(20), task VARCHAR(100) );")

}

func GetDB() *sql.DB {
	return db_sql
}
