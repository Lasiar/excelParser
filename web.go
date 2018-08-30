package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type message struct {
	Select preload `json:"select"`
	Month  string  `json:"month"`
}

func ProcessedHandle(w http.ResponseWriter, r *http.Request) {
	m := new(message)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		log.Println(err)
	}

	if m.Month == "" {
		log.Println("bad reqest")
		return
	}

	date, err := time.Parse("2006-01", m.Month)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=\"Результаты анализа стабильности оборудования за "+date.Format("01 2006")+".xlsx\"")

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

	if err := existPre.Save(); err != nil {
		log.Println(err)
	}

}

//func DeleteHandle(w http.ResponseWriter, r *http.Request){
//	pre := new(preload)
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&pre); err != nil {
//		log.Fatal(err)
//	}
//	existPre := GetPreload()
//
//}
