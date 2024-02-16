package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWritter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Methos is not supported", http.StatusNoContent)
		return
	}

	fmt.Fprint(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server ata port 8000\n")
	if err := http.ListenAndServer(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
