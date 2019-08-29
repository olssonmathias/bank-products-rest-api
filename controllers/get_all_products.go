package controllers

import (
	"bank-products-rest-api/data"
	"encoding/json"
	"net/http"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.ProductData)
}
