package main

import (
	"fmt"
)

func main() {
	Problem1()
	Problem2()
	Problem3()
	Problem4()
}

func Problem1() {
	x, y, limit := 3, 5, 1000
	sum := Multiples(x, y, limit)
	fmt.Printf("Problem 1:\n\tSum of multiples of %d and %d < %d is %d\n", x, y, limit, sum)
}

func Problem2() {
	limit := 4000000
	sum := EvenFibonacci(limit)
	fmt.Printf("Problem 2:\n\tSum of even Fibonacci numbers <= %d is %d\n", limit, sum)
}

func Problem3() {
	n := 600851475143
	p := LargestPrimeFactor(n)
	fmt.Printf("Problem 3:\n\tLargest prime factor of %d is %d\n", n, p)
}

func Problem4() {
	upper := 999
	lower := 100
	p := PalindromeProduct(upper, lower)
	fmt.Printf("Problem 4:\n\tLargest palindromic product of two 3-digit numbers: %d\n", p)
}
