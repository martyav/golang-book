package main

import (
	"fmt"
	"math"
)

// Circle takes the stage: Go has structs
type Circle struct { // `type` keyword indicates a new type definition
	x float64
	y float64
	r float64
}

func main() {
	// 4 ways to make a struct

	var a Circle // This sets every field to 0, "", or nil

	b := new(Circle) // Same as above, plus returns a *pointer

	c := Circle{x: 5, y: 5, r: 10} // Usual initialization

	d := Circle{25, 25, 14} // Shorthand with positional params

	fmt.Println(a.x, b.y, c.r, d.x) // Fields are accessed with dot notation
}

// Methods have a special syntax, with a receiver in the signature.
// The receiver syntax allows us to attach a fuction to a struct.
func (c *Circle) area() float64 { // the receiver is `(c *Circle)`
	return math.Pi * c.r * c.r
}

// Go has interfaces
type Carnivore interface {
	hunting() bool
}

// We can define fields on one line if they all share a single type
type Felid struct {
	species, genus, scratch, bite string
}

// Because Felid has a hunting method now, it implements `Carnivore`
func (f *Felid) hunting(hungry, preyAvailable bool) bool {
	return hungry && preyAvailable
}

// Embedded types are anonymous objects that allow us to call
// one struct's fields and methods on another
type Lion struct {
	Felid
	roar string
}
