package main

func EvenFibonacci(limit int) int {
	a := 1
	b := 2
	c := a + b
	sum := 2
	for c <= limit {
		if c%2 == 0 {
			sum += c
		}
		a, b, c = b, c, b+c
	}
	return sum
}
