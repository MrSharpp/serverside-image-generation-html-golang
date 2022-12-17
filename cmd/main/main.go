package main

import (
	"fmt"
	"log"
	"net/http"
	"testproject/pkg/config"
	"testproject/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.Connect()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Server started lisening to 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
