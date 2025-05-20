package handlers

import (
	"net/http"
	"go-ec-docker/database"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	row := database.DB.QueryRow(`
		SELECT COALESCE(SUM(p.price * c.quantity), 0)
		FROM cart c
		JOIN products p ON p.id = c.product_id
	`)
	var total int
	row.Scan(&total)

	_, err := database.DB.Exec("INSERT INTO orders (total) VALUES ($1)", total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Exec("DELETE FROM cart")

	w.WriteHeader(http.StatusCreated)
}
