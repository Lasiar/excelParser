package main

import "net/http"

func main() {
	http.HandleFunc("/",ProcessedHandle)
	http.ListenAndServe(":8080", nil)
}
