package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articles")
	Articles := make([]Article, 2)
	Articles[0] = Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"}
	Articles[1] = Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"}
	json.NewEncoder(w).Encode(Articles)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
