package go_k_means

import (
	"reflect"
	"testing"
)

func TestPoint_Add(t *testing.T) {
	type args struct {
		other Vector
	}
	tests := []struct {
		name string
		v    Vector
		args args
		want Vector
	}{
		{
			name: "add same length",
			v:    Vector{1, 2, 3},
			args: args{
				other: Vector{1, 2, 3},
			},
			want: Vector{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, wantK %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Distance(t *testing.T) {
	type args struct {
		other Vector
	}
	tests := []struct {
		name string
		v    Vector
		args args
		want float64
	}{
		{
			name: "distance 0",
			v:    Vector{1, 2, 3},
			args: args{
				other: Vector{1, 2, 3},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Distance(tt.args.other)
			if got != tt.want {
				t.Errorf("Distance() got = %v, wantK %v", got, tt.want)
			}
		})
	}
}

func TestPoint_Div(t *testing.T) {
	type args struct {
		scalar float64
	}
	tests := []struct {
		name string
		v    Vector
		args args
		want Vector
	}{
		{
			name: "divide by 2 scalar",
			v:    Vector{1, 2, 3},
			args: args{
				scalar: 2,
			},
			want: Vector{0.5, 1, 1.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Div(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Div() = %v, wantK %v", got, tt.want)
			}
		})
	}
}
