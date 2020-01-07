package main

import (
	"fmt"
	"math"
)

func main() {
	x, y, limit := 3, 5, 1000
	answer := multiples(x, y, limit)
	fmt.Printf("Problem 1:\n\tSum of multiples of %d and %d < %d is %d\n", x, y, limit, answer)

	limit = 4000000
	answer = evenFibonacci(limit)
	fmt.Printf("Problem 2:\n\tSum of even Fibonacci numbers <= %d is %d\n", limit, answer)

	n := 600851475143
	answer = largestPrimeFactor(n)
	fmt.Printf("Problem 3:\n\tLargest prime factor of %d is %d\n", n, answer)

	lower := 100
	upper := 999
	answer = palindromeProduct(lower, upper)
	fmt.Printf("Problem 4:\n\tLargest palindromic product of two 3-digit numbers: %d\n", answer)

	lower = 1
	upper = 20
	answer = smallestMultiple(lower, upper)
	fmt.Printf("Problem 5:\n\tSmallest multiple that divides [1, 20]: %d\n", answer)
}

func multiples(x, y, limit int) int {
	sum := 0
	for i := 3; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func evenFibonacci(limit int) int {
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

func largestPrimeFactor(n int) int {
	var factorize func(n, factor, largest, limit int) int
	var removeFactor func(n, factor int) int

	removeFactor = func(n, factor int) int {
		for n%factor == 0 {
			n /= factor
		}
		return n
	}

	factorize = func(n, factor, largest, limit int) int {
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

	return factorize(removeFactor(n, 2), 3, 2, int(math.Ceil(math.Sqrt(float64(n)))))
}

func palindromeProduct(lower, upper int) int {
	var isPalindrome func(n int) bool

	isPalindrome = func(n int) bool {
		if n < 10 {
			return true
		}
		high := int(math.Pow(10, math.Floor(math.Log10(float64(n)))+1))
		low := 10
		for high > low {
			x := (n % high) / (high / 10)
			y := (n % low) / (low / 10)
			if x != y {
				return false
			}
			high /= 10
			low *= 10
		}
		return true
	}

	max := 0
	for i := upper; i >= lower; i-- {
		for j := i - 1; j >= lower; j-- {
			if i*j < max {
				break
			}
			if isPalindrome(i * j) {
				max = i * j
			}
		}
	}
	return max
}

func smallestMultiple(lower, upper int) int {
	var gcd func(a, b int) int
	var lcm func(a, b int) int

	// https://en.wikipedia.org/wiki/Euclidean_algorithm
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}

	// https://en.wikipedia.org/wiki/Least_common_multiple
	lcm = func(a, b int) int {
		return (a / gcd(a, b)) * b
	}

	l := 1
	for i := lower; i < upper; i++ {
		l = lcm(l, i)
	}
	return l
}
