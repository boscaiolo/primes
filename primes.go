package main

import (
	"fmt"
	"math"
)
import . "gonum.org/v1/gonum/mat"

func primeVector(n int) []float64 {

	lim := n + 1

	for lim <= (n+1)*int(math.Log(float64(lim)))  {
		lim++
	}

	f := make([]bool, lim)

	for i := 2; i <= int(math.Sqrt(float64(lim))); i++ {
		if f[i] == false {
			for j := i * i; j < lim; j += i {
				f[j] = true
			}
		}
	}

	p := make([]float64, n+1)
	p[0] = 1
	j := 1
	i := 2
	for  i < lim && j <= n {
		if f[i] == false {
			p[j] = float64(i)
			j++
		}
		i++
	}
	return p
}

func MultyTable(n int) Dense {

	x := n + 1

	r := NewDense(x, x, nil)
	c := NewDense(x, x, nil)

	var primes = primeVector(n + 1)

	r.SetRow(0, primes)
	c.SetCol(0, primes)

	var t Dense
	t.Mul(c, r)
	return t
}

func main() {
	var n int
	fmt.Println("Enter an integer value >=1 : ")
	_, err := fmt.Scanf("%d", &n)

	if err != nil {
		fmt.Println(err)
	} else if n >= 1 {

		var c = MultyTable(n)

		fc := Formatted(&c, Squeeze())
		fmt.Printf("%v", fc)
	}
}
