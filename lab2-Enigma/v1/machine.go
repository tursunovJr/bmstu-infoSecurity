package v1

import (
	"bufio"
	"bytes"
	"encoding/base32"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EnigmaMachine struct {
	rotors []string
	reflectorKey, ringSettings string
	reflectorList []Rotor
	rotorList []Rotor
	//rotorsCount int
}

func existCheck(symbol, searchString string) bool {
	for _, val := range searchString {
		if string(val) == symbol {
			return true
		}
	}
	return false
}

func checkIsDigit(symbol string) bool {
	val, _ := strconv.Atoi(symbol)
	for i := 0; i < 10; i++ {
		if val == i {
			return true
		}
	}
	return false
}

type Machine interface {
	Init() EnigmaMachine
	Rotate()
	ProcessText(text string) string
	Encode(symbol string) string
	Decode(symbol string) string
}


func (m EnigmaMachine) Init() EnigmaMachine {
	if len(m.rotors) != 3 && len(m.rotors) !=4 {
		panic("Rotors count doesn't match")
	}
	for _, val := range(m.rotors) {
		m.rotorList = append(m.rotorList, createRotor(val))
	}
	m.reflectorList = append(m.reflectorList, createReflector(m.reflectorKey))
	return m
}

func getReflectorValue(m EnigmaMachine, key, step string) string {
	var rotor string
	pos := m.rotorList[0].alphaMap[key] - m.rotorList[0].alphaMap[step]
	if pos < 0 {
		pos = 26 + pos
	}
	if pos > 25 {
		pos = pos - 26
	}
	for key, val := range m.rotorList[0].alphaMap {
		if val == pos {
			rotor = key
		}
	}
	return rotor
}

func (m EnigmaMachine) Encode(symbol string) string {
	var rotor, stepping string
	//fmt.Println("-------------------------------------------------------------------------------")
	//fmt.Println("Symbol = ", symbol)
	for i := 0; i < len(m.rotorList); i++ {
		//fmt.Println("------------------Before---------------")
		//fmt.Println("i = ", i, "rotor = ", rotor)
		stepping = m.rotorList[i].stepping
		//fmt.Println("Current stepping = ", stepping)
		steppingPos := m.rotorList[i].alphaMap[stepping]
		if i == 0 {
			pos := m.rotorList[i].alphaMap[symbol] + steppingPos
			if pos > 25 {
				pos = pos - 26
			}
			for key, val := range m.rotorList[i].posMap {
				if val == pos {
					rotor = key
				}
			}
		} else {
			backStepping := m.rotorList[i-1].stepping
			//fmt.Println("Back stepping = ", backStepping)
			backSteppingPos := m.rotorList[i-1].alphaMap[backStepping]
			steppingDifference := steppingPos - backSteppingPos
			//fmt.Println("SteppingDifference = ", steppingDifference)
			pos := m.rotorList[i].alphaMap[rotor] + steppingDifference
			if pos < 0 {
				pos = 26 + pos
			}
			if pos > 25 {
				pos = pos - 26
			}
			//fmt.Println("POS = ", pos)
			for key, val := range m.rotorList[i].posMap {
				if val == pos {
					rotor = key
				}
			}
		}
		//fmt.Println("------------after--------------")
		//fmt.Println("i = ", i, "rotor = ", rotor)
	}
	rotor = getReflectorValue(m, rotor, stepping)
	return rotor
}

func (m EnigmaMachine) Decode(symbol string) string {
	var signal, stepping string
	fmt.Println("symbol = ", symbol)
	pos := m.rotorList[0].alphaMap[symbol]
	for key, val := range m.reflectorList[0].posMap {
		if val == pos {
			signal = key
		}
	}
	//fmt.Println("Signal before = ", signal)
	for i := len(m.rotorList) - 1; i >= 0; i-- {
		stepping = m.rotorList[i].stepping
		steppingPos := m.rotorList[i].alphaMap[stepping]
		if i == len(m.rotorList) - 1 {
			pos := m.rotorList[i].alphaMap[signal] + steppingPos
			if pos > 25 {
				pos = pos - 26
			}
			for key, val := range m.rotorList[i].alphaMap {
				if val == pos {
					signal = key
				}
			}
		} else {
			backStepping := m.rotorList[i+1].stepping
			backSteppingPos := m.rotorList[i+1].alphaMap[backStepping]
			steppingDifference := backSteppingPos - steppingPos
			pos := m.rotorList[i].alphaMap[signal] - steppingDifference
			if pos < 0 {
				pos = 26 + pos
			}
			if pos > 25 {
				pos = pos - 26
			}
			for key, val := range m.rotorList[i].alphaMap {
				if val == pos {
					signal = key
				}
			}
		}
	}
	signal = getReflectorValue(m, signal, stepping)
	return signal
}

func (m EnigmaMachine) ProcessText(text, mode string) string {
	var result, c string
	for _, val := range text {
		c = strings.ToUpper(string(val))
		if existCheck(c, ALPHA_LABELS) == false {
			if c == "=" || checkIsDigit(c) {
				result += c
				continue
			} else {
				panic("Illegal symbol")
			}
		}
		if mode == "encode" {
			result += m.Encode(c)
		} else {
			result += m.Decode(c)
		}


	}
	return result
}


func main() {
	mode := "decode"

	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := bytes.Buffer{}
	buf := bufio.NewScanner(file)
	for buf.Scan() {
		content.WriteString(buf.Text())
	}
	res := base32.StdEncoding.EncodeToString([]byte(content.String()))
	tmp := EnigmaMachine{rotors: []string{"I", "II", "III"}, reflectorKey: "B"}
	tmp = tmp.Init()
	_, err = f.WriteString(tmp.ProcessText(res, mode))
	if err != nil {
		panic(err)
	}
}

