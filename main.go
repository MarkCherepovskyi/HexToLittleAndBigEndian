package main

import (
	"Lab/practic2/lib"
	"log"
)

func main() {
	str := "aaaa000000000000000000000000000000000000000000000000000000000000"
	size := 32
	res := lib.SearchingValue(str)
	log.Println("search", res)
	res2 := lib.ConvertHexToDec(res)
	log.Println("little endian ", res2)
	res3 := lib.ConvertHexToDec(str)
	log.Println("big endian ", res3)
	res4 := lib.ConvertToHex(res3, size)
	log.Println("big to hex ", res4)
	res5 := lib.ConvertToHex(res2, size)
	log.Println("little to hex ", res5)
}
