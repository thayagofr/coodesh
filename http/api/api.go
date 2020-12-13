package api

import (
	"encoding/json"
	"github.com/thyagofr/coodesh/desafio/service"
	"github.com/thyagofr/coodesh/desafio/utils"
	"net/http"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now().UTC()
}

// Home - Home
func Home(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse(w, http.StatusOK)
	response := service.GetInfo()
	response.ElapseTime = time.Since(startTime).String()
	_ = json.NewEncoder(w).Encode(&response)
}

// UpdateProduct - Handler to update a product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	code, err := utils.GetCode(r, "code")
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err.Error())
	}
	var prod utils.UpdateProductRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&prod); err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = service.UpdateProduct(code, prod)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(w, http.StatusAccepted)
}

// RemoveProduct - Handler to "remove" a product
func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	code, err := utils.GetCode(r, "code")
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = service.RemoveProduct(code)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(w, http.StatusAccepted)
}

// GetProduct - Handler to find a product
func GetProduct(w http.ResponseWriter, r *http.Request) {
	code, err := utils.GetCode(r, "code")
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	response, err := service.GetProduct(code)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(response)
	utils.SuccessResponse(w, 0)
}

// GetProducts - Handler to return all products
func GetProducts(w http.ResponseWriter, r *http.Request) {

	page, size := utils.GetQueryParams(r)
	response, err := service.GetProducts(page, size)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(response)
	utils.SuccessResponse(w, 0)
}
