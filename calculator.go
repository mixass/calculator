package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started script")
	startFunction()
}

var accept string
var operation string
var number1 int
var number2 int

func startFunction() {
	println("Operation type (+,-,x,%)")
	fmt.Scan(&operation)
	if operation == "x" {
		SelectedMultiply()
	} else if operation == "+" {
		SelectedPlus()
	} else if operation == "-" {
		SelectedMinus()
	} else if operation == "%" {
		SelectedDivide()
	} else {
		println("Invalid operation")
	}
}

func SelectedMultiply() {
	fmt.Println("You want to operation x . Do you accept? (Y/N)")
	fmt.Scan(&accept)
	if accept == "y" {
		fmt.Println("Enter the first number.")
		fmt.Scan(&number1)
		fmt.Println("Enter the second number.")
		fmt.Scan(&number2)
		multiply := number1 * number2
		fmt.Println("\nResult : ", multiply)
	} else if accept == "n" {
		fmt.Println("Canceled.")
	} else {
		fmt.Println("Invalid Value")
	}
}

func SelectedPlus() {
	fmt.Println("You want to operation toplama. Do you accept? (y/n)")
	fmt.Scan(&accept)
	if accept == "y" {
		fmt.Println("accepted")
		fmt.Println("Enter the first number.")
		fmt.Scan(&number1)
		fmt.Println("Enter the second number.")
		fmt.Scan(&number2)
		plus := number1 + number2
		fmt.Println("\nResult : ", plus)
	} else if accept == "n" {
		fmt.Println("Canceled.")
	} else {
		fmt.Println("Invalid Value")
	}
}

func SelectedMinus() {
	fmt.Println("You want to operation cikartma. Do you accept? (Y/N)")
	fmt.Scan(&accept)
	if accept == "y" {
		fmt.Println("accepted")
		fmt.Println("Enter the first number.")
		fmt.Scan(&number1)
		fmt.Println("Enter the second number.")
		fmt.Scan(&number2)
		minus := number1 - number2
		fmt.Println("\nResult : ", minus)
	} else if accept == "n" {
		fmt.Println("Canceled.")
	} else {
		fmt.Println("Invalid Value")
	}
}

func SelectedDivide() {
	fmt.Println("You want to operation bolum. Do you accept? (Y/N)")
	fmt.Scan(&accept)
	if accept == "y" {
		fmt.Println("accepted")
		fmt.Println("Enter the first number.")
		fmt.Scan(&number1)
		fmt.Println("Enter the second number.")
		fmt.Scan(&number2)
		divide := number1 % number2
		fmt.Println("\nResult : ", divide)
	} else if accept == "n" {
		fmt.Println("Canceled.")
	} else {
		fmt.Println("Invalid Value")
	}
}
