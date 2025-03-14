package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hak-vichet1111/reamstack-golang/db"
)

var DB *sql.DB

func main() {
	// connect database 
	DB = db.ConnectDB()

	// Initialize router
	router := mux.NewRouter()

	// Define a simple endpoint
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	router.HandleFunc("/users", getUserHandler).Methods("GET")

	// Start server
	fmt.Println("Server is running on port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write a simple response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, world!")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM public.profiles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var id int 
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, fmt.Sprintf("ID: %d, Name: %s", id, name))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v", users)
}


// Run project: compiledaemon --command="./reamstack-golang"