package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func ProcessedHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=\"Результаты анализа стабильности оборудования.xlsx\"")
	begin, _ := time.Parse("2006-01-02 15:03:05", "2018-07-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:03:05", "2018-07-30 00:00:00")

	data := parse(begin, end)
	file := create(data)
	fmt.Println(file.DefinedNames)
	file.Write(w)
}

func AddElement(w http.ResponseWriter, r *http.Request) {
	pre := new(preload)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pre); err != nil {
		log.Fatal(err)
	}
	existPre := GetPreload()
	for k, v := range pre.Apn {
		if _, ok := pre.Apn[k]; !ok {
			existPre.Apn[k] = v
		}
		existPre.Apn[k] = append(existPre.Apn[k], pre.Apn[k]...)
	}
	file, err := os.Create("preload.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(existPre); err != nil {
		log.Fatal(err)
	}
}

func PrintElement(w http.ResponseWriter, r *http.Request) {
	pre := GetPreload()
	fmt.Println(pre.Converter())
}
