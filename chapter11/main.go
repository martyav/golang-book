package main

import "fmt"
import m "golang-book/chapter11/math" // Directory must be in $GOPATH

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
}

// Inside ./math, run `go install` to create a linkable object file
// Then, go to ./chapter11 and run `go run main.go`
