package api

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/thyagofr/coodesh/desafio/database"
	"github.com/thyagofr/coodesh/desafio/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Arquivo .env nao encontrado")
	}
	mongoClient = database.InitDatabase()
}

func TestRemoveProductNotFound(t *testing.T) {
	ts := httptest.NewServer(Routes(mongoClient))
	defer ts.Close()
	client := ts.Client()
	request, err := http.NewRequest("DELETE", ts.URL + "/api/v1/products/1", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("WANT : %d, GOT : %d", http.StatusNotFound, response.StatusCode )
	}
}

func TestRemoveProductAccept(t *testing.T) {
	ts := httptest.NewServer(Routes(mongoClient))
	defer ts.Close()
	client := ts.Client()
	request, err := http.NewRequest("DELETE", ts.URL + "/api/v1/products/0681131911962", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusAccepted {
		t.Errorf("WANT : %d, GOT : %d", http.StatusAccepted, response.StatusCode )
	}
}

func TestGetProductOk(t *testing.T) {
	ts := httptest.NewServer(Routes(mongoClient))
	defer ts.Close()
	client := ts.Client()
	request, err := http.NewRequest("GET", ts.URL + "/api/v1/products/0681131911962", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("WANT : %d, GOT : %d", http.StatusAccepted, response.StatusCode )
	}
}

func TestGetProductNotFound(t *testing.T) {
	ts := httptest.NewServer(Routes(mongoClient))
	defer ts.Close()
	client := ts.Client()
	request, err := http.NewRequest("GET", ts.URL + "/api/v1/products/0075011", nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("WANT : %d, GOT : %d", http.StatusAccepted, response.StatusCode )
	}
}

func TestUpdateProductNotFound(t *testing.T) {
	ts := httptest.NewServer(Routes(mongoClient))
	defer ts.Close()
	client := ts.Client()

	request, err := http.NewRequest("PUT", ts.URL + "/api/v1/products/0075011", strings.NewReader(JsonValid()))
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("WANT : %d, GOT : %d", http.StatusNotFound, response.StatusCode )
	}
}

func TestUpdateProductAccept(t *testing.T) {
	ts := httptest.NewServer(Routes(mongoClient))
	defer ts.Close()
	client := ts.Client()

	request, err := http.NewRequest("PUT", ts.URL + "/api/v1/products/0681131911962", strings.NewReader(JsonValid()))
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusAccepted {
		t.Errorf("WANT : %d, GOT : %d", http.StatusAccepted, response.StatusCode )
	}
}

func JsonValid() string {
	body := utils.UpdateProductRequest{
		MainCategory: "Teste",
		NutriscoreGrade: "10",
		NutriscoreScore: 10,
		IngredientsText: "Lalalalalal",
		Brands: "Nestle",
		Quantity: "10",
		ProductName: "Produto caro",
	}
	bodyJSON , _ := json.Marshal(&body)
	return string(bodyJSON)
}