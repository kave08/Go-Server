package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on Port 8080 \n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
	http.Error(w, "404 not found", http.StatusNotFound)
	return
	}
	if r.Method != "GET" {
	http.Error(w ,"method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w,"hello")
}
