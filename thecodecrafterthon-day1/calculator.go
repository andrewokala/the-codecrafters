package main

import (
	"fmt"
	"strconv"
)


	func add(input1, input2 int) int {
	return input1 + input2
	}
	func sub(input1, input2 int) int {
	return input1 - input2
	}

	func mul(input1, input2 int) int {
	return input1 * input2
	}

	func div(input1, input2 int) int {
	return input1 / input2
	}

func main() {
	for	{
		var operation string

		var num1, num2 string
		fmt.Println("Enter Input 1: ")
		fmt.Scan(&num1)
		input1, err1 := strconv.Atoi(num1)
		if err1 != nil {
			fmt.Println("Must Be A Number")
			continue
			}
		fmt.Println("Enter Input 2: ")
		fmt.Scan(&num2)
		input2, err2 := strconv.Atoi(num2)
		if err2 != nil {
			fmt.Println("Must Be A Number")
			continue
		}

		fmt.Println("Welcome For Calculation")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Println("help for Use cases")
		fmt.Scan(&operation)

		switch operation {
		case "1":
			fmt.Println(add(input1, input2))
		case "2":
			fmt.Println(sub(input1, input2))
		case "3":
			fmt.Println(mul(input1, input2))
		case "4":
			if input2 == 0 {
				fmt.Println("Not A Divisor")
				continue
			}
			fmt.Println(div(input1, input2))
		case "5":
			fmt.Println("Goodbye ROYALTY")
			return
		case "help":
			fmt.Println("USE CASE")
			fmt.Println("Input 1 For Addition")
			fmt.Println("Input 2 For Subtraction")
			fmt.Println("Input 3 For Multiplication")
			fmt.Println("Input 4 For Addition")
			fmt.Println("Input 5 To Exit")
		
		default:
			fmt.Println("Not Supported")
			continue
		}
	}
}