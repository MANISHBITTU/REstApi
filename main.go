package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticle(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{Title: "Text Title", Desc: "Text Description", Content: "Hello word"},
	}
	fmt.Println("Endpoint hit: All Article Endpoint")
	json.NewEncoder(w).Encode(article)
}

func main() {
	handleRequest()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my page")
	fmt.Println("Endpoint hit here")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homepage)
	router.HandleFunc("/articles", allArticle)
	log.Fatal(http.ListenAndServe(":8090", router))
}
