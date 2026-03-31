package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hexToDec(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

func decBinHex(decStr string) (string, string) {
	dec, err := strconv.ParseInt(decStr, 10, 64)
	if err != nil {
		fmt.Println("Not a Decimal")
	}
	return strconv.FormatInt(dec, 2), strings.ToUpper(strconv.FormatInt(dec, 16))
}

func main() {
	for {
		var operator string
		var input string

		fmt.Println("Enter input: ")
		fmt.Scan(&input)

		fmt.Println("Welcome To Base Converter")
		fmt.Println("What Base Would You Like To Convert To?")
		fmt.Println("1. HEXADECIMAL")
		fmt.Println("2. BINARY")
		fmt.Println("3. DECIMAL")
		fmt.Println("4. Exit")
		fmt.Scan(&operator)

		switch operator {
		case "1":
			fmt.Println(hexToDec(input))
		case "2":
			fmt.Println(binToDecimal(input))
		case "3":
			bin, hex := decBinHex(input)
			fmt.Println("Binary: ", bin)
			fmt.Println("HexaDecimal: ", hex)
		case "4":
			fmt.Println("Thank You ROYALTY")
			return
		}
	}
}
