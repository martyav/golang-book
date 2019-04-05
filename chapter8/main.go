package main

import "fmt"

/*
Pointers allow us to pass arguments by reference, and alter the
location in memory represented by a variable.
*/

func setToZero(pointer *int) { // pointer type int
	*pointer = 0 // dereference -- access the value currently pointed to
}

func iterate(pointer *int) {
	*pointer++
}

func main() {
	x := 55555

	setToZero(&x) // find the address x points to & return the pointer

	fmt.Println(x)

	/*
		From https://www.golang-book.com/books/intro/8

		In Go a pointer is represented using the `*` (asterisk) character followed
		by the type of the stored value. In the `zero` function `pointer` is a pointer
		to an int.

		`*` is also used to “dereference” pointer variables. Dereferencing a pointer
		gives us access to the value the pointer points to. When we write
		`*pointer = 0` we are saying “store the int 0 in the memory location
		`pointer` refers to”. If we try `pointer = 0` instead, we will get compiler
		error because `pointer` is not an int -- it's a *int, which can only be
		given another *int.

		Finally we use the `&` operator to find the address of a variable. `&x`
		returns a `*int` (pointer to an int) because `x` is an int. This is what
		allows us to modify the original variable. `&x` in main and `pointer` in
		zero refer to the same memory location.
	*/

	newInt := new(int)

	iterate(newInt)

	fmt.Println(*newInt) // x is 1

	iterate(newInt)

	fmt.Println(*newInt) // x is 2

	iterate(newInt)

	fmt.Println(*newInt) // x is 3

	/*
		We can use the `new` keyword to allocate space for a new int in memory.
		This keyword returns a pointer.

		Since we don't give `newInt` a value, the value at `newInt` defaults to 0.

		Note the we don't need to use `&` when passing newInt as an argument --
		it's already a pointer -- however, we do need to call newInt with `*`
		inside `fmt.Println` to dereference it.
	*/

	a := new(int)
	b := 1313

	fmt.Println("A:", *a, "B:", b)

	swap(a, &b)

	fmt.Println("A:", *a, "B:", b)

	/*
		We always need to call a with `*` because the variable is a pointer
	*/
}

// Problems

func square(x *float64) {
	*x = *x * *x
}

func swap(x, y *int) {
	tempY := *y
	tempX := *x

	*y = tempX
	*x = tempY
}
