package tests

import (
	"bank-products-rest-api/controllers"
	"bank-products-rest-api/data"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestDeleteProduct(t *testing.T) {
	data.PopulateDummyData()
	router := mux.NewRouter().StrictSlash(true)

	for i := 1; i < len(data.ProductData)+2; i++ {
		id := strconv.Itoa(i)

		request := buildDeleteRequest(t, id)
		response := httptest.NewRecorder()

		router.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
		router.ServeHTTP(response, request)

		if validateDeleteStatus(t, response) == http.StatusConflict {
			validateDeleteHeaders(t, response, id)
		}
	}
}

func buildDeleteRequest(t *testing.T, id string) *http.Request {
	request, error := http.NewRequest("DELETE", "/product/"+id, nil)
	if error != nil {
		t.Errorf("ERROR: %v", error)
	}
	return request
}

func validateDeleteStatus(t *testing.T, response *httptest.ResponseRecorder) int {
	status := response.Code
	if status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("Unexpected HTTP response code: %d", status)
	}
	return status
}

func validateDeleteHeaders(t *testing.T, response *httptest.ResponseRecorder, id string) {
	expectedLocation := "/product/" + id
	location := response.Header().Get("Location")

	if location != expectedLocation {
		t.Errorf("Location (expected): %v", expectedLocation)
		t.Errorf("Location (actual): %v", location)
	}
}
