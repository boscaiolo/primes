package main

import "math"

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func NextPrime(x int) int {
	if x < 2  {
		return 2
	}

	next:=x+1
	nextSqr:=next*next
	for ; next<=nextSqr && !IsPrime(next);  {
		next++
	}

	return next
}

func main() {

	
}
