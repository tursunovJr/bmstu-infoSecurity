package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertNum(val string, base, toBase int) string {
	i, _ := strconv.ParseInt(val, base, 64)
	return strconv.FormatInt(i, toBase)
}

func fullPack(num string, bitNum int) string {
	for len(num) < bitNum {
		num = num + "0"
	}
	return num
}

func moveLeft(arr []int, shift int) []int {
	for i := 0; i < shift; i++ {
		tmp := arr[0]
		for j := 0; j < len(arr) - 1; j++ {
			arr[j] = arr[j + 1]
		}
		arr[len(arr) - 1] = tmp
	}
	return arr
}

func initHalfKey(from, to, in []int) {
	for i, val := range in{
		to[i] = from[val]
	}
}

func randKeyGen(key []int) [16][48]int {
	var resKeys [16][48]int
	var currentKey [48]int
	left := make([]int, 28, 28)
	right := make([]int, 28, 28)
	initHalfKey(key, left, C0())
	initHalfKey(key, right, D0())
	for i := 1; i < 17; i++ {
		left = moveLeft(left, Shifts()[i-1])
		right = moveLeft(right, Shifts()[i-1])
		for j, val := range CP() {
			if j < len(CP())/2 {
				currentKey[j] = left[val - 1]
			} else {
				currentKey[j] = right[val - 29]
			}
		}
		//fmt.Println(currentKey)
		resKeys[i-1] = currentKey
	}
	return resKeys
}

func permut(data []int, table []int) []int {
	res := make([]int, len(table), len(table))
	for i, val := range table{
		res[i] = data[val-1]
	}
	return res
}

func xor(a []int, b [48]int) []int {
	res := make([]int, len(b), len(b))
	if len(a) != 48 {
		panic("Xor can`t be done. Slices have different size")
	} else {
		for i := 0; i < len(a); i++ {
			res[i] = a[i]^b[i]

		}
	}
	return res
}

func splitToBlocks(data []int, size int) [][]int{
	count := len(data) / size
	var blocks [][]int
	for i := 0; i < count; i++ {
		blocks = append(blocks, data[size * i : size * (i + 1)])
	}
	return blocks
}

func cipherFeistel(data []int, key [48]int) []int {
	var z, res []int
	data = permut(data, E())
	z = xor(data, key)
	blocks := splitToBlocks(z, 6)
	for i := 0 ; i < len(blocks); i++ {
		block := blocks[i]
		row, _ := strconv.Atoi(convertNum(strconv.Itoa(block[0]) + strconv.Itoa(block[5]), 2, 10))
		column, _ := strconv.Atoi(convertNum(strings.Trim(strings.Replace(fmt.Sprint(block[1:5]), " ", "", -1), "[]"),2,10))
		value := S()[i][row][column]
		binValue := convertNum(strconv.Itoa(value), 10, 2)
		binValue = fullPack(binValue, 4)

		for _, value := range binValue {
			valToInt, _ := strconv.Atoi(string(value))
			res = append(res, valToInt)
		}
	}
	return permut(res, P())
}




