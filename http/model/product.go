package model

import (
	"encoding/json"
	"time"
)

type Product struct {
	Id              string      `json:"id"`
	Code            string      `json:"code"`
	Brands          string      `json:"brands"`
	Creator         string      `json:"creator"`
	Categories      string      `json:"categories"`
	IngredientsText string      `json:"ingredients_text"`
	Labels          string      `json:"labels"`
	ProductName     string      `json:"product_name"`
	PurchasePlaces  string      `json:"purchase_places"`
	Quantity        string      `json:"quantity"`
	ServingQuantity json.Number `json:"serving_quantity"`
	ServingSize     string      `json:"serving_size"`
	Stores          string      `json:"stores"`
	Traces          string      `json:"traces"`
	MainCategory    string      `json:"main_category"`
	ImageURL        string      `json:"image_url"`
	Cities          []string    `json:"cities"`
	URL             string      `json:"url"`
	Status          string      `json:"status"`
	ImportedT       time.Time   `json:"imported_t"`
	CreatedT        int         `json:"created_t"`
	LastModifiedT   int         `json:"last_modified_t"`
	NutriscoreScore int         `json:"nutriscore_score"`
	NutriscoreGrade string      `json:"nutriscore_grade"`
}
