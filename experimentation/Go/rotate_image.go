package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	printImage(arr)
	fmt.Println("-----------")
	result := rotateImage(arr)
	printImage(result)
}

func printImage(image [][]int) {
	length := len(image[0])
	for y := 0; y < length; y++ {
		for x := 0; x < length; x++ {
			e := image[y][x]
			fmt.Printf("%d ", e)
		}
		fmt.Print("\n")
	}
}

func rotateImage(a [][]int) [][]int {
	n := len(a[0])
	for y := 0; y < n; y++ {
		for x := y; x < ((n - y) - 1); x++ {
			rotatePixel(y, x, &a)
		}
	}
	return a
}

func rotatePixel(y int, x int, a *([][]int)) {
	newValue := -1
	for i := 0; i < 5; i++ {
		nextY := x
		nextX := len((*a)[0]) - (y + 1)
		if newValue != -1 {
			temp := (*a)[y][x]
			(*a)[y][x] = newValue
			newValue = temp
		} else {
			newValue = (*a)[y][x]
		}
		y = nextY
		x = nextX
	}
}
