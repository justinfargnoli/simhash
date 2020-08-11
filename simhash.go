package simhash

import "math/rand"

// Simhash is the simhash data sketch of an []float64
type Simhash []uint8

// NewSimhash constructs a simhash data sketch of the givec vector
func NewSimhash(hyperplanes []Hyperplane, vector []float64) Simhash {
	simhash := make(Simhash, len(hyperplanes))
	for i, hyperplane := range hyperplanes {
		var dotProduct float64 // the dot product of hyperplanes[i] and vector
		for j, v := range vector {
			dotProduct += hyperplane[j] * v
		}
		if dotProduct >= 0 {
			simhash[i] = uint8(1)
		} else {
			simhash[i] = uint8(0)
		}
	}
	return simhash
}

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