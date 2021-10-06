package main

import (
	"fmt"
	"math/rand"
)

func main() {
	key := make([]int, 64, 64)
	for i := 0; i < len(key); i++ {
		key[i] = rand.Intn(256)
	}
	fmt.Println(key)
	fmt.Println("Results = ",randKeyGen(key))

}
