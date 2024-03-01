package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "210804Ernur"
	dbname   = "golang"
)

type Credentials struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Link        string  `json:"link"`
}

func itemsAdd(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Connect to PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Insert user credentials into the auth table
	_, err = db.Exec("INSERT INTO items (name, description, price, link) VALUES ($1, $2, $3, $4)", creds.Name, creds.Description, creds.Price, creds.Link)
	if err != nil {
		http.Error(w, "Failed to insert item into database", http.StatusInternalServerError)
		fmt.Println("Error inserting item into database:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Item added successfully")
}
func itemsUpd(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Connect to PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Update item in the items table
	_, err = db.Exec("UPDATE items SET name=$1, description=$2, price=$3, link=$4 WHERE id=$5", creds.Name, creds.Description, creds.Price, creds.Link, id)
	if err != nil {
		http.Error(w, "Failed to update item in database", http.StatusInternalServerError)
		fmt.Println("Error updating item in database:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Item updated successfully")
}

func itemsDel(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Connect to PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Update item in the items table
	_, err = db.Exec("DELETE FROM items WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Failed to delete item in database", http.StatusInternalServerError)
		fmt.Println("Error deleting item in database:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Item deleted successfully")
}

func itemsGet(w http.ResponseWriter, r *http.Request) {
	// Connect to PostgreSQL database
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query all items from the items table
	rows, err := db.Query("SELECT name, description, price, link FROM items")
	if err != nil {
		http.Error(w, "Failed to retrieve items from database", http.StatusInternalServerError)
		fmt.Println("Error retrieving items from database:", err)
		return
	}
	defer rows.Close()

	var items []Credentials
	for rows.Next() {
		var item Credentials
		if err := rows.Scan(&item.Name, &item.Description, &item.Price, &item.Link); err != nil {
			http.Error(w, "Failed to scan item row", http.StatusInternalServerError)
			fmt.Println("Error scanning item row:", err)
			return
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over item rows", http.StatusInternalServerError)
		fmt.Println("Error iterating over item rows:", err)
		return
	}

	// Encode the items to JSON and write response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, "Failed to encode items to JSON", http.StatusInternalServerError)
		fmt.Println("Error encoding items to JSON:", err)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items/itemsAdd", itemsAdd).Methods("POST")
	r.HandleFunc("/items/itemsUpd/{id}", itemsUpd).Methods("POST")
	r.HandleFunc("/items/itemsDel/{id}", itemsDel).Methods("DELETE")
	r.HandleFunc("/items/itemsGet", itemsGet).Methods("GET")
	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", r)

	// Convert port to string
	portStr := strconv.Itoa(port)

	// Connection string
	connStr := "postgres://" + user + ":" + password + "@" + host + ":" + portStr + "/" + dbname + "?sslmode=disable"

	// Connect to  database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection by pinging the database
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("Connected to the PostgreSQL database successfully!")

}
