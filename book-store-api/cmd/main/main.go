package main

import (
	"book-store-api/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Start servert at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
