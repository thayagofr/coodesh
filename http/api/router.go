package api

import (
	"github.com/gorilla/mux"
	"github.com/thyagofr/coodesh/desafio/http/utils"
)

func Routes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(utils.LogMiddleware)
	router.HandleFunc("/api/v1/", Home)
	router.HandleFunc("/api/v1/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/products/{code}", GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/products/{code}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/products/{code}", RemoveProduct).Methods("DELETE")
	return router

}
