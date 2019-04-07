package math

import "testing"

// Go has a builtin package for testing. Woooo!
// Run this with `go test`
func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}

// From https://www.golang-book.com/books/intro/12
//
// Tests are identified by starting a function with the word Test and taking
// one argument of type *testing.T.
