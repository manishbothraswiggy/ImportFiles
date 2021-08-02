package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/manishbothraswiggy/ImportFiles/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()

	jsonResponse, err := json.Marshal(productList)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
	rw.Write(jsonResponse)
}
