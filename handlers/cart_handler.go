package handlers

import (
	"encoding/json"
	"net/http"
	"go-ec-docker/database"
	"go-ec-docker/models"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var item models.CartItem
	json.NewDecoder(r.Body).Decode(&item)

	_, err := database.DB.Exec("INSERT INTO cart (product_id, quantity) VALUES ($1, $2)", item.ProductID, item.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	rows, _ := database.DB.Query("SELECT id, product_id, quantity FROM cart")
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var i models.CartItem
		rows.Scan(&i.ID, &i.ProductID, &i.Quantity)
		items = append(items, i)
	}
	json.NewEncoder(w).Encode(items)
}
