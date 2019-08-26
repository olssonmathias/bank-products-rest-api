package main

import (
	"bank-products-rest-api/controllers"
	"bank-products-rest-api/data"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	myRouter.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	myRouter.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
	myRouter.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	myRouter.HandleFunc("/product", controllers.PostProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	data.PopulateDummyData()
	handleRequests()
}
