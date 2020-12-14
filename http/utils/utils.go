package utils

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/thyagofr/coodesh/desafio/model"
	"net/http"
	"strconv"
)

type UpdateProductRequest struct {
	ProductName     string `json:"product_name"`
	Quantity        string `json:"quantity"`
	Brands          string `json:"brands"`
	IngredientsText string `json:"igredients_text"`
	NutriscoreScore int    `json:"nutriscore_score"`
	NutriscoreGrade string `json:"nutriscore_grade"`
	MainCategory    string `json:"main_category"`
}

type PaginateResponse struct {
	TotalElements int64
	TotalPages    int64
	ActualPage    int64
	ActualSize    int64
	Content          []model.Product
}

// GetCode - Funcao para pegar o parametro da URI
func GetCode(r *http.Request, param string) (string, error) {

	vars := mux.Vars(r)
	paramValue := vars[param]
	if paramValue == "" {
		return "", errors.New("Identificador inválido")
	}
	return paramValue, nil
}

// GetQueryParams
func GetQueryParams(r *http.Request) (int64, int64) {

	var page, size int64
	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	if pageStr == "" {
		page = 0
	} else {
		page, _ = strconv.ParseInt(pageStr, 10, 64)
	}
	if sizeStr == "" {
		size = 10
	} else {
		size, _ = strconv.ParseInt(sizeStr, 10, 64)
	}
	return page, size

}

func SuccessResponse(w http.ResponseWriter, status int) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	if status == 0 {
		return w
	}
	w.WriteHeader(status)
	return w
}
