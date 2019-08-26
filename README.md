bank-products-rest-api
----------------------

This is a very basic example of a REST API written in Go.

**Prerequisites**

1. Go distribution: https://golang.org/doc/install
2. A 3rd party package (makes routing of requests easier): https://github.com/gorilla/mux  
`go get -u github.com/gorilla/mux`
3. Some kind of API testing tool: https://www.getpostman.com/downloads/

**Try it out**

1. `git clone https://github.com/olssonmathias/bank-products-rest-api.git`
2. `go run main.go`
3. http://localhost:8081

```
GET     /products     Return all products  
GET     /product/2    Return one product by id  
DELETE  /product/2    Delete one product by id  
PUT     /product/2    Update one product by id and as according to the request body  
POST    /product      Create one product as according to the request body  
```

When testing with a POST or PUT method, use the following request body format.

```
{
  "id": "2",
  "bank": "Bank X",
  "name": "Product Y",
  "aer": "percentage"
}
```
