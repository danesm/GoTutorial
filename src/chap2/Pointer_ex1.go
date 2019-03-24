package main

import "fmt"

func main() {

	num := 10

	fmt.Println("Value of num before double func call:=> ", num)
	double(num)
	fmt.Println("Value of num after double func call:=> ", num)
	doublePoint(&num)
	fmt.Println("Value of num after double func call:=> ", num)

}

// pass by value doesn't change value.
func double(a int) {

	a += 10

}

// pointer changed the value
func doublePoint(a *int) {

	*a += 10
}
