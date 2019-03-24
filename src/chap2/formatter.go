package main

import "fmt"

func main() {

	fmt.Println("Println output:", 1.0/3.0)

	// Printf func == Print with formatting.
	fmt.Printf("Printf output:  %0.2f\n", 1.0/3.0)

	/*The Sprintf function (also part of the fmt package) works just like Printf,
		except that it returns a formatted string instead of printing it.

		1. Formatting verbs (the %0.2f in the strings above is a verb
	  2. Value widths (that’s the 0.2 in the middle of the verb) */

	newstring := fmt.Sprintf("%0.2f\n", 1.0/3.0)
	fmt.Println("Sprintf output:", newstring)

	fmt.Printf("A Float : %f\n", 3.1415)

	// To return error from the func , use fmt.Errorf("error msg").
	err := fmt.Errorf("Test error")
	fmt.Println(err)

}
