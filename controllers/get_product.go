package controllers

import (
	"bank-products-rest-api/data"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, product := range data.ProductData {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	w.Header().Set("location", "/product/"+id)
	w.WriteHeader(http.StatusNotFound)
}
