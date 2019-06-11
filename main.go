package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"build-rest-api/controllers"
)

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	myRouter.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Rest API")
	handleRequest()
}
