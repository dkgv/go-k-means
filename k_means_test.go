package go_k_means

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestKMeans_Cluster(t *testing.T) {
	type args struct {
		data []Datum
		k    int
	}
	tests := []struct {
		name string
		conf []Option
		args args
		want [][]Datum
	}{
		{
			name: "kpp, 1 point 1 cluster",
			conf: []Option{WithKMeansPlusPlusCentroidStrategy(0)},
			args: args{
				data: []Datum{{Vector: Vector{1, 1}}},
				k:    1,
			},
			want: [][]Datum{
				{
					{Vector: Vector{1, 1}, cluster: 0},
				},
			},
		},
		{
			name: "random, 1 point 1 cluster",
			conf: []Option{WithRandomCentroidStrategy(0)},
			args: args{
				data: []Datum{{Vector: Vector{1, 1}}},
				k:    1,
			},
			want: [][]Datum{
				{
					{Vector: Vector{1, 1}, cluster: 0},
				},
			},
		},
		{
			name: "kpp, 10 points 4 clusters",
			conf: []Option{
				WithKMeansPlusPlusCentroidStrategy(0),
				WithMaxIterations(1),
			},
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{2, 2}},
					{Vector: Vector{30, 30}},
					{Vector: Vector{31, 31}},
					{Vector: Vector{33, 33}},
					{Vector: Vector{155, 155}},
					{Vector: Vector{100, 100}},
					{Vector: Vector{40, 40}},
					{Vector: Vector{5, 5}},
					{Vector: Vector{3333, 3333}},
				},
				k: 4,
			},
			want: [][]Datum{
				{
					{Vector: Vector{30, 30}, cluster: 0},
					{Vector: Vector{31, 31}, cluster: 0},
					{Vector: Vector{33, 33}, cluster: 0},
					{Vector: Vector{40, 40}, cluster: 0},
				},
				{
					{Vector: Vector{3333, 3333}, cluster: 1},
				},
				{
					{Vector: Vector{155, 155}, cluster: 2},
					{Vector: Vector{100, 100}, cluster: 2},
				},
				{
					{Vector: Vector{1, 1}, cluster: 3},
					{Vector: Vector{2, 2}, cluster: 3},
					{Vector: Vector{5, 5}, cluster: 3},
				},
			},
		},
		{
			name: "kpp, 4 points 2 clusters",
			conf: []Option{WithKMeansPlusPlusCentroidStrategy(0)},
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{200, 200}},
					{Vector: Vector{30, 30}},
					{Vector: Vector{31, 31}},
				},
				k: 2,
			},
			want: [][]Datum{
				{
					{Vector: Vector{1, 1}, cluster: 0},
					{Vector: Vector{30, 30}, cluster: 0},
					{Vector: Vector{31, 31}, cluster: 0},
				},
				{
					{Vector: Vector{200, 200}, cluster: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.conf...)
			got := m.Cluster(tt.args.data, tt.args.k)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cluster() got = %v, wantK %v", got, tt.want)
			}
		})
	}
}

func TestKMeans_Fit(t *testing.T) {
	type args struct {
		data []Datum
		maxK int
	}
	tests := []struct {
		name  string
		args  args
		wantK int
	}{
		{
			name: "k=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{2, 2}},
				},
				maxK: 4,
			},
			wantK: 2,
		},
		{
			name: "k=3",
			args: args{
				data: flatten(
					generateData(500, 600, 500, 600, 10),
					generateData(5000, 5100, 5000, 5100, 10),
					generateData(9000, 9100, 9000, 9100, 10),
				),
				maxK: 5,
			},
			wantK: 3,
		},
		{
			name: "k=8",
			args: args{
				data: flatten(
					generateData(1000, 1100, 1000, 1100, 10),
					generateData(2000, 2100, 2000, 2100, 10),
					generateData(3000, 3100, 3000, 3100, 10),
					generateData(4000, 4100, 4000, 4100, 10),
					generateData(5000, 5100, 5000, 5100, 10),
					generateData(6000, 6100, 6000, 6100, 10),
					generateData(7000, 7100, 7000, 7100, 10),
					generateData(8000, 8100, 8000, 8100, 10),
					generateData(9000, 9100, 9000, 9100, 10),
				),
				maxK: 10,
			},
			wantK: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(WithKMeansPlusPlusCentroidStrategy(0))
			if got := m.Fit(tt.args.data, tt.args.maxK); len(got) != tt.wantK {
				t.Errorf("Fit() = %v, wantK %v", len(got), tt.wantK)
			}
		})
	}
}

func flatten(data ...[]Datum) []Datum {
	var result []Datum
	for i := range data {
		result = append(result, data[i]...)
	}
	return result
}

func generateData(minX, maxX, minY, maxY, num int) []Datum {
	data := make([]Datum, num)
	for i := 0; i < num; i++ {
		data[i] = Datum{
			Vector: Vector{float64(minX + rand.Intn(maxX-minX)), float64(minY + rand.Intn(maxY-minY))},
		}
	}
	return data
}
