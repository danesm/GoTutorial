//READING FILE PROPERTIES
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileinfo, error := os.Stat("hello.go")

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(fileinfo.Size())

}
