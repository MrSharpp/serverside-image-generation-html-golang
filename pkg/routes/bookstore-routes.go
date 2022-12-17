package routes

import (
	bookController "testproject/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", bookController.CreateBook).Methods("POST")
	router.HandleFunc("/book/", bookController.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", bookController.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", bookController.UpDateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", bookController.DeleteBook).Methods("DELETE")
}
