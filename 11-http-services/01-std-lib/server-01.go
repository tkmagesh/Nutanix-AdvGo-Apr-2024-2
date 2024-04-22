package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products = []Product{
	{100, "pen", 10},
	{101, "marker", 50},
	{102, "pencil", 5},
}

type AppServer struct {
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "error parsing data", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "error parsing input data", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			w.WriteHeader(http.StatusCreated)
		default:
			http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All customer details will be served")
	default:
		http.Error(w, "resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		fmt.Println("error starting server :", err)
	}
}
