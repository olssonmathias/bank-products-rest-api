package controllers

import (
	"bank-products-rest-api/data"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, product := range data.ProductData {
		if product.ID == id {
			data.ProductData = append(data.ProductData[:index], data.ProductData[index+1:]...)
			fmt.Fprintln(w, "Removed product "+id)
			return
		}
	}
	fmt.Fprintln(w, "Could not find product "+id)
}
