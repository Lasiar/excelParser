package main

import (
	"net/http"
)

func ProcessedHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=\"Результаты анализа стабильности оборудования.xlsx\"")
	//begin, _ := time.Parse("2006-01-02 15:03:05", "2018-07-01 00:00:00")
	//end, _ := time.Parse("2006-01-02 15:03:05", "2018-07-30 00:00:00")

	//parse(begin, end)
	//	file := create(data)
	//	fmt.Println(file.DefinedNames)
	//	file.Write(w)
}
