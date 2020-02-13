package main

import (
	"fmt"
	"strings"

	"github.com/teohrt/kPalindromes/kpalindromes"
)

// Given an int k, and a string s, write a function that returns a bool
// representing if k number of palindromes can be created from
// permutations of s that utilize ALL characters of s.

func main() {
	k := 3
	s := "aabbaa"
	res := kpalindromes.Kpalindromes(k, s)
	fmt.Println(res)
}

// O(n) time | O(n) space
func Kpalindromes(k int, str string) bool {
	n := len(str)
	if n <= 1 && k <= 1 || k == 0 {
		return true
	}
	arr := strings.Split(str, "")
	charCounts := map[string]int{}
	for _, s := range arr {
		if _, ok := charCounts[s]; ok {
			charCounts[s]++
		} else {
			charCounts[s] = 1
		}
	}
	oddCharacterCount := 0
	for key, val := range charCounts {
		if val%2 != 0 {
			oddCharacterCount++
		}
		charCounts[key] = charCounts[key] / 2
	}
	if oddCharacterCount > 1 && n > 2 {
		return false
	}
	possiblePalendromeCount := 0
	for i := 0; i < n/2; i++ {
		count := (n / 2) - i
		possiblePalendromeCount += count
	}
	return possiblePalendromeCount-oddCharacterCount >= k
}
