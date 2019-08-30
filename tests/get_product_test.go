package tests

import (
	"bank-products-rest-api/controllers"
	"bank-products-rest-api/data"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

type GETObject struct {
	ID   string `json:"id,omitempty"`
	Bank string `json:"bank,omitempty"`
	Name string `json:"name,omitempty"`
	AER  string `json:"aer,omitempty"`
}

func TestGetProduct(t *testing.T) {
	data.PopulateDummyData()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")

	for i := 1; i < len(data.ProductData)+2; i++ {
		id := strconv.Itoa(i)

		request := buildGetRequest(t, id)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		status := validateGetStatus(t, response)
		validateGetHeaders(t, response, status, id)
		if status == http.StatusOK {
			validateGetBody(t, id, response)
		}
	}
}

func buildGetRequest(t *testing.T, id string) *http.Request {
	request, error := http.NewRequest("GET", "/product/"+id, nil)
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}
	return request
}

func validateGetStatus(t *testing.T, response *httptest.ResponseRecorder) int {
	status := response.Code
	if status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("Unexpected HTTP response code: %d", status)
	}
	return status
}

func validateGetHeaders(t *testing.T, response *httptest.ResponseRecorder, status int, id string) {
	if status == http.StatusOK {
		expectedContentType := "application/json"
		contentType := response.Header().Get("Content-Type")

		if contentType != expectedContentType {
			t.Errorf("Content-Type (expected): %v", expectedContentType)
			t.Errorf("Content-Type (actual): %v", contentType)
		}
	} else if status == http.StatusNotFound {
		expectedLocation := "/product/" + id
		location := response.Header().Get("Location")

		if location != expectedLocation {
			t.Errorf("Location (expected): %v", expectedLocation)
			t.Errorf("Location (actual): %v", location)
		}
	}
}

func validateGetBody(t *testing.T, id string, response *httptest.ResponseRecorder) {
	var expectedBody GETObject

	error := json.Unmarshal(response.Body.Bytes(), &expectedBody)
	if error != nil {
		t.Errorf("Response body: %v", response.Body.String())
		t.Errorf("ERROR: %v", error)
	}
}
