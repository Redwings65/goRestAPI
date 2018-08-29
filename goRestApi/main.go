package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Shoe Size Restful API")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", welcome).Methods("GET")
	myRouter.HandleFunc("/addshoesize/{size}", addShoeSize).Methods("GET", "POST")
	myRouter.HandleFunc("/gettotal", returnTrueSize).Methods("GET")
	myRouter.HandleFunc("/getall", getAllShoes).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	initialMigration()
	fmt.Print("Running Server...")
	handleRequests()
}
