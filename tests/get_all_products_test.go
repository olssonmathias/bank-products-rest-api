package tests

import (
	"bank-products-rest-api/controllers"
	"bank-products-rest-api/data"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type ExpectedJSONObject struct {
	ID   string `json:"id,omitempty"`
	Bank string `json:"bank,omitempty"`
	Name string `json:"name,omitempty"`
	AER  string `json:"aer,omitempty"`
}

type ExpectedJSONObjects []ExpectedJSONObject

func TestGetAllProducts(t *testing.T) {
	data.PopulateDummyData()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")

	request := buildRequest(t)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	validateGetAllStatus(t, response)
	validateGetAllHeaders(t, response)
	validateGetAllBody(t, response)
}

func buildRequest(t *testing.T) *http.Request {
	request, error := http.NewRequest("GET", "/products", nil)
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}
	return request
}

func validateGetAllStatus(t *testing.T, response *httptest.ResponseRecorder) {
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Unexpected HTTP response code: %d", status)
	}
}

func validateGetAllHeaders(t *testing.T, response *httptest.ResponseRecorder) {
	expectedContentType := "application/json"
	contentType := response.Header().Get("Content-Type")

	if contentType != expectedContentType {
		t.Errorf("Content-Type (expected): %v", expectedContentType)
		t.Errorf("Content-Type (actual): %v", contentType)
	}
}

func validateGetAllBody(t *testing.T, response *httptest.ResponseRecorder) {
	var expectedJSON ExpectedJSONObjects

	error := json.Unmarshal(response.Body.Bytes(), &expectedJSON)
	if error != nil {
		t.Errorf("Response body: %v", response.Body.String())
		t.Errorf("ERROR: %v", error)
	}
}
