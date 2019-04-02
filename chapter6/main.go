package main

import (
	"fmt"
)

// Arrays!

func main() {
	var x [5]int // Arrays have a size and a given type they are allowed to hold
	x[4] = 100
	fmt.Println(x) // the default int value is 0

	total := 0.0

	testScores := [3]float64{93, 48, 80}

	for _, value := range testScores {
		/* _ is the index variable, value is the value at that index, range
		testScores means "index within the range from 0 to len(testScores)"
		*/

		total += value
	}

	avg := total / float64(len(testScores))
	rounded := fmt.Sprintf("%.f", avg)

	/*
		Rounds to nearest whole number

		For 2 decimal places, use "%.2f", etc.

		For no rounding, just use "%f" or "%1f"
	*/

	fmt.Println(avg)
	fmt.Println("Rounded to...", rounded)

	var sliceAndDice []float64
	/*
		Initializing an array with no length creates a slice.

		Unfortunately, all slices are associated with an underlying array...
		...and we initialized this with the empty array.
	*/

	anotherSlice := make([]float64, 3, 10)
	/*
		Associate a slice with an array via `make`. The slice can be as long as,
		or shorter than, the underlying array. This slice is set at a length of 3,
		but the underlying array is as long as 10.
	*/

	yetanotherSlice := []int{1, 2, 3}
	yetanotherSlice = append(yetanotherSlice, 4, 5, 6)

	/*
		We can add to the length of slices using the append function.
	*/

	sliceAndDice = append(sliceAndDice, 666)

	/*
		We can even do so if we start with an empty array.
	*/

	arr := [5]string{"pepperoni", "vodka", "bbq", "mushroom", "white"}
	pizzaSlice := arr[2:3]
	wholePie := arr[:] // Pythony!

	fmt.Println("SliceAndDice:", sliceAndDice, "AnotherSlice:", anotherSlice, "YetanotherSlice:", yetanotherSlice, "PizzaSlice:", pizzaSlice, "WholePie:", wholePie)

	mappy := make(map[string]int) // maps!

	mappy["Abc"] = 123 // Create

	fmt.Println(mappy["Abc"]) // Read

	mappy["Abc"] = 0 // Update

	delete(mappy, "Abc") // Delete

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	for _, value := range elements {
		fmt.Println(value)
	}

	rarr := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := rarr[0]

	for _, value := range rarr {
		if value < smallest {
			smallest = value
		}
	}

	fmt.Println(smallest)
}
