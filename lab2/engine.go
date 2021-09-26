package main

type Rotor struct {
	rotor map[int]int
	shift int
}

func shiftRotor(rotor map[int]int) map[int]int {
	tmp := rotor[255]
	for i := 255; i > 0 ; i-- {
		rotor[i] = rotor[i - 1]
	}
	rotor[0] = tmp
	return rotor
}

func encode(key int, rotor1, rotor2, rotor3 Rotor, reflector map[int]int) int{
	tmp1 := rotor1.rotor[key]
	tmp2 := rotor2.rotor[tmp1]
	tmp3 := rotor3.rotor[tmp2]
	refTmp := reflector[tmp3]
	tmp3Back := rotor3.rotor[refTmp]
	tmp2Back := rotor2.rotor[tmp3Back]
	tmp1Back := rotor1.rotor[tmp2Back]
	return tmp1Back
}
func dataProcessing(rotorsList []map[int]int, reflector map[int]int, data []int) []int {
	var res []int
	rotor1 := Rotor{rotor: rotorsList[0], shift: 0}
	rotor2 := Rotor{rotor: rotorsList[1], shift: 0}
	rotor3 := Rotor{rotor: rotorsList[2], shift: 0}
	for val := range data {
		res = append(res, encode(val, rotor1, rotor2, rotor3, reflector))
		rotor1.shift++
		rotor1.rotor = shiftRotor(rotor1.rotor)
		if rotor1.shift > 256 {
			rotor1.shift = 0
			rotor2.shift++
			rotor2.rotor = shiftRotor(rotor2.rotor)
			if rotor2.shift > 256 {
				rotor1.shift = 0
				rotor1.shift = 0
				rotor3.shift++
				rotor3.rotor = shiftRotor(rotor3.rotor)
			}
		}
	}
	return res
}
