package main

import (
	"github.com/tealeg/xlsx"
	"log"
	"strings"
	"time"
)

const (
	locateName  = 3
	deviceName  = 4
	model       = 5
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

func (d Device) stringFind(findString ...string) string {
	str := strings.Join(findString, "")
	find := ""
	for k := range d {
		if strings.Index(str, k) > -1 {
			find = k
		}
	}
	return find
}

type ValueByMaint map[int]float64

func parse(begin, end time.Time) DataTable {

	xlFile, err := xlsx.OpenFile("test.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	xlFile.Date1904 = false

	sheet := xlFile.Sheet["БД"]

	dataTable := DataTable{}

	dataTable.AllHours = float64(end.Unix()-begin.Unix()) / 3600

	preload := GetPreload()

	list := preload.Converter()

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

		if _, ok := list[cells[locateName].String()]; !ok {
			continue
		}

		failure := list[cells[locateName].String()].stringFind(cells[deviceName].String(), cells[model].String())

		if _, ok := list[cells[locateName].String()][failure]; !ok {
			continue
		}

		v, ok := preload.Reason[cells[CauseColumn].String()]
		if !ok {
			continue
		}

		list[cells[locateName].String()][failure][v] += dt

	}
	dataTable.Apn = list
	return dataTable
}
