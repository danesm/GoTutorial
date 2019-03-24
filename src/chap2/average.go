package main

import "fmt"

func main() {

	numbers := [3]float64{2.2, 2.1, 2.4}
	var sum float64 = 0.0
	for _, num := range numbers {
		sum += num

	}

	avg := sum / float64(len(numbers))

	fmt.Println("average : ", avg)

}
