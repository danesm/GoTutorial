// Func returning mulitple values.
package main

import "fmt"

func returnmany() (int, bool, string) {

	return 2, true, "Hello"
}

func main() {

	num, boole, strg := returnmany()

	fmt.Println(num, boole, strg)

}
