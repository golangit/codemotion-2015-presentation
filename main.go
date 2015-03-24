package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "" || r.URL.Path == "/" {
		http.Error(w, "Must Contain a Name", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s.\n", r.URL.Path[1:])
}

func main() {
	listenAndServe(helloHandler)
}

func listenAndServe(handlerFunc http.HandlerFunc) {
	log.Println("Starting 8080")
	http.HandlerFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

func mainChaining() {
	listenAndServe(logger(single(helloHandler, "localhost:8080")))
}
