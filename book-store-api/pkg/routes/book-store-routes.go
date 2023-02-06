package routes

import (
	"book-store-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
