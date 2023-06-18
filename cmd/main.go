package main

import (
	"log"
	"net/http"

	// "gorm.io/driver/mysql"
	"readinglist-api/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.UseReadingListRoutes(r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
