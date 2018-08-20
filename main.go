package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/add", AddElement)
	http.HandleFunc("/print", PrintElement)
	http.HandleFunc("/", ProcessedHandle)

	http.ListenAndServe(":8080", nil)
}
