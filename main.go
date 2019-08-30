package main

import (
	"bank-products-rest-api/controllers"
	"bank-products-rest-api/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	data.PopulateDummyData()
	handleRequests()
}
