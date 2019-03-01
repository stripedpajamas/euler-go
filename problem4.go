package main

import (
	"math"
)

func isPalindrome(n int) bool {
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

func PalindromeProduct(upper, lower int) int {
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
