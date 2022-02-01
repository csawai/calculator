package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64
	var operator string

	fmt.Print("enter the first number: ")
	fmt.Scanln(&number1)

	fmt.Print("enter the second number: ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the operator: ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println("divide by zero situation")
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f", number1, operator, number2, number1/number2)
		}

	default:
		fmt.Println("invalid operator")
	}

}
