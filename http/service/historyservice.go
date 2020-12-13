package service

import (
	"context"
	"github.com/thyagofr/coodesh/desafio/model"
	"github.com/thyagofr/coodesh/desafio/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type HService struct {
	Client *mongo.Client
}

func (h *HService) LastMigration() string {
	collection := h.Client.Database("banco").Collection(utils.GetCollection(utils.HISTORY))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.FindOneOptions{}
	opt.SetSort(bson.D{{"runningt" , -1 }})
	var hist model.History
	err := collection.FindOne(ctx, bson.D{} , &opt).Decode(&hist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "Migração ainda não realizada"
		}
	}
	return hist.RunningT.String()
}


func (h *HService) Ping() string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := h.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return "Database OFF"
	}
	return "Database ON"
}

