package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Trishank-K/BookStore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	if os.Getenv("GO_ENV") == "" {
		log.Println("Running in development mode. Set GO_ENV=production for production mode.")
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
