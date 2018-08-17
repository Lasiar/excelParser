package main

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

const (
	allColumn = 10
)

type PreLoad struct {
	Device `json:"Device"`
}

func loadFailure() (map[int]string, map[string]int) {
	failureSource := make(map[string]int)
	file, err := os.Open("failure.json")
	if err != nil {
		log.Fatal(err)
	}
	en := json.NewDecoder(file)
	en.Decode(&failureSource)

	failure := make(map[int]string)

	for k, v := range failureSource {
		failure[v] = k
	}
	return failure, failureSource
}

func create(data DataTable) *xlsx.File {
	headers := [allColumn]string{"Наименование оборудования",
		"Количество часов", "Отключение электроэнергии, часов",
		"Неисправность оборудования, часов",
		"Сбой программного обеспечения, часов",
		"Выработка ресурса сенсора, часов",
		"Отключение для технического обслуживания, часов",
		"Отключение для метрологического обслуживания, часов",
		"Стабильная работа, часов",
		"Стабильная работа*, %"}

	totalStyle := xlsx.NewStyle()

	totalStyle.Font.Bold = true

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("анализ работы оборудования")
	if err != nil {
		log.Fatal(err)
	}

	sheet.SetColWidth(0, 0, 34.40)

	row := sheet.AddRow()
	for _, header := range headers {
		cell := row.AddCell()
		cell.SetValue(header)
	}

	for point := range data.Apn {
		rowName := sheet.AddRow()

		cell := rowName.AddCell()
		cell.SetValue(point)
		total := [6]float64{}
		for device := range data.Apn[point] {
			row := sheet.AddRow()
			cell := row.AddCell()
			cell.SetValue(device)
			cell = row.AddCell()
			cell.SetFloat(data.AllHours)
			var unstableWork float64
			for i := 0; i < 6; i++ {
				cell := row.AddCell()
				cell.SetFloat(data.Apn[point][device][i])
				unstableWork += data.Apn[point][device][i]
				total[i] += data.Apn[point][device][i]
			}
			cell = row.AddCell()
			cell.SetFloat(data.AllHours - unstableWork)
			cell = row.AddCell()
			cell.SetFloat(100 - (unstableWork / data.AllHours * 100))
			cell.SetStyle(totalStyle)
		}
		row := sheet.AddRow()
		cell = row.AddCell()
		cell.SetValue("итого")
		cell.SetStyle(totalStyle)
		cell = row.AddCell()
		cell.SetFloat(data.AllHours * float64(len(data.Apn[point])))
		for i := 0; i < 6; i++ {
			cell := row.AddCell()
			cell.SetStyle(totalStyle)
			cell.SetFloat(total[i])
		}
	}
	return file
}
