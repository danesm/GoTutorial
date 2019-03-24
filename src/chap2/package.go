package main

import (
	"fmt"
	"keyboard"
)

func main() {

	// Get user input from keyboard.
	floatval, err := keyboard.GetFloat()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(" Invoked keyboard package inside 'package.go' to get float : ", floatval)

}
