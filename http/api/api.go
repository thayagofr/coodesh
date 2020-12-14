package api

import (
	"encoding/json"
	"fmt"
	"github.com/thyagofr/coodesh/desafio/service"
	"github.com/thyagofr/coodesh/desafio/utils"
	"net/http"
	"runtime"
	"time"
)

type Info struct {
	Name             string
	Author           string
	MemoryUsage      string
	ElapseTime       string
	LastExecutedTime string
	Connection       string
	Version          string
}

var startTime time.Time

func init() {
	startTime = time.Now().UTC()
}

type Application struct {
	PService *service.PService
	HService *service.HService
}

// Home - Home
func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	utils.SuccessResponse(w, http.StatusOK)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	response := Info{
		Name:             "API FoodFacts V1",
		Author:           "Thyago Freitas da Silva",
		ElapseTime:       time.Since(startTime).String(),
		LastExecutedTime: app.HService.LastMigration(),
		Connection:       app.HService.Ping(),
		MemoryUsage:      fmt.Sprintf("%d MB", m.Sys/1024/1024),
		Version: "V1",
	}
	_ = json.NewEncoder(w).Encode(&response)
}

// UpdateProduct - Handler to update a product
func (app *Application) UpdateProduct(w http.ResponseWriter, r *http.Request) {

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
	err = app.PService.UpdateProduct(code, prod)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(w, http.StatusAccepted)
}

// RemoveProduct - Handler to "remove" a product
func (app *Application) RemoveProduct(w http.ResponseWriter, r *http.Request) {
	code, err := utils.GetCode(r, "code")
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	err = app.PService.RemoveProduct(code)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	utils.SuccessResponse(w, http.StatusAccepted)
}

// GetProduct - Handler to find a product
func (app *Application) GetProduct(w http.ResponseWriter, r *http.Request) {
	code, err := utils.GetCode(r, "code")
	if err != nil {
		utils.HandlerError(w, r, http.StatusBadRequest, err.Error())
		return
	}
	response, err := app.PService.GetProduct(code)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(response)
	utils.SuccessResponse(w, 0)
}

// GetProducts - Handler to return all products
func (app *Application) GetProducts(w http.ResponseWriter, r *http.Request) {

	page, size := utils.GetQueryParams(r)
	response, err := app.PService.GetProducts(page, size)
	if err != nil {
		utils.HandlerError(w, r, http.StatusNotFound, err.Error())
		return
	}
	_ = json.NewEncoder(w).Encode(response)
	utils.SuccessResponse(w, 0)
}
