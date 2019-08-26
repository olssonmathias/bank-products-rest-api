package controllers

import (
	"bank-products-rest-api/data"
	"encoding/json"
	"fmt"
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
			fmt.Fprintln(w, "Updated product "+id)
			return
		}
	}
	fmt.Fprintln(w, "Could not find product "+id)
}
