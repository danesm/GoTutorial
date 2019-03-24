//Find a squreroot of a float64/int.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func squareroot(number float64) (float64, error) {

	if number < 0 {
		return 0, errors.New("-ve number:can't get squre root for this")

		//OR

		//return 0, fmt.Errorf("-ve number:can't get squre root for this")

	}

	return math.Sqrt(number), nil
}

func main() {

	// Taking input from std input.
	fmt.Print("Enter a float number: ")
	inb := bufio.NewReader(os.Stdin)
	inputStr, err := inb.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	inputStr = strings.TrimSpace(inputStr)
	input, err := strconv.ParseFloat(inputStr, 64)

	if err != nil {
		fmt.Println(err)
	}
	// Calculate Square root
	sroot, err := squareroot(input)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%0.2f\n", sroot)

}
