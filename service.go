package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	_oncePreload sync.Once
	_preload     *preload
)

type preload struct {
	Reason []string         `json:"reason"`
	Device []string         `json:"device"`
	Apn    map[string][]int `json:"apn"`
}

func (p *preload) IdReason(find string) int {
	for i, reason := range p.Reason {
		if reason == find {
			return i
		}
	}
	return -1
}

func (p *preload) Converter() Apn {
	apn := make(Apn)

	for point, idsDevice := range p.Apn {

		if _, ok := apn[point]; !ok {
			apn[point] = make(Device)
		}

		for idDevice := range idsDevice {

			if _, ok := apn[point][p.Device[idDevice]]; !ok {
				apn[point][p.Device[idDevice]] = make(ValueByMaint)
			}

			for i := range p.Reason {
				apn[point][p.Device[idDevice]][i] = 0.0
			}

		}

	}
	return apn
}

func (p *preload) load() {
	file, err := os.Open("preload.json")
	if err != nil {
		log.Fatal(err)
	}
	jsDec := json.NewDecoder(file)
	if err := jsDec.Decode(p); err != nil {
		log.Fatal(err)
	}
}

func GetPreload() *preload {
	_oncePreload.Do(func() {
		_preload = new(preload)
		_preload.load()
	})
	return _preload
}
