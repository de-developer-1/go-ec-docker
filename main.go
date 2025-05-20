package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"go-ec-docker/database"
	"go-ec-docker/handlers"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")

	r.HandleFunc("/cart", handlers.AddToCart).Methods("POST")
	r.HandleFunc("/cart", handlers.GetCart).Methods("GET")

	r.HandleFunc("/order", handlers.CreateOrder).Methods("POST")

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
