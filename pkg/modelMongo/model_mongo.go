package modelMongo

import (
	"context"
	"github.com/Sumitk99/crud_api_dbconfig.git/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type BsonEntry struct {
	Date  string `json:"date" bson:"date"`
	Day   string `json:"day" bson:"day"`
	Tasks string `json:"tasks" bson:"tasks"`
}

var Client *mongo.Client
var Ctx context.Context
var Cancel context.CancelFunc

func init() {
	config.ConnectMongo("mongodb://localhost:27017")
	Client, Ctx, Cancel = config.GetMongoSession()
}

func ReadMongo(db, col string, query, field interface{}) []BsonEntry {

	collection := Client.Database(db).Collection(col)
	var filter, option interface{}
	filter = bson.D{{}}
	option = bson.D{{"_id", 0}}
	cursor, err := collection.Find(Ctx, filter, options.Find().SetProjection(option))
	if err != nil {
		panic(err)
	}
	var result []bson.D
	defer cursor.Close(Ctx)
	cursor.All(Ctx, &result)
	var output []BsonEntry
	for _, entry := range result {
		a := entry[0].Value.(string)
		b := entry[1].Value.(string)
		c := entry[2].Value.(string)
		output = append(output, BsonEntry{Date: a, Day: b, Tasks: c})
	}
	return output
}

func (entry BsonEntry) WriteMongo(db, col string) {

	new_entry := bson.D{
		{"date", entry.Date},
		{"day", entry.Day},
		{"tasks", entry.Tasks},
	}
	collection := Client.Database(db).Collection(col)
	_, err := collection.InsertOne(Ctx, new_entry)
	if err != nil {
		log.Fatal(err)
	}
}

func (entry BsonEntry) DeleteMongo(db, col string) {

	new_entry := bson.D{
		{"date", entry.Date},
	}
	collection := Client.Database(db).Collection(col)
	_, err := collection.DeleteOne(Ctx, new_entry)
	if err != nil {
		log.Fatal(err)
	}
}

func (entry BsonEntry) UpdateMongo(db, col string) {

	update := bson.D{
		{"$set", bson.D{{"day", entry.Day}, {"tasks", entry.Tasks}}},
	}
	filter := bson.D{{"date", entry.Date}}

	collection := Client.Database(db).Collection(col)

	_, err := collection.UpdateOne(Ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
}
