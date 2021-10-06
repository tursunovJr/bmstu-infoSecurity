package main

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
		resKeys[i-1] = currentKey
	}
	return resKeys
}


