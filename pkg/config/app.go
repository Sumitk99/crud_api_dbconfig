package config

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var db_sql *sql.DB
var MongoSession *mgo.Session

func GetConfig() bool {
	return false
}
func ConnectSQL() {
	d, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dbnew")

	if err != nil {
		panic(err)
	}
	err = d.Ping()
	db_sql = d
	//	_, err = db_sql.Exec("CREATE TABLE entry_table(date VARCHAR(20) PRIMARY KEY , day VARCHAR(20), task VARCHAR(100) );")
}

func GetSQL() *sql.DB {
	return db_sql
}

func GetMongoSession() (*mongo.Client, context.Context, context.CancelFunc) {
	return Client, Ctx, Cancel
}

var Client *mongo.Client
var Ctx context.Context
var Cancel context.CancelFunc

func ConnectMongo(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	Ctx = ctx
	Cancel = cancel

}

func Ping(client *mongo.Client, ctx context.Context) error {

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("Connection Successful")
	return nil
}

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) error {
	defer cancel()
	err := client.Disconnect(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Connection Closed")
	return nil
}
