package tests

import (
	"bank-products-rest-api/controllers"
	"bank-products-rest-api/data"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type POSTObject struct {
	ID   string `json:"id,omitempty"`
	Bank string `json:"bank,omitempty"`
	Name string `json:"name,omitempty"`
	AER  string `json:"aer,omitempty"`
}

func TestCreateProduct(t *testing.T) {
	data.PopulateDummyData()
	existingProduct := POSTObject{
		ID:   "1",
		Bank: "Nationwide",
		Name: "FlexDirect",
		AER:  "5%",
	}
	newProduct := POSTObject{
		ID:   "4",
		Bank: "Nationwide",
		Name: "FlexAccount",
		AER:  "0%",
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")

	runPostTest(t, router, newProduct, newProduct.ID)
	runPostTest(t, router, existingProduct, existingProduct.ID)
}

func runPostTest(t *testing.T, router *mux.Router, product POSTObject, id string) {
	request := buildPostRequest(t, product)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if validatePostStatus(t, response) == http.StatusConflict {
		validatePostHeaders(t, response, id)
	}
}

func buildPostRequest(t *testing.T, product POSTObject) *http.Request {
	jsonProduct, error := json.Marshal(product)
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}

	request, error := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonProduct))
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}

	return request
}

func validatePostStatus(t *testing.T, response *httptest.ResponseRecorder) int {
	status := response.Code
	if status != http.StatusOK && status != http.StatusConflict {
		t.Errorf("Unexpected HTTP response code: %d", status)
	}
	return status
}

func validatePostHeaders(t *testing.T, response *httptest.ResponseRecorder, id string) {
	expectedLocation := "/product/" + id
	location := response.Header().Get("Location")

	if location != expectedLocation {
		t.Errorf("Location (expected): %v", expectedLocation)
		t.Errorf("Location (actual): %v", location)
	}
}
