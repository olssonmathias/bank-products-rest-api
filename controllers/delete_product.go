package controllers

import (
	"bank-products-rest-api/data"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, product := range data.ProductData {
		if product.ID == id {
			data.ProductData = append(data.ProductData[:index], data.ProductData[index+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.Header().Set("location", "/product/"+id)
	w.WriteHeader(http.StatusNotFound)
}
