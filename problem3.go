package main

import "math"

func LargestPrimeFactor(n int) int {
	return factorize(removeFactor(n, 2), 3, 2, int(math.Ceil(math.Sqrt(float64(n)))))
}

func removeFactor(n, factor int) int {
	for n%factor == 0 {
		n /= factor
	}
	return n
}

func factorize(n, factor, largest, limit int) int {
	if factor > limit {
		return largest
	}
	nextN, nextFactor, nextLargest := n, factor+2, largest

	if n%factor == 0 {
		nextN = removeFactor(n, factor)
		if factor > largest {
			nextLargest = factor
		}
	}
	return factorize(nextN, nextFactor, nextLargest, limit)
}
