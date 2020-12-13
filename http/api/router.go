package api

import (
	"github.com/gorilla/mux"
	"github.com/thyagofr/coodesh/desafio/service"
	"github.com/thyagofr/coodesh/desafio/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(client *mongo.Client) *mux.Router {

	pservice := service.PService{Client: client}
	hservice := service.HService{Client: client}
	app := Application{
		PService: &pservice,
		HService: &hservice,
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Use(utils.LogMiddleware)
	router.HandleFunc("/api/v1/", app.Home)
	router.HandleFunc("/api/v1/products", app.GetProducts).Methods("GET")
	router.HandleFunc("/api/v1/products/{code}", app.GetProduct).Methods("GET")
	router.HandleFunc("/api/v1/products/{code}", app.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/v1/products/{code}", app.RemoveProduct).Methods("DELETE")
	return router
}
