package main

/*
This one uses functions separated out into different files.

To run this script, you need to specify the two other files:

go run main.go celsius.go meters.go

or...

go run *.go
*/

import (
	"fmt"
)

const doubleIt int = 2

func main() {
	fmt.Println("Give me a number and I'll double it!")

	var input int
	fmt.Scanf("%f", &input)

	/*
		The first argument is the kind of data to look for while
		scanning from standard input.

		`%d` stands for base 10 int.
		The corresponding format for floats is `%f`.

		The second (and any successive arguments) is where to store the
		input.

		`&` is a refererence, allowing us to access what is
		at that pointer's address.

		See https://gist.github.com/josephspurrier/7686b139f29601c3b370
		for a pointer cheatsheet.

		Also: The two numbers' types MUST match. Go won't coerce them.
	*/

	output := input * doubleIt

	fmt.Println(output)

	var flinput float64

	fmt.Println("Give me a temperature in Fahrenheit: ")

	fmt.Scanf("%f", &flinput)

	floutput := celsiusToFahrenheit(flinput)

	fmt.Println("That's", floutput, "in Celsius.")

	fmt.Println("Now give me a measurement in feet: ")

	fmt.Scanf("%f", &flinput)

	floutput = FeetToMeters(flinput)

	fmt.Println("That's", floutput, "in meters.")
}
