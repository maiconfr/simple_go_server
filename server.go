package main

import (
	"fmt"
	"log"
	"net/http"
	"server-go/functions"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", functions.FormHandler)
	http.HandleFunc("/hello", functions.HelloHandler)

	fmt.Printf("Starting server at port 8080.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
