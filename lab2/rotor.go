package main

import (
	"strings"
)

const ALPHA_LABELS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type RotorStruct struct {
	wiring, stepping string
}

var RotorsList = map[string]RotorStruct {
	"I" : {
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		"Q",
	},
	"II" : {
		"AJDKSIRUXBLHWTMCQGZNPYFVOE",
		"E",
	},
	"III" : {
		"BDFHJLCPRTXVZNYEIWGAKMUSQO",
		"V",
	},
	"IV" : {
		"ESOVPZJAYQUIRHXLNFTGKDCMWB",
		"J",
	},
	"V" : {
		"VZBRGITYUPSDNHLXAWMJQOFECK",
		"Z",
	},
}

var ReflectorsList = map[string]string {
	"B" : "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	"C" : "FVPJIAOYEDRZXWGCTKUQSBNMHL",
}

type Rotor struct {
	modelName, wiring, stepping string
	alphaMap, posMap map[string]int
	//ringSettings int

}

type RotorFunctionality interface {
	Init() Rotor
}

func (r Rotor) Init() Rotor{
	if len(r.wiring) != 26{
		panic("Invalid wiring length")
	}
	for _, val := range r.wiring {
		if strings.Contains(ALPHA_LABELS, string(val)) == false {
			panic("Invalid wiring")
		}
	}
	if r.alphaMap == nil {
		r.alphaMap = make(map[string]int)
	}
	if r.posMap == nil {
		r.posMap = make(map[string]int)
	}
	for _, val := range ALPHA_LABELS {
		r.alphaMap[string(val)] = int(val - 65)
	}
	for pos, val := range r.wiring {
		r.posMap[string(val)] = pos
	}
	//fmt.Println("alphamap = ", r.alphaMap)
	//fmt.Println("posmap = ", r.posMap)

	//todo Добавить проверку на частоту повтора букв
	//todo Проверка ring_settings
	return r
}

func createRotor(model string) Rotor{
	for key, _ := range RotorsList {
		if key == model {
			data := RotorsList[key]
			obj := Rotor{modelName: model, wiring: data.wiring, stepping: data.stepping}
			return obj.Init()
		}
	}
	return Rotor{}
}

func createReflector(model string) Rotor{
	for key, _ := range ReflectorsList {
		if key == model {
			obj := Rotor{modelName: model, wiring: ReflectorsList[key]}
			return obj.Init()
		}
	}
	return Rotor{}
}

