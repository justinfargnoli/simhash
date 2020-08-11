package simhash

import "math/rand"

// Hyperplane is a dim dimensinoal hyperplane
type Hyperplane []float64

// NewHyperplanes constusts a hyperplane given number of hyperplanes to 
// construct and the dimension of each hyperplane
func NewHyperplanes(count, dim uint) []Hyperplane {
	hyperplanes := make([]Hyperplane, count)
	for i := uint(0); i < count; i++ {
		hyperplane := make(Hyperplane, dim)
		for j := uint(0); j < dim; j++ {
			hyperplane[j] = rand.NormFloat64()
		}
		hyperplanes[i] = hyperplane
	}
	return hyperplanes
}