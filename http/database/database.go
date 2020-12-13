package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// InitDatabase - InitDatabase
func InitDatabase() *mongo.Client {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	URI := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.zzc1h.mongodb.net/%s?retryWrites=true&w=majority",
		os.Getenv("MONGODB_USER"),
		os.Getenv("MONGODB_PASSWORD"),
		os.Getenv("MONGODB_DATABASE"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//
//// GetCollection - Return a collection to the database
//func GetCollection(collection string) *mongo.Collection {
//
//	return client.Database("banco").Collection(collection)
//
//}
