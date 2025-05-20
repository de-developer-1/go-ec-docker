package handlers

import (
	"encoding/json"
	"net/http"
	"go-ec-docker/database"
	"go-ec-docker/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, _ := database.DB.Query("SELECT id, name, price FROM products")
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	json.NewDecoder(r.Body).Decode(&p)

	_, err := database.DB.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", p.Name, p.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
