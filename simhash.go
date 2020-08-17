package simhash

import "math/rand"

// Online builds simhash data sketches online
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
func (o Online) Hash(vector []float64) Simhash {
	return NewSimhash(*o.hyperplanes, vector)
}

// Offline sketches each vector in an offline fashion
func Offline(vectors [][]float64, hyperplaneCount uint) []Simhash {
	simhashs := make([]Simhash, len(vectors))
	hyperplanes := NewHyperplanes(hyperplaneCount, uint(len(vectors[0])))
	for i, vector := range vectors {
		simhashs[i] = NewSimhash(hyperplanes, vector)
	}
	return simhashs
}

// Simhash is the simhash data sketch of an []float64
type Simhash []uint8

// NewSimhash constructs a simhash data sketch of the given vector
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
