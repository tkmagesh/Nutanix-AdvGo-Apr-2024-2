package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

type AppServer struct {
	handlers    map[string]http.HandlerFunc
	middlewares []MiddlewareFunc
}

func NewAppServer() *AppServer {
	return &AppServer{
		handlers: make(map[string]http.HandlerFunc),
	}
}

func (appServer *AppServer) Use(pattern string, handler http.HandlerFunc) {
	for _, middleware := range appServer.middlewares {
		handler = middleware(handler)
	}
	appServer.handlers[pattern] = handler
}

func (appServer *AppServer) UseMiddleware(middleware MiddlewareFunc) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := appServer.handlers[r.URL.Path]
	if handler != nil {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

// application
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello World!")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All customer details will be served")
}

func log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v %v\n", r.Method, r.URL.Path)
		handler(w, r)
		return
	}
}

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}

func timeout(d time.Duration) MiddlewareFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			timeoutCtx, cancel := context.WithTimeout(r.Context(), d)
			go func() {
				<-timeoutCtx.Done()
				fmt.Fprintln(w, "timeout occurred", http.StatusRequestTimeout)
			}()
			defer cancel()
			handler(w, r.WithContext(timeoutCtx))
			if err := timeoutCtx.Err(); err != nil {
				fmt.Println("timeout :", err)
			}
		}
	}
}

func main() {
	appServer := NewAppServer()

	appServer.UseMiddleware(timeout(3 * time.Second))
	appServer.UseMiddleware(log)
	appServer.UseMiddleware(profile)

	appServer.Use("/", IndexHandler)
	appServer.Use("/products", ProductsHandler)
	appServer.Use("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		fmt.Println("error starting server :", err)
	}
}
