package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Please enter your percentage: ")
	//Reading from standard input.
	reader := bufio.NewReader(os.Stdin)
	//Read until user press enter.
	input, error := reader.ReadString('\n')

	/*The input string still has a newline character on the end,
	  from when the user pressed the Enter key. */

	input = strings.TrimSpace(input)
	grade, error := strconv.ParseFloat(input, 64)

	if error != nil {
		log.Fatal(error)
	}

	//Note result is defined out of block scope of if-else.
	var result string
	if grade >= 60 && grade <= 100 {
		result = "Passed!!"
	} else if grade > 100 {
		result = "Your result can not be more than 100% try again"
	} else {
		result = "failed"

	}
	fmt.Println(result)

}
