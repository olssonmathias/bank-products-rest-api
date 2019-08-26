package controllers

import (
	"bank-products-rest-api/data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func PostProduct(w http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		log.Fatal(error)
	}

	var newProduct data.Product
	json.Unmarshal(requestBody, &newProduct)

	for _, existingProduct := range data.ProductData {
		if existingProduct.ID == newProduct.ID {
			fmt.Fprintln(w, "Product with ID "+newProduct.ID+" already exists.")
			return
		}
	}
	data.ProductData = append(data.ProductData, newProduct)
	fmt.Fprintln(w, "Added product "+newProduct.ID)
}
