// Package simhash hashes []float64 using the simhash data sketching algorithm.
// simhash is described in the following paper (https://www.cs.princeton.edu/courses/archive/spring04/cos598B/bib/CharikarEstim.pdf).
package simhash

import "math/rand"

// Online builds simhash data sketches in an online fashion
type Online struct {
	hyperplanes *[]Hyperplane
}

// NewOnline constructs a simhash.Online builder given the number of hyperplanes
// to construct and the dimension of the input vectors
func NewOnline(hyperplaneCount, dim uint) Online {
	hyperplanes := NewHyperplanes(hyperplaneCount, dim)
	return Online{
		hyperplanes: &hyperplanes,
	}
}

// Hash constructs a simhash data sketch of the given vector
func (o Online) Hash(vector *[]float64) *[]Bit {
	return NewSimhash(o.hyperplanes, vector)
}

// Offline sketches each vector in an offline fashion
func Offline(vectors *[][]float64, hyperplaneCount uint) *[][]Bit {
	simhashs := make([][]Bit, len(*vectors))
	hyperplanes := NewHyperplanes(hyperplaneCount, uint(len((*vectors)[0])))

	for i, vector := range *vectors {
		simhashs[i] = *NewSimhash(&hyperplanes, &vector)
	}

	return &simhashs
}

// Bit represents a bit in an element's hash bit array
type Bit uint8

// NewSimhash constructs a simhash data sketch of the given vector
func NewSimhash(hyperplanes *[]Hyperplane, vector *[]float64) *[]Bit {
	simhash := make([]Bit, len(*hyperplanes))

	for i, hyperplane := range *hyperplanes {
		// The dot product of `hyperplanes[i]` and `vector`
		var dotProduct float64
		for j, v := range *vector {
			dotProduct += hyperplane[j] * v
		}

		if dotProduct >= 0 {
			simhash[i] = Bit(1)
		} else {
			simhash[i] = Bit(0)
		}
	}

	return &simhash
}

// Hyperplane is a `n` dimensional hyperplane
type Hyperplane []float64

// NewHyperplanes constructs `count` hyperplanes, each with dimension `dim`
func NewHyperplanes(count, dim uint) []Hyperplane {
	hyperplanes := make([]Hyperplane, count)

	for i := uint(0); i < count; i++ {
		// Initialize the current hyperplan
		hyperplane := make(Hyperplane, dim)

		// Assign a random number for each dimension of the hyperplane
		for j := uint(0); j < dim; j++ {
			hyperplane[j] = rand.NormFloat64()
		}

		hyperplanes[i] = hyperplane
	}

	return hyperplanes
}
