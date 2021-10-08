package main

import (
	"fmt"
	"math/rand"
)

func main() {
	key := make([]int, 64, 64)
	data := make([]int, 32, 32)
	for i := 0; i < len(key); i++ {
		key[i] = rand.Intn(2)
	}
	for i := 0; i < 32; i++ {
		data[i] = rand.Intn(2)
	}
	rand16Key := randKeyGen(key)
	fmt.Println(len(cipherFeistel(data, rand16Key[1])))

}
