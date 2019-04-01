package main

import "fmt"

const doubleIt int = 2

func main() {
	fmt.Println("Give me a number and I'll double it!")

	var input int
	fmt.Scanf("%f", &input)

	output := input * doubleIt

	fmt.Println(output)

	/*
		The two number's types MUST match. Go won't coerce them.

		`%d` stands for base 10 int.
		The corresponding format for floats is `%f`.

		`&` is a reference to a pointer's address. In this case, it's
		a reference to the address of input, allowing us to access what is
		inside at the moment.

		See https://gist.github.com/josephspurrier/7686b139f29601c3b370
		for a pointer cheatsheet.
	*/
}
