package go_k_means

import (
	"time"
)

type Options struct {
	maxIterations    int
	centroidStrategy CentroidStrategy
}

type Option func(*Options)

func WithMaxIterations(maxIterations int) Option {
	return func(options *Options) {
		options.maxIterations = maxIterations
	}
}

func WithRandomCentroidStrategy(seed ...int64) Option {
	return func(options *Options) {
		s := time.Now().UnixMilli()
		if seed != nil {
			s = seed[0]
		}
		options.centroidStrategy = RandomCentroidStrategy{
			randomSeed: s,
		}
	}
}

func WithKMeansPlusPlusCentroidStrategy(seed ...int64) Option {
	return func(options *Options) {
		s := time.Now().UnixMilli()
		if seed != nil {
			s = seed[0]
		}
		options.centroidStrategy = KMeansPlusPlusCentroidStrategy{
			randomSeed: s,
		}
	}
}

type CentroidStrategy interface {
	Select(data []Datum, k int) []Datum
}

type KMeans struct {
	conf Options
}

func DefaultOptions() Options {
	return Options{
		maxIterations: 100,
		centroidStrategy: KMeansPlusPlusCentroidStrategy{
			randomSeed: time.Now().UnixMilli(),
		},
	}
}

func New(opts ...Option) *KMeans {
	conf := DefaultOptions()
	for _, opt := range opts {
		opt(&conf)
	}

	return &KMeans{conf: conf}
}

type Datum struct {
	Vector  Vector
	cluster int
}

func (m KMeans) Fit(data []Datum, maxK int) [][]Datum {
	bestScore := -1.0
	var best [][]Datum
	for k := 2; k < maxK; k++ {
		clusters := m.Cluster(data, k)
		score := Score(clusters)
		if score > bestScore {
			bestScore = score
			best = clusters
		}
	}

	return best
}

func (m KMeans) Cluster(data []Datum, k int) [][]Datum {
	// Pick initial centroids
	centroids := m.conf.centroidStrategy.Select(data, k)
	for i := range centroids {
		centroids[i].cluster = i + 1
	}

	// Assign all unassigned data to nearest centroid
	for i := 0; i < m.conf.maxIterations; i++ {
		for _, datum := range data {
			if datum.cluster != 0 {
				continue
			}
			index := nearestDataToDatum(centroids, datum)
			datum.cluster = index + 1
		}
	}

	for i := 0; i < m.conf.maxIterations; i++ {
		// Adjust centroids
		clusterSizes := make([]int, k)
		for j := range data {
			datum := data[j]
			if datum.cluster == 0 {
				continue
			}
			centroid := centroids[datum.cluster-1].Vector
			centroid = centroid.Add(datum.Vector)
			clusterSizes[datum.cluster-1]++
		}

		for j := range centroids {
			centroid := centroids[j].Vector
			centroid = centroid.Div(float64(clusterSizes[j]))
		}

		// Reassign data to nearest centroid
		changes := 0
		for j := range data {
			index := nearestDataToDatum(centroids, data[j])
			prevIndex := data[j].cluster - 1
			if prevIndex <= 0 || prevIndex != index {
				data[j].cluster = index + 1
				changes++
			}
		}

		if changes == 0 {
			break
		}
	}

	clusters := make([][]Datum, k)
	for i := range data {
		data[i].cluster--
		cluster := data[i].cluster
		clusters[cluster] = append(clusters[cluster], data[i])
	}

	return clusters
}
