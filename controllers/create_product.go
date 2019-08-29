package controllers

import (
	"bank-products-rest-api/data"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var newProduct data.Product
	json.Unmarshal(requestBody, &newProduct)

	for _, existingProduct := range data.ProductData {
		if existingProduct.ID == newProduct.ID {
			w.Header().Set("location", "/product/"+existingProduct.ID)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
	data.ProductData = append(data.ProductData, newProduct)
	w.WriteHeader(http.StatusOK)
}
