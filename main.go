// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program sees if your integer is positive or negative

package main

import "fmt"

func main() {
	var userNumber int

	fmt.Println("Enter your integer to see if it’s positive or negative:")
	fmt.Println()
	fmt.Print("Integer: ")
	fmt.Scanln(&userNumber)

	if userNumber < 0 {
		print("Your integer is a negative number.")
	} else {
		print("Your integer is a positive number.")
	}
}
