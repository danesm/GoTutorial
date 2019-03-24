package main

import "fmt"

func main() {

	//var citylist [4]string = [4]string {"A","B","C,"D"}

	clist := [4]string{"A", "B", "C", "D"}

	fmt.Println("array length :", len(clist))

	fmt.Println("From Normal For loop:")

	for i := 0; i < len(clist); i++ {
		fmt.Println(i, clist[i])
	}

	// looping arrays using 'range'
	fmt.Println("From Range Loop : ")
	for index, value := range clist {

		fmt.Println(index, value)
	}

	//if index not needed. Use blank identifier.
	fmt.Println("From Range Loop with blank identifier (return value only) : ")
	for _, val := range clist {
		fmt.Println(val)
	}

}
