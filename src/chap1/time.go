// Using time package
package main

import (
	"fmt"
	"time"
)

func main() {

	//var now time.Time = time.Now()
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()

	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

}
