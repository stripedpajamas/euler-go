package main

// https://en.wikipedia.org/wiki/Euclidean_algorithm
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// https://en.wikipedia.org/wiki/Least_common_multiple
func lcm(a, b int) int {
	return (a / gcd(a, b)) * b
}

func multiLcm(nums ...int) int {
	l := 1
	for _, num := range nums {
		l = lcm(l, num)
	}
	return l
}

func SmallestMultiple(lower, upper int) int {
	nums := make([]int, upper-lower)
	for i := lower; i < upper; i++ {
		nums[i-lower] = i
	}
	return multiLcm(nums...)
}
