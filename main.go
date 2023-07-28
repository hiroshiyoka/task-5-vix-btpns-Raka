package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/controllers/authcontroller"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/controllers/photocontroller"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/middleware"
	"github.com/hiroshiyoka/task-5-vix-btpns-Raka/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter().StrictSlash(true)

	// User
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	// Photo
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/photos", photocontroller.Index).Methods("GET")
	api.HandleFunc("/photo/{id}", photocontroller.Show).Methods("GET")
	api.HandleFunc("/photo", photocontroller.Create).Methods("POST")
	api.HandleFunc("/photo/{id}", photocontroller.Update).Methods("PUT")
	api.HandleFunc("/photo", photocontroller.Delete).Methods("DELETE")

	// Middleware
	api.Use(middleware.JwtMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
