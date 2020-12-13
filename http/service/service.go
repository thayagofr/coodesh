package service

import (
	"context"
	"errors"
	"github.com/thyagofr/coodesh/desafio/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/thyagofr/coodesh/desafio/database"
	"github.com/thyagofr/coodesh/desafio/model"
	"go.mongodb.org/mongo-driver/bson"
)

// GetProducts - Get all products
func GetProducts(page, size int64) ([]model.Product, error) {
	collection := database.GetCollection(utils.GetCollection(utils.PRODUCTS))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.FindOptions{}
	opt.SetLimit(size)
	opt.SetSkip(page * size)
	cursor, err := collection.Find(
		ctx,
		bson.D{},
		&opt,
	)
	if err != nil {
		return nil, err
	}
	var products []model.Product
	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProduct - Find a product by code
func GetProduct(code string) (*model.Product, error) {
	collection := database.GetCollection(utils.GetCollection(utils.PRODUCTS))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product := model.Product{}
	err := collection.FindOne(ctx, bson.M{"code": code}).Decode(&product)
	return &product, err
}

// UpdateProduct - Update data of a product by code
func UpdateProduct(code string, request utils.UpdateProductRequest) error {
	collection := database.GetCollection(utils.GetCollection(utils.PRODUCTS))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"ProductName", request.ProductName},
		{"Quantity", request.Quantity},
		{"Brands", request.Brands},
		{"IngredientsText", request.IngredientsText},
		{"NutriscoreScore", request.NutriscoreScore},
		{"NutriscoreGrade", request.NutriscoreGrade},
		{"MainCategory", request.MainCategory},
	},
	},
	}
	result , err := collection.UpdateOne(ctx, bson.M{"code": code}, update)
	if err != nil {
		return nil
	}
	if result.MatchedCount == 0 {
		return errors.New("Product not found.")
	}
	return err
}

// RemoveProduct - Remove a product by code
func RemoveProduct(code string) error {
	collection := database.GetCollection(utils.GetCollection(utils.PRODUCTS))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.D{{"$set" , bson.D{
		{"status", utils.GetStatus(utils.TRASH)},
	}}}
	res, err := collection.UpdateOne(ctx, bson.M{"code": code} , update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("Product not found")
	}
	return nil
}

func LastMigration() string {
	collection := database.GetCollection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.FindOneOptions{}
	opt.SetSort(bson.D{{"running_t" , -1 }})
	var hist model.History
	err := collection.FindOne(ctx, bson.D{} , &opt).Decode(&hist)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "Migração ainda não realizada"
		}
	}
	return hist.RunningT.String()
}


func GetInfo() utils.Info {
	info := utils.Info{
		Name:             "Food API",
		Author:           "Thyago Freitas da Silva",
		LastExecutedTime: LastMigration(),
		Connection:       database.Ping(),
	}
	return info
}

