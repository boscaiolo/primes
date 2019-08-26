package main

import (
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

func TestNextPrime1(t *testing.T) {
	type args struct {
		x int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{name: "test -1", args: args{x: -1}, want: 2},
		{name: "test 0", args: args{x: 0}, want: 2},
		{name: "test 1", args: args{x: 1}, want: 2},
		{name: "test 2", args: args{x: 2}, want: 3},
		{name: "test 3", args: args{x: 3}, want: 5},
		{name: "test 4", args: args{x: 4}, want: 5},
		{name: "test 5", args: args{x: 5}, want: 7},
		{name: "test 6", args: args{x: 6}, want: 7},
		{name: "test 7", args: args{x: 7}, want: 11},
		{name: "test 7", args: args{x: 11}, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NextPrime(tt.args.x); got != tt.want {
				t.Errorf("NextPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test -1", args: args{-1}, want: false},
		{name: "test 0", args: args{0}, want: false},
		{name: "test 1", args: args{1}, want: false},
		{name: "test 2", args: args{2}, want: true},
		{name: "test 3", args: args{3}, want: true},
		{name: "test 4", args: args{4}, want: false},
		{name: "test 5", args: args{5}, want: true},
		{name: "test 6", args: args{6}, want: false},
		{name: "test 7", args: args{7}, want: true},
		{name: "test 7", args: args{8}, want: false},
		{name: "test 7", args: args{11}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.value); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultyTable(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want mat.Dense
	}{
		{"test n=1",args{1},*mat.NewDense(2, 2, []float64{1, 2, 2, 4})},
		{"test n=2",args{2},*mat.NewDense(3, 3, []float64{1, 2, 3, 2, 4, 6, 3, 6, 9})},
		{"test n=3",args{3},
		*mat.NewDense(4, 4,
				[]float64{1, 2, 3, 5,
			  			  2, 4, 6, 10,
			  			  3, 6, 9, 15,
			  			  5, 10, 15, 25})},
		{"test n=4",args{4},
		*mat.NewDense(5, 5,
				[]float64{1, 2, 3, 5, 7,
			  			  2, 4, 6, 10, 14,
			  			  3, 6, 9, 15, 21,
			  			  5, 10, 15, 25, 35,
			  			  7, 14, 21, 35, 49})},
		{"test n=5",args{5},
		*mat.NewDense(6, 6,
				[]float64{1, 2, 3, 5, 7, 11,
			  			  2, 4, 6, 10, 14, 22,
			  			  3, 6, 9, 15, 21, 33,
			  			  5, 10, 15, 25, 35, 55,
			  			  7, 14, 21, 35, 49, 77,
			  			  11, 22, 33, 55, 77, 121})},
		{"test n=6",args{6},
		*mat.NewDense(7, 7,
				[]float64{1, 2, 3, 5, 7, 11, 13,
			  			  2, 4, 6, 10, 14, 22, 26,
			  			  3, 6, 9, 15, 21, 33, 39,
			  			  5, 10, 15, 25, 35, 55, 65,
			  			  7, 14, 21, 35, 49, 77, 91,
			  			  11, 22, 33, 55, 77, 121, 143,
						  13, 26, 39, 65, 91, 143, 169})},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultyTable(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultyTable() = %v, want %v", got, tt.want)
			}
		})
	}
}