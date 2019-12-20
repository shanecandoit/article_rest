package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// follow https://tutorialedge.net/golang/creating-restful-api-with-golang/

type Article struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// keep in memory to simulate a database
var Articles []Article

// router
var myRouter *mux.Router

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	// router
	myRouter = mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	// single article
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	// all articles
	myRouter.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":9090", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {

	// articles
	Articles = []Article{
		Article{
			ID:      "2",
			Title:   "Home",
			Desc:    "Home Desc",
			Content: "Home Content"},
		Article{
			ID:      "2",
			Title:   "About Us",
			Desc:    "About Description",
			Content: "About Content"},
	}

	handleRequests()
}
