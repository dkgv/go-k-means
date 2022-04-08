package go_k_means

import "math"

func Score(clusters [][]Datum) float64 {
	si := 0.0 // Silhouette coefficient
	denominator := 0.0
	for _, cluster := range clusters {
		for _, datum := range cluster {
			ai := avgDistance(cluster, datum)                         // Measure of fit in own cluster
			bi := avgDistanceToNearestForeignCluster(clusters, datum) // Dissimilarity measure to neighboring cluster
			si += (bi - ai) / math.Max(ai, bi)
			denominator++
		}
	}

	return si / denominator
}
