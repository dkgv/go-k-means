package go_k_means

import "testing"

func Test_nearestDataToDatum(t *testing.T) {
	type args struct {
		data   []Datum
		target Datum
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 Datum
	}{
		{
			name: "n=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
				},
				target: Datum{Vector: Vector{1, 1}},
			},
			want: 0,
		},
		{
			name: "n=2",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{10, 10}},
				},
				target: Datum{Vector: Vector{10, 10}},
			},
			want: 1,
		},
		{
			name: "n=3",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{2, 2}},
					{Vector: Vector{10, 10}},
				},
				target: Datum{Vector: Vector{10, 10}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nearestDataToDatum(tt.args.data, tt.args.target)
			if got != tt.want {
				t.Errorf("nearestDataToDatum() got = %v, wantK %v", got, tt.want)
			}
		})
	}
}

func Test_avgDistance(t *testing.T) {
	type args struct {
		data   []Datum
		target Datum
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "n=1",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
				},
				target: Datum{Vector: Vector{1, 1}},
			},
			want: 0,
		},
		{
			name: "n=2",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{10, 10}},
				},
				target: Datum{Vector: Vector{10, 10}},
			},
			want: 6.363961030678928,
		},
		{
			name: "n=3",
			args: args{
				data: []Datum{
					{Vector: Vector{1, 1}},
					{Vector: Vector{2, 2}},
					{Vector: Vector{10, 10}},
				},
				target: Datum{Vector: Vector{10, 10}},
			},
			want: 8.01387685344754,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avgDistance(tt.args.data, tt.args.target); got != tt.want {
				t.Errorf("avgDistance() = %v, wantK %v", got, tt.want)
			}
		})
	}
}

func Test_avgDistanceToNearestForeignCluster(t *testing.T) {
	type args struct {
		clusters [][]Datum
		target   Datum
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "n=1",
			args: args{
				clusters: [][]Datum{
					{
						{Vector: Vector{1, 1}, cluster: 1},
					},
				},
				target: Datum{Vector: Vector{1, 1}, cluster: 1},
			},
			want: 0,
		},
		{
			name: "n=2",
			args: args{
				clusters: [][]Datum{
					{
						{Vector: Vector{1, 1}, cluster: 0},
					},
					{
						{Vector: Vector{10, 10}, cluster: 1},
					},
				},
				target: Datum{Vector: Vector{1, 1}, cluster: 0},
			},
			want: 12.727922061357855,
		},
		{
			name: "n=3",
			args: args{
				clusters: [][]Datum{
					{
						{Vector: Vector{1, 1}, cluster: 0},
					},
					{
						{Vector: Vector{2, 2}, cluster: 1},
						{Vector: Vector{10, 10}, cluster: 1},
					},
				},
				target: Datum{Vector: Vector{1, 1}, cluster: 0},
			},
			want: 7.0710678118654755,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avgDistanceToNearestForeignCluster(tt.args.clusters, tt.args.target); got != tt.want {
				t.Errorf("avgDistanceToNearestForeignCluster() = %v, wantK %v", got, tt.want)
			}
		})
	}
}
