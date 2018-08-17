package main

import (
	"github.com/tealeg/xlsx"
	"log"
	"time"
)

const (
	locateName  = 3
	deviceName  = 4
	timeBegin   = 8
	timeEnd     = 9
	CauseColumn = 23
)

type DataTable struct {
	Apn
	AllHours float64
}

type Apn map[string]Device
type Device map[string]ValueByMaint
type ValueByMaint map[int]float64

func parse(begin, end time.Time) DataTable {

	xlFile, err := xlsx.OpenFile("table.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	xlFile.Date1904 = false

	sheet := xlFile.Sheet["БД"]

	dataTable := DataTable{}

	dataTable.AllHours = float64(end.Unix()-begin.Unix()) / 3600

	cause := GetPreload()

	for i, row := range sheet.Rows {

		if i < 2 {
			continue
		}

		cells := row.Cells

		if cells[timeBegin].String() == "" {
			continue
		}

		t0, err := cells[timeBegin].GetTime(false)
		if err != nil {
			log.Printf("Ошибка чтении первой даты %v %v", err, i)
		}

		t1, err := cells[timeEnd].GetTime(false)
		if err != nil {
			t1 = end
		}

		if t0.After(end) || t1.Before(begin) {
			continue
		}

		if t0.Before(begin) {
			t0 = begin
		}
		if t1.After(end) {
			t1 = end
		}

		dt := float64(t1.Unix()-t0.Unix()) / 3600

		if _, ok := cause[cells[locateName].String()]; !ok {
			cause[cells[locateName].String()] = make(Device)
		}
		if _, ok := cause[cells[locateName].String()][cells[deviceName].String()]; !ok {
			cause[cells[locateName].String()][cells[deviceName].String()] = make(ValueByMaint)
		}

		v, ok := failure[cells[CauseColumn].String()]
		if !ok {
			continue
		}

		cause[cells[locateName].String()][cells[deviceName].String()][v] += dt

	}
	dataTable.Apn = cause
	return dataTable
}
