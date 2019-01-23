package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
n  : 0, 1, 2, 3, 4, 5, 6, 7,  ...
fib: 0, 1, 1, 2, 3, 5, 8, 13, ...
*/

func main() {
	memo := make(map[int]int)
	n, _ := strconv.Atoi(os.Args[1])

	//fmt.Println(fibNaive(n))
	fmt.Println(fibDP(n, memo))
}

// Recursively returns the Nth Fibonacci number utilizing memoization!!
// Time  complexity: O(n)
// Space complexity: O(n)
func fibDP(n int, memo map[int]int) int {
	// Check if the value is in the memo
	value, found := memo[n]
	if found {
		return value
	}
	if n <= 1 {
		value = n
	} else {
		value = fibDP(n-1, memo) + fibDP(n-2, memo)
	}

	memo[n] = value
	return value
}

// Naive solution
// Time  Complexity: O(2^n)
// Space Complexity: O(n)
func fibNaive(n int) int {
	if n <= 1 {
		return n
	}
	return fibNaive(n-1) + fibNaive(n-2)
}
