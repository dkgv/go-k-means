package go_k_means

import (
	"reflect"
	"testing"
)

func TestKMeansPlusPlusCentroidStrategy_Select(t *testing.T) {
	type args struct {
		data []Datum
		k    int
	}
	tests := []struct {
		name string
		args args
		want []Datum
	}{
		{
			name: "1 point, k=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
				},
				k: 1,
			},
			want: []Datum{
				{Vector: Vector{1, 2, 3}},
			},
		},
		{
			name: "2 points, k=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
					{Vector: Vector{2, 3, 4}},
				},
				k: 1,
			},
			want: []Datum{
				{Vector: Vector{1, 2, 3}},
			},
		},
		{
			name: "3 points, k=2",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
					{Vector: Vector{2, 3, 4}},
					{Vector: Vector{20, 30, 40}},
				},
				k: 2,
			},
			want: []Datum{
				{Vector: Vector{1, 2, 3}},
				{Vector: Vector{20, 30, 40}},
			},
		},
		{
			name: "4 points, k=3",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
					{Vector: Vector{2, 3, 4}},
					{Vector: Vector{20, 30, 40}},
					{Vector: Vector{200, 300, 400}},
				},
				k: 3,
			},
			want: []Datum{
				{Vector: Vector{20, 30, 40}},
				{Vector: Vector{200, 300, 400}},
				{Vector: Vector{2, 3, 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := KMeansPlusPlusCentroidStrategy{randomSeed: 0}
			if got := s.Select(tt.args.data, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, wantK %v", got, tt.want)
			}
		})
	}
}

func TestRandomCentroidStrategy_Select(t *testing.T) {
	type args struct {
		data []Datum
		k    int
	}
	tests := []struct {
		name string
		args args
		want []Datum
	}{
		{
			name: "1 point, k=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
				},
				k: 1,
			},
			want: []Datum{
				{Vector: Vector{1, 2, 3}},
			},
		},
		{
			name: "2 points, k=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
					{Vector: Vector{2, 3, 4}},
				},
				k: 1,
			},
			want: []Datum{
				{Vector: Vector{1, 2, 3}},
			},
		},
		{
			name: "3 points, k=2",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 2, 3}},
					{Vector: Vector{2, 3, 4}},
					{Vector: Vector{20, 30, 40}},
				},
				k: 2,
			},
			want: []Datum{
				{Vector: Vector{1, 2, 3}},
				{Vector: Vector{2, 3, 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RandomCentroidStrategy{randomSeed: 0}
			if got := s.Select(tt.args.data, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, wantK %v", got, tt.want)
			}
		})
	}
}
