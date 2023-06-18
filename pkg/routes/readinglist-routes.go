package routes

import (
	"readinglist-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var UseReadingListRoutes = func(router *mux.Router) {
	router.HandleFunc("/readinglist", controllers.GetList).Methods("GET")
	router.HandleFunc("/readinglist", controllers.AddBook).Methods("POST")
	router.HandleFunc("/readinglist/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/readinglist/{id}", controllers.Modify).Methods("PUT")
	router.HandleFunc("/readinglist/{id}", controllers.Delete).Methods("DELETE")
}
