package go_k_means

import (
	"math"
)

type Vector []float64

func (v Vector) Distance(other Vector) float64 {
	sum := 0.0
	for i := 0; i < len(v); i++ {
		sum += (v[i] - other[i]) * (v[i] - other[i])
	}
	return math.Sqrt(sum)
}

func (v Vector) Add(other Vector) Vector {
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] + other[i]
	}
	return result
}

func (v Vector) Div(scalar float64) Vector {
	result := make(Vector, len(v))
	for i := range v {
		result[i] = v[i] / scalar
	}
	return result
}
