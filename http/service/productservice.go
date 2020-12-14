package service

import (
	"context"
	"errors"
	"github.com/thyagofr/coodesh/desafio/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/thyagofr/coodesh/desafio/model"
	"go.mongodb.org/mongo-driver/bson"
)

type PService struct {
	Client *mongo.Client
}

// GetProducts - Get all products
func (p *PService) GetProducts(page, size int64) (*utils.PaginateResponse, error) {
	collection := p.Client.Database("banco").Collection(utils.GetCollection(utils.PRODUCTS))
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
	number, _ := collection.CountDocuments(ctx, bson.D{})
	response := utils.PaginateResponse{
		ActualPage:    page,
		ActualSize:    size,
		Content:       products,
		TotalElements: number,
		TotalPages:    number / size,
	}
	return &response, nil
}

// GetProduct - Find a product by code
func (p *PService) GetProduct(code string) (*model.Product, error) {
	collection := p.Client.Database("banco").Collection(utils.GetCollection(utils.PRODUCTS))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product := model.Product{}
	err := collection.FindOne(ctx, bson.M{"code": code}).Decode(&product)
	return &product, err
}

// UpdateProduct - Update data of a product by code
func (p *PService) UpdateProduct(code string, request utils.UpdateProductRequest) error {
	collection := p.Client.Database("banco").Collection(utils.GetCollection(utils.PRODUCTS))
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
	result, err := collection.UpdateOne(ctx, bson.M{"code": code}, update)
	if err != nil {
		return nil
	}
	if result.MatchedCount == 0 {
		return errors.New("Product not found.")
	}
	return err
}

// RemoveProduct - Remove a product by code
func (p *PService) RemoveProduct(code string) error {
	collection := p.Client.Database("banco").Collection(utils.GetCollection(utils.PRODUCTS))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	update := bson.D{{"$set", bson.D{
		{"status", utils.GetStatus(utils.TRASH)},
	}}}
	res, err := collection.UpdateOne(ctx, bson.M{"code": code}, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("Product not found,")
	}
	return nil

}

//
//func GetInfo() utils.Info {
//	info := utils.Info{
//		Name:             "Food API",
//		Author:           "Thyago Freitas da Silva",
//		LastExecutedTime: LastMigration(),
//		Connection:       database.Ping(),
//	}
//	return info
//}
