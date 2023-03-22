package main

import (
	"fmt"
	"log"
	"net/http"
	"test/controller"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/insert", controller.Insert).Methods("POST")
	router.HandleFunc("/update", controller.Update).Methods("POST")
	router.HandleFunc("/delete", controller.Delete).Methods("DELETE")
	router.HandleFunc("/viewall", controller.ViewAll).Methods("GET")
	router.HandleFunc("/findproduct", controller.FindProduct).Methods("GET")
	router.HandleFunc("/likeordislike", controller.LikeOrDislike).Methods("POST")
	router.HandleFunc("/", helloWorld)
	log.Fatal(http.ListenAndServe(":8080", router))
}
