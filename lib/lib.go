package lib

import (
	"fmt"
	"log"
	"math/big"
)

var (
	valuesInt = map[string]int64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}
)

func ConvertHexToDec(hexNum string) *big.Int {
	hexNum = hexNum[2:]
	bufDecimal := new(big.Int)
	for index, _ := range hexNum {
		data := fmt.Sprintf("%c", hexNum[index])
		var bufData int64
		for key, value := range valuesInt {
			if data == key {
				bufData = value
				break
			}
		}
		powBig := new(big.Int).Exp(big.NewInt(16), big.NewInt(int64(len(hexNum)-(index+1))), nil)
		multyBig := new(big.Int).Mul(powBig, big.NewInt(bufData))
		bufDecimal = bufDecimal.Add(multyBig, bufDecimal)
	}
	return bufDecimal
}

func SearchingValue(hexNum string) string {
	var bufIndex int
	for index, data := range hexNum {
		if data != '0' {
			bufIndex = index
		}

	}
	res := hexNum[0 : bufIndex+1]
	return res
}

func ConvertToHex(num *big.Int, size int) string {

	bufferRes := fmt.Sprintf("%x", num)
	log.Println(bufferRes)

	if len(bufferRes) != size*2 {
		for i := len(bufferRes); i < size*2; i++ {
			bufferRes += "0"
		}
	}
	res := "0x" + bufferRes
	return res
}
