package main

import (
	"fmt"
)

func main() {
	var one int = 2 // When assigning a new whole number, type int is implicit
	var two int8 = 4
	var three int16 = 8
	var four int32 = 16
	var five int64 = 32

	fmt.Println("The powers of 2 are ", one, two, three, four, five)

	var onePointOh float32 = 1.5
	var doubleOh float64 = 0.07 // When assigning a new float, float64 is implicit

	fmt.Println("Go has two different kinds of floats, but they aren't called shorts and longs or doubles.", onePointOh, "is a float32, and", doubleOh, "is a float64.")

	/*
		We can print longer text simply by separating values with commas. We don't
		need to add spaces within the text, because spaces will be added between
		each comma-separated value.
	*/

	fmt.Println("1 + 1 * 1 / 1 - 2 = ", 1+1*1/1-2)
	fmt.Println("1 + 1 * 1 / 1.0 - 2.0 = ", 1+1*1/1.0-2.0)

	word := "word"

	fmt.Println("The length of `word` is", len(word)) // Python 2-style, eh?
	fmt.Println("the first letter of `word` is", word[0])
	fmt.Println("Wait, what?")

	fmt.Println("True and true is", true && true)
	fmt.Println("False and false is, naturally,", false && false)
	fmt.Println("False and true is", false && true)
	fmt.Println("...but false or true is", false || true)
	fmt.Println("And not false is", !false)
	fmt.Println("While not true is", !true)

	/*
		True and false in Go are untyped Boolean constants.

		From https://blog.golang.org/constants

		 An untyped constant is just a value, one not yet given a defined type that
		 would force it to obey the strict rules that prevent combining differently
		 typed values.
	*/
}
