//A simple porgram to check pass or fail based on percentage of the student.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Enter a value: ")
	reader := bufio.NewReader(os.Stdin)
	// In GO: function can return any number of values. Unlike other prog languages.

	//userinput, error := reader.ReadString('\n')  - This also fails as we don't use error var in our prog.

	// Using blank identifier (single underscore) to show that we are ignoring second returned value from the func ReadString.
	//	userinput, _ := reader.ReadString('\n')

	//Second option is to handle the error using log.Fatal func.

	userinput, error := reader.ReadString('\n')

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(userinput)

}
