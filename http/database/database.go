package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://thyagofr:392035@cluster0.zzc1h.mongodb.net/banco?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

}

func Ping() string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return "Database OFF"
	}
	return "Database ON"
}

//
//func LastMigration() string {
//	collection := GetCollection("history")
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	err := client.Ping(ctx, readpref.Primary())
//	if err != nil {
//		return "Database OFF"
//	}
//	return "Database ON"
//}

// GetCollection - Return a collection to the database
func GetCollection(collection string) *mongo.Collection {

	return client.Database("banco").Collection(collection)

}
