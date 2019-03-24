//Go is pass by value by default. To get reference use pointers.
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int
	var b float64

	a = 10
	b = 3.1415

	//Print variable's value.
	fmt.Println(a)
	// Print variable's address.
	fmt.Println(&a)
	add(a)

	//after calling add func.Print variable's value unchanged.
	fmt.Println(a)

	//Pointer type * symbol e.g. *int is pointer to int.

	var bAddress *float64
	bAddress = &b
	fmt.Println("b's address", bAddress, "stores type of : ", reflect.TypeOf(&b))

	fmt.Println(" value of b= ", b)

	// If you have address of the variable i.e. pointer than you can print its value by using *.
	// using * you can change value also.
	fmt.Println("Printing value at pointer *bAddress => ", *bAddress)

	*bAddress = 10.00

	// Value changed using pointers.
	fmt.Println("Printing new value at pointer *bAddress => ", *bAddress)
	fmt.Println("New value of b= ", b)

}

func add(a int) {
	a += a

	fmt.Println("inside add func:  ", a)
}
