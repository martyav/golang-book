package main

import "fmt"

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

func main() {
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

	fmt.Println("2 + 2 =", add(2, 2))

	fmt.Println(numerology(10001))
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
