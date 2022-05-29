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
		"a": 10,
		"b": 11,
		"c": 12,
		"d": 13,
		"e": 14,
		"f": 15,
	}
)

func ConvertHexToDec(hexNum string) *big.Int {
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

	res := fmt.Sprintf("%x", num)
	log.Println(res)

	if len(res) != size*2 {
		for i := len(res); i < size*2; i++ {
			res += "0"
		}
	}
	return res
}
