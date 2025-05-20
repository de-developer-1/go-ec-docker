package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dsn := "host=localhost port=5432 user=ecuser password=ecpass dbname=ecdemo sslmode=disable"

	var err error
	for i := 0; i < 10; i++ {
		DB, err = sql.Open("postgres", dsn)
		if err == nil && DB.Ping() == nil {
			break
		}
		log.Println("Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	createTables()
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			price INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS cart (
			id SERIAL PRIMARY KEY,
			product_id INTEGER,
			quantity INTEGER,
			FOREIGN KEY (product_id) REFERENCES products(id)
		);`,
		`CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			total INTEGER NOT NULL
		);`,
	}
	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatal("Error creating table:", err)
		}
	}
}
