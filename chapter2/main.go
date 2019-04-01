package main

/*
Text below from https://golang.org/doc/code.html#PackageNames

The first statement in a Go source file must be

`package <name>`

where <name> is the package's default name for imports. (All files in a package
must use the same name.)

Go's convention is that the package name is the last element of the
import path: the package imported as "crypto/rot13" should be named `rot13`.

Executable commands must always use `package main`.
*/

import (
	"bufio" // Buffered I/O
	"fmt"   // Formatted I/O operations, similar to C's printf and scanf
	"os"    // Platform independent (tho Unix-like) os interface
)

func main() {
	var greet string // Assignments start with var; type follows the name
	greet = "Hello, world!"
	fmt.Println(greet)

	/*
		Semicolons after statements are optional; the compiler inserts them
		automatically after line breaks that immediately follow certain symbols
		or keywords -- for example, `continue`, string literals, or ().

		Because the compiler inserts semicolons automatically after line
		breaks, you cannot do C#-style "Put the brackets on the next line",
		because the compiler will insert a semicolon before the bracket!

		func() {

		}

		NOT

		func()
		{

		}
	*/

	var scanner = bufio.NewReader(os.Stdin)

	fmt.Println("What is your name?")
	text, _ := scanner.ReadString('\n')

	/*
		Multiple assignment on the same line
		Implicit typing with := (only available within a function)
		'\n' is the character to be used as a delimiter by the scanner
	*/

	fmt.Println("Hello, " + text)
}
