package main

import (
	"fmt"
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
		var operation int
		fmt.Println("Welcome For Calculation")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Scan(&operation)

		if operation > 5 {
			fmt.Println("Invalid Input!!!!!")
			break
		}

		if operation == 5 {
			fmt.Println("Goodbye ROYALTY")
			break
		}

				var input1, input2 int
		fmt.Println("Enter Input 1: ")
		fmt.Scan(&input1)
		fmt.Println("Enter Input 2: ")
		fmt.Scan(&input2)

		switch operation {
		case 1:
			fmt.Println(add(input1, input2))
		case 2:
			fmt.Println(sub(input1, input2))
		case 3:
			fmt.Println(mul(input1, input2))
		case 4:
			fmt.Println(div(input1, input2))
		default:
			fmt.Println("Not Supported")
			continue
		}
	}


}