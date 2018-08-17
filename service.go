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
	Reason map[string]int      `json:"reason"`
	Apn    map[string][]string `json:"apn"`
}

func (p *preload) Converter() Apn {
	apn := make(Apn)
	for point, wantedArr := range p.Apn {
		if _, ok := apn[point]; !ok {
			apn[point] = make(Device)
		}
		for _, wanted := range wantedArr {
			if _, ok := apn[point][wanted]; !ok {
				apn[point][wanted] = make(ValueByMaint)
			}
		}
	}
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
