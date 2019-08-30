package controllers

import (
	"bank-products-rest-api/data"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var productData data.Product
	json.Unmarshal(requestBody, &productData)

	for index, product := range data.ProductData {
		if product.ID == id {
			data.ProductData[index] = productData
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.Header().Set("location", "/product/"+id)
	w.WriteHeader(http.StatusNotFound)
}
