# go-k-means
[![Coverage Status](https://coveralls.io/repos/github/dkgv/go-k-means/badge.svg?branch=master)](https://coveralls.io/github/dkgv/go-k-means?branch=master)

A simple Go k-means clustering library with support for [k-means++](https://en.wikipedia.org/wiki/K-means%2B%2B) centroid initialization and k value evaluation via [silhouette scoring](https://en.wikipedia.org/wiki/Silhouette_(clustering)).

## Usage

```go
package main

import k_means "github.com/dkgv/go-k-means"

func main() {
    data := []k_means.Datum{
        // ...
    }
	
    randomSeed := 0
    km := k_means.New(
        k_means.WithMinIterations(1)
        k_means.WithMaxIterations(100),
        k_means.WithKMeansPlusPlusCentroidStrategy(randomSeed),
        // k_means.WithRandomCentroidStrategy(randomSeed),
    )
	
    // Find k value and return clustering
    maxK := 5
    clusters := km.Fit(data, maxK)

    // Score clustering
    score := k_means.Score(clusters)	
	
    // Cluster data using a specific k value
    k := 3
    clusters := km.Cluster(data, k)
}
```
