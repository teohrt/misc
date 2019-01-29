package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

/*
n  : 0, 1, 2, 3, 4, 5, 6, 7,  ...
fib: 0, 1, 1, 2, 3, 5, 8, 13, ...

This code takes an integer from arguments and
calculates the Nth Fibonacci number using 3 approches
and logs their excecution times

The three approaches are:

Iteratively with memoization,
Recursively with memoization,
and the naive recursive solution
*/

type FibFunc func(int) int

// Decerator for logging the excecution time of the Fibonacci algorithms
func excecutionTimeLogger(f FibFunc) FibFunc {
	return func(n int) int {
		defer func(t time.Time) {
			fmt.Printf("%s - runtime: %d\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), time.Since(t))
		}(time.Now())
		return f(n)
	}
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	excecutionTimeLogger(recursiveFibDP)(n)
	excecutionTimeLogger(iterativeFibDP)(n)
	excecutionTimeLogger(fibNaive)(n)
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
func RecursiveFibDP(n int, memo map[int]int) int {
	// Check if the value is in the memo
	value, found := memo[n]
	if found {
		return value
	}
	if n <= 1 {
		value = n
	} else {
		value = RecursiveFibDP(n-1, memo) + RecursiveFibDP(n-2, memo)
	}

	memo[n] = value
	return value
}

// Helper function
func recursiveFibDP(n int) int {
	memo := make(map[int]int)
	return RecursiveFibDP(n, memo)
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
