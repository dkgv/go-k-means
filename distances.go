package go_k_means

func nearestDataToDatum(data []Datum, target Datum) int {
	nearestIndex := 0
	nearestDatum := data[nearestIndex]
	minDistance := nearestDatum.Vector.Distance(target.Vector)
	for i := range data[1:] {
		distance := data[i+1].Vector.Distance(target.Vector)
		if distance < minDistance {
			minDistance = distance
			nearestIndex = i + 1
			nearestDatum = data[i]
		}
	}

	return nearestIndex
}

func avgDistance(data []Datum, target Datum) float64 {
	var sum float64
	for _, datum := range data {
		sum += datum.Vector.Distance(target.Vector)
	}

	return sum / float64(len(data))
}

func avgDistanceToNearestForeignCluster(clusters [][]Datum, target Datum) float64 {
	distance := 1e10
	for i, cluster := range clusters {
		if i == target.cluster {
			continue
		}

		foreignDistance := avgDistance(cluster, target)
		if foreignDistance < distance {
			distance = foreignDistance
		}
	}

	return distance
}
