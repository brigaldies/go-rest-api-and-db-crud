package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func CreateTable(db *sql.DB) (sql.Result, error) {
	result, err := db.Exec(tableCreationQuery)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func ClearTable(db *sql.DB) {
	fmt.Println("clearTable...")
	_, err := db.Exec("DELETE FROM products")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func AddProducts(db *sql.DB, count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		_, err := db.Exec("INSERT INTO products(name, price) VALUES($1, $2)", "Product "+strconv.Itoa(i), (i+1.0)*10)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
