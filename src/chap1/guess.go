//Game
// Challange a player to guess a random number in 10 attempts.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	/*
	 To get different random numbers, we're going to need to pass a value to the rand. Seed function.
	 That will "seed" the random number generator, that is, give it a value that it will use to generate other random values.
	 But if we keep giving it the same seed value, it will keep giving us the same random values, and we'll be back where we started.
	*/

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	//Addin 1 to target to becuase we want to range it between 1-100.
	target := rand.Intn(100) + 1

	fmt.Println("Target number is randomly choosen between 1-100. Lets see if you can guess it? ")

	success := false
	//fmt.Print("\n Guess the number: ")
	inputbuf := bufio.NewReader(os.Stdin)

	for guess := 1; guess < 10; guess++ {

		fmt.Println("You have", 10-guess, "guesses left after this!")
		fmt.Print("\n \t Make a guess:   ")
		inputstr, error := inputbuf.ReadString('\n')
		inputstr = strings.TrimSpace(inputstr)
		userinput, error := strconv.Atoi(inputstr)

		if error != nil {
			log.Fatal(error)
		}

		if userinput < target {

			fmt.Println("Oops! your guess was LOW.. ")
		} else if userinput > target {

			fmt.Println("Oops! your guess was HIGH! .. ")
		} else {
			success = true

			fmt.Println("Well done ..you got it right")
			break
		}

	}

	if success != true {
		fmt.Println("Sorry you couldn't guess the right number. Which is  : ", target)
	}

}
