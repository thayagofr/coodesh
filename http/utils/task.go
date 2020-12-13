package utils

import (
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/thyagofr/coodesh/desafio/database"
	"github.com/thyagofr/coodesh/desafio/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// InitializeCron - InitializeCron
func InitializeCron() {
	c := cron.New()
	_, err := c.AddFunc(
		"34 11 * * *",
		LoadData,
	)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Houve um erro ao agendar tarefa de sincronizacao...")
	}
	c.Start()
}

func GetFileName() string {

	URL := os.Getenv("OPENFOOD_FILES_LIST")
	response, err := http.Get(URL)
	if err != nil {
		log.Println("Nao foi possível obter a lista de arquivos.")
		log.Fatal(err.Error())
	}
	defer response.Body.Close()
	scanner := bufio.NewScanner(response.Body)
	scanner.Scan()
	return scanner.Text()
}

func LoadData() {
	name := GetFileName()
	URL := fmt.Sprintf("https://static.openfoodfacts.org/data/delta/%s", name)
	response, _ := http.Get(URL)
	defer response.Body.Close()
	gzipReader, err := gzip.NewReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(gzipReader)
	convert := strings.ReplaceAll(string(data), "}\n{", "},{")
	newString := "[" + convert[:len(convert)-1] + "]"

	var products []model.Product
	err = json.NewDecoder(strings.NewReader(newString)).Decode(&products)
	if err != nil {
		log.Println(err)
	}
	total := 0
	for _, p := range products {
		p.Status = GetStatus(PUBLISHED)
		p.ImportedT = time.Now().UTC()
		total += InsertProduct(p)
		if total == 100 {
			break
		}
	}
	LoadDataLog()
}

func InsertProduct(product model.Product) int {
	collection := database.GetCollection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Println(err)
		log.Print("Erro ao inserir produto")
		return 0
	}
	return 1
}

func LoadDataLog() {
	collection := database.GetCollection("history")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, model.History{
		RunningT: time.Now(),
	})
	if err != nil {
		log.Println("Houve um erro ao registrar execução do CRON")
	}
}
