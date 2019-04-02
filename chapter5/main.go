package main

import (
	"fmt"
	"strconv" // Convert ints to string
)

func main() {
	// Introducing...loops!

	for i := 0; i <= 100; i = i + 3 {
		if i%3 == 0 && i != 0 {
			fmt.Println(i)
		}
	}

	for i := 1; i <= 100; i++ {
		fizzAndOrBuzz := ""

		if i%3 != 0 && i%5 != 0 {
			fizzAndOrBuzz = strconv.Itoa(i)

			/*
				Converts int to string. Attempting to use string(i) will have the
				number interpretted as a unicode value -- fun!
			*/

		} else {
			if i%3 == 0 {
				fizzAndOrBuzz = "Fizz"
			}

			if i%5 == 0 {
				fizzAndOrBuzz += "Buzz"
			}
		}

		fmt.Println(fizzAndOrBuzz)
	}
}
