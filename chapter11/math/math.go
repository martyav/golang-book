package math

// Average is a function for computing a mathematical average
// Since it's capitalized, it's an export
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
