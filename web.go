package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type message struct {
	Select preload `json:"select"`
	Month  string  `json:"month"`
}

func ProcessedHandle(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	//w.Header().Set("Content-Disposition", "attachment; filename=\"Результаты анализа стабильности оборудования.xlsx\"")

	m := new(message)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		log.Println(err)
	}

	if m.Month == "" {
		log.Println("bad reqest")
	}

	date, err := time.Parse("2006-01", m.Month)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(m)

	data := parse(date, date.AddDate(0, 1, 0).Add(-time.Nanosecond), m.Select)

	fmt.Println(data)

	file := create(data)

	file.Write(w)
}

func AddElement(w http.ResponseWriter, r *http.Request) {
	pre := new(preload)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pre); err != nil {
		log.Fatal(err)
	}
	existPre := GetPreload()

	if pre.Device != nil {
		existPre.Device = append(existPre.Device, pre.Device...)
	}

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
