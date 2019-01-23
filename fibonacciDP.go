package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
n  : 0, 1, 2, 3, 4, 5, 6, 7,  ...
fib: 0, 1, 1, 2, 3, 5, 8, 13, ...

This code takes an integer from user input and
calculates the Nth Fibonacci number using 3 approches:

Iteratively with memoization,
Recursively with memoization,
and the naive recursive solution
*/

func main() {
	memo := make(map[int]int)
	n, _ := strconv.Atoi(os.Args[1])

	fmt.Println(recursiveFibDP(n, memo))
	fmt.Println(iterativeFibDP(n))
	fmt.Println(fibNaive(n))
}

// Iteratively return the Nth Fibonacci number utilizing memoization!!
// Time complexity : O(n):
// Space complexity: O(n):
func iterativeFibDP(n int) int {
	var slice = make([]int, n+1)
	slice[0] = 0
	slice[1] = 1
	for i := 2; i <= n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}

	return slice[n]
}

// Recursively returns the Nth Fibonacci number utilizing memoization!!
// Time  complexity: O(n)
// Space complexity: O(n)
func recursiveFibDP(n int, memo map[int]int) int {
	// Check if the value is in the memo
	value, found := memo[n]
	if found {
		return value
	}
	if n <= 1 {
		value = n
	} else {
		value = recursiveFibDP(n-1, memo) + recursiveFibDP(n-2, memo)
	}

	memo[n] = value
	return value
}

// Naive solution:
// Time  Complexity: O(2^n)
// Space Complexity: O(n)
func fibNaive(n int) int {
	if n <= 1 {
		return n
	}
	return fibNaive(n-1) + fibNaive(n-2)
}
