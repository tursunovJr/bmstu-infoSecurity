package main

import (
	"fmt"
	"math/rand"
	"os"
)

func createRotorMap() map[int]int {
	res := make(map[int]int, 256)
	nums := make(map[int]bool)
	for i := 0; i < 256; i++ {
		nums[i] = false
	}
	for i := 0; i < 256; i++ {
		tmp := rand.Intn(256)
		for nums[tmp] != false {
			tmp = rand.Intn(256)
		}
		res[i] = tmp
		nums[tmp] = true
	}
	return res
}

func createReflectorMap() map[int]int {
	res := make(map[int]int, 256)
	for i := 0; i < 256; i++ {
		tmp := rand.Intn(256)
		_, exist1 := res[tmp]
		_, exist2 := res[i]
		if exist1 == false && exist2 == false {
			res[i] = tmp
			res[tmp] = i
		}
	}
	return res
}

func main() {
	var rotorsList []map[int]int
	var data []int
	file, err := os.ReadFile("archive.zip")
	if err != nil {
		panic(err)
	}
	for _, val := range(file) {
		data = append(data, int(val))
	}
	//fmt.Println("Before: ", file)
	for i := 0; i < 3; i++ {
		rotorsList = append(rotorsList, createRotorMap())
	}
	reflectorMap := createReflectorMap()
	fmt.Println("After: ", dataProcessing(rotorsList, reflectorMap, data))

}
