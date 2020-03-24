package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DB - embeds the *sql.DB
// so we can use it's methods
type DB struct {
	*sql.DB
}

// Error Model
type Error struct {
	Message string `json:"message"`
}

func main() {
	UserModel.initialiseTable()
	d := UserRepo.getTableData()
	fmt.Println(d)

	router := mux.NewRouter()

	router.HandleFunc("/signup", Signup).Methods("POST")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(protectedEndpoint)).Methods("POST")

	log.Println("Listen on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protected Endpoint invoked")
}
