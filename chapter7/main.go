package main

import "fmt"

func isValueInAndWhere(arr []string, value string) (bool, int) {
	for i, text := range arr {
		if value == text {
			return true, i
		}
	}

	return false, -1
}

/*
	Go allows functions to return multiple things
*/

func main() {
	texts := []string{"bunny", "chick", "calf"}
	shouldBeFalse, shouldBeNegativeOne := isValueInAndWhere(texts, "puppy")
	shouldBeTrue, shouldBeOne := isValueInAndWhere(texts, "chick")

	fmt.Println(shouldBeFalse, shouldBeNegativeOne)
	fmt.Println(shouldBeTrue, shouldBeOne)

	add := func(x, y int) int { return x + y }

	/*
		We can define functions within functions by using closures
	*/

	fmt.Println(add(2, 2))
}
