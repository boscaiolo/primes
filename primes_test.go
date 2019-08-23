package main

import "testing"

func TestNextPrime1(t *testing.T) {
	type args struct {
		x int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{name: "test -1", args: args{x:-1}, want: 2},
		{name: "test 0", args: args{x:0}, want: 2},
		{name: "test 1", args: args{x:1}, want: 2},
		{name: "test 2", args: args{x:2}, want: 3},
		{name: "test 3", args: args{x:3}, want: 5},
		{name: "test 4", args: args{x:4}, want: 5},
		{name: "test 5", args: args{x:5}, want: 7},
		{name: "test 6", args: args{x:6}, want: 7},
		{name: "test 7", args: args{x:7}, want: 11},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.value); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}