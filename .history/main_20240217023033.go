package main

import (
	"fmt"
	"net/http"
)

func main() {
fileServer := http.FileServer(http.Dir("./static"))
http.Handle("/", fileServer)
http.HandleFunc("/form",formHandler)
http.HandleFunc("/hello", helloHandler)

fmt.Println("Starting server ata port 8000\n")
if err  http.ListenAndServer(":8000", nil);
err != nil {
	log.Fatal(err)
}

}