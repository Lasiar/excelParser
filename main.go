package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"log"
)

func main() {
	p := GetPreload()
	p.Converter()
	begin, _ := time.Parse("2006-01-02 15:03:05", "2018-07-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:03:05", "2018-10-30 00:00:00")
	b := parse(begin, end)
	fmt.Println(b)
	file := create(b)
	w, err := os.Create("out.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(w)
	http.HandleFunc("/", ProcessedHandle)
	http.ListenAndServe(":8080", nil)
}
