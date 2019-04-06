package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 3; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// Go supports concurrency

func main() {
	for i := 0; i < 3; i++ {
		go f(i)

		// The `go` keyword forces the compiler to move on, whether
		// f(0) has returned yet or not. It's a bit like `async`
		// in Swift. Both line 21 and our fave, main(), are examples
		// of Goroutines.
	}

	var input string
	fmt.Scanln(&input)

	// The call to Scanln is required to keep the script from finishing
	// before the goroutine on line 21 has returned anything.
}
