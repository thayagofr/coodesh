package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// InitDatabase - InitDatabase
func InitDatabase() {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	URI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.zzc1h.mongodb.net/%s?retryWrites=true&w=majority",
		os.Getenv("MONGODB_USER"),
		os.Getenv("MONGODB_PASSWORD"),
		os.Getenv("MONGODB_DATABASE"))
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(URI))
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

// GetCollection - Return a collection to the database
func GetCollection(collection string) *mongo.Collection {

	return client.Database("banco").Collection(collection)

}
