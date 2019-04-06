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

// Go has channels, which allow 2 goroutines to sync up.
// The functions signature below contain a channel named `c`, which
// transmits strings.

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // Send "ping" via `c`
	}
}

func printer(c chan string) {
	for {
		msg := <-c // Receive and store the value in `c`
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

/*
	From https://www.golang-book.com/books/intro/10

	Using a channel like this synchronizes the two goroutines. When pinger
	attempts to send a message on the channel it will wait until printer is
	ready to receive the message. (this is known as blocking)
*/

func ponger(d chan<- string) { // chan<- specifies directionality; can only send
	for i := 0; ; i++ {
		d <- "pong"
	}
}

func pronter(d <-chan string) { // <-chan specifies can only receive
	for {
		msg := <-d // Receive and store the value in `c`
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
