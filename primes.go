package main

import "math"
import . "gonum.org/v1/gonum/mat"

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func NextPrime(x int) int {
	if x < 2 {
		return 2
	}

	next := x + 1
	nextSqr := next * next
	for next <= nextSqr && !IsPrime(next) {
		next++
	}

	return next
}

func MultyTable(n int) Dense {

	x := n+1

	r := NewDense(x, x, nil)
	c := NewDense(x, x, nil)
	
	v := make([]float64, x)
	v[0]=1
	for i := 1; i < x; i++ {
		v[i] = float64(NextPrime(int(v[i-1])))
	}

	r.SetRow(0, v)
	c.SetCol(0, v)

	var t Dense
	t.Mul(c,r)
	return t
}

func main() {

}
