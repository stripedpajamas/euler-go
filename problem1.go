package main

func Multiples(x, y, limit int) int {
	sum := 0
	for i := 3; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
