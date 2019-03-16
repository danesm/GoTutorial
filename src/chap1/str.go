//Using strings package.
package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "G# is C##l"
	//using string package's NewReplacer
	replacer := strings.NewReplacer("#", "o")
	str2 := replacer.Replace(str1)
  fmt.Println(str2)

}
