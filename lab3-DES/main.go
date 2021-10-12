package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var data []int
	var res []string
	var cnt int
	key := make([]int, 64, 64)

	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	for _, val := range file {
		data = append(data, int(val))
	}
	for i := 0; i < len(key); i++ {
		key[i] = rand.Intn(2)
	}
	rand16Key := randKeyGen(key)
	if len(data) % 8 != 0 {
		for len(data) % 8 != 0 {
			data = append(data, 0)
			cnt++
		}
	}
	for i := 0; i < len(data)/8; i++ {
		arr := data[8 * i : 8 * (i + 1)]
		cipher := process(arr, rand16Key, "encrypt")
		res = append(res, strings.Trim(strings.Replace(fmt.Sprint(cipher), " ", ",", -1), "[]"))
	}
	fmt.Println(res[:len(res) - cnt])

}

