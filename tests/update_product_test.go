package tests

import (
	"bank-products-rest-api/controllers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type PUTObject struct {
	ID   string `json:"id,omitempty"`
	Bank string `json:"bank,omitempty"`
	Name string `json:"name,omitempty"`
	AER  string `json:"aer,omitempty"`
}

func TestUpdateProduct(t *testing.T) {
	existingProduct := PUTObject{
		ID:   "1",
		Bank: "Nationwide",
		Name: "FlexDirect",
		AER:  "5%",
	}
	newProduct := PUTObject{
		ID:   "4",
		Bank: "Nationwide",
		Name: "FlexAccount",
		AER:  "0%",
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")

	runPutTest(t, router, newProduct, newProduct.ID)
	runPutTest(t, router, existingProduct, existingProduct.ID)
}

func runPutTest(t *testing.T, router *mux.Router, product PUTObject, id string) {
	request := buildPutRequest(t, product, id)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if validatePutStatus(t, response) == http.StatusNotFound {
		validatePutHeaders(t, response, id)
	}
}

func buildPutRequest(t *testing.T, product PUTObject, id string) *http.Request {
	jsonProduct, error := json.Marshal(product)
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}

	request, error := http.NewRequest("PUT", "/product/"+id, bytes.NewBuffer(jsonProduct))
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}

	return request
}

func validatePutStatus(t *testing.T, response *httptest.ResponseRecorder) int {
	status := response.Code
	if status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("Unexpected HTTP response code: %d", status)
	}
	return status
}

func validatePutHeaders(t *testing.T, response *httptest.ResponseRecorder, id string) {
	expectedLocation := "/product/" + id
	location := response.Header().Get("Location")

	if location != expectedLocation {
		t.Errorf("Location (expected): %v", expectedLocation)
		t.Errorf("Location (actual): %v", location)
	}
}
