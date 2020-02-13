package main

import (
	"fmt"
	"strings"
)

func getPermutations(s string) (perms []string) {
	array := strings.Split(s, "")
	permutationsHelper(0, array, &perms)
	return perms
}

func permutationsHelper(i int, array []string, perms *[]string) {
	if i == len(array)-1 {
		newPerm := strings.Join(array, "")
		*perms = append(*perms, newPerm)
	} else {
		for j := i; j < len(array); j++ {
			swap(array, i, j)
			permutationsHelper(i+1, array, perms)
			swap(array, i, j)
		}
	}
}

func swap(array []string, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func printPerms(perms []string) {
	for i := 0; i < len(perms); i++ {
		fmt.Println(perms[i])
	}
}

func main() {
	s := "Trace"
	perms := getPermutations(s)
	printPerms(perms)
}
