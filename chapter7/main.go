package main

import (
	"fmt"
	"math"
)

/*
	Go allows functions to return multiple things
*/

func isValueInAndWhere(arr []string, value string) (bool, int) {
	for i, text := range arr {
		if value == text {
			return true, i
		}
	}

	return false, -1
}

/*
	Go also has variadic functions
*/

func variadic(args ...string) int {
	return len(args)
}

/*
	Go has a special keyword for invoking run-time errors: `panic`

	A panic immediately stops the function dead.
*/

func panicAtTheRepo() {
	panic("I write sins not try-catch statements")
}

func main() {
	fmt.Println(fib(7))

	length := variadic("one", "**", "3", "IV", "|||||")

	fmt.Println("The length of all these items is", length)

	texts := []string{"bunny", "chick", "calf"}
	shouldBeFalse, shouldBeNegativeOne := isValueInAndWhere(texts, "puppy")
	shouldBeTrue, shouldBeOne := isValueInAndWhere(texts, "chick")

	fmt.Println("Should be false, should be -1:", shouldBeFalse, shouldBeNegativeOne)
	fmt.Println("Should be true, should be 1:", shouldBeTrue, shouldBeOne)

	/*
		We can define functions within functions by using closures
	*/

	add := func(x, y int) int { return x + y }

	/*
		The `defer` keyword lets us delay when a function is run
	*/

	defer fmt.Println("2 + 2 =", add(2, 2))

	fmt.Println("If we take all the digits in 10001 and add them together, we get", numerology(10001))

	/*
		the `recover` keyword reports back what caused the panic.

		Using `defer` with `recover` allows us to see what's going on.
	*/

	defer func() {
		str := recover()
		fmt.Println(str)
	}() // immediately calls the closure after it is defined

	panicAtTheRepo()
}

/*
	Go also supports recursion.

	Note that Go will not allow you to define a named function inside another
	named function!
*/

func numerology(num int) int {
	if num < 10 {
		return num
	}

	return numerology(num/10) + (num % 10)
}

// Problems

func sum(args ...int) int {
	total := 0

	for _, value := range args {
		total += value
	}

	return total
}

func half(num float64) (float64, bool) {
	halved := num / 2.0

	return math.Floor(halved), halved > 1 && math.Floor(halved) == halved
}

func findGreatest(args ...int) int {
	greatest := 0

	for _, value := range args { // Using len(args) with cause a type error
		if value > greatest {
			greatest = value
		}
	}

	return greatest
}

func makeOddGenerator() func() uint {
	initial := uint(1)

	return func() (iterate uint) {
		iterate = initial

		initial += 2

		return
	}
}

func fib(num int) int {
	if num < 2 {
		return num
	}

	return fib(num-1) + fib(num-2)
}
