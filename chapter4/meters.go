package main

// FeetToMeters Every function outside main.go must follow a very strict format
// where the line immediately previous to the function's signature
// is a // style comment that starts with the function's name, then some text,
// and continues for however many lines...a /* */ style comment won't do.
// Neither will a comment that starts with anything other than the
// function's name, a comment that's separated from the function
// by a space, or a comment with a newline following the function name.
func FeetToMeters(units float64) float64 {
	return units * 0.3048
}

/*
	Capitalized == Exported method
*/
