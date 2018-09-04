package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/add", AddElement)
	http.HandleFunc("/", ProcessedHandle)
	//	http.HandleFunc("/delete", DeleteHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
