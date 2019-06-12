package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"build-rest-api/controllers"
	"build-rest-api/app"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(app.JwtAuthentication)
	myRouter.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	myRouter.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	myRouter.HandleFunc("/api/contact/new", controllers.CreateContact).Methods("POST")
	myRouter.HandleFunc("/api/{id}/contacts", controllers.GetContacts).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API")
	handleRequest()
}
