package go_k_means

import (
	"math/rand"
)

type RandomCentroidStrategy struct {
	randomSeed int64
}

func (s RandomCentroidStrategy) Select(data []Datum, k int) []Datum {
	rand.Seed(s.randomSeed)
	visited := make([]bool, len(data))
	centroids := make([]Datum, k)
	for i := 0; i < k; {
		index := rand.Intn(len(data))
		if visited[index] {
			continue
		}
		visited[index] = true
		centroids[i] = data[index]
		i++
	}
	return centroids
}

type KMeansPlusPlusCentroidStrategy struct {
	randomSeed int64
}

// https://en.wikipedia.org/wiki/K-means%2B%2B
func (s KMeansPlusPlusCentroidStrategy) Select(data []Datum, k int) []Datum {
	rand.Seed(s.randomSeed)

	centroids := make([]Datum, k)
	centroids[0] = data[rand.Intn(len(data))]

	distances := make([]float64, len(data))
	for i := 1; i < k; i++ {
		sum := 0.0
		for j, datum := range data {
			minDistance := 10e10
			for l := range centroids[:i] {
				distance := datum.Vector.Distance(centroids[l].Vector)
				if distance < minDistance {
					minDistance = distance
				}
			}
			distances[j] = minDistance * minDistance
			sum += distances[j]
		}
		nextIndex := rand.Float64() * sum
		j := 0
		for sum = distances[j]; sum < nextIndex; sum += distances[j] {
			j++
		}
		centroids[i] = data[j]
	}
	return centroids
}
