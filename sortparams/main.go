package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	strs := os.Args[1:]
	arr := []string{}

	arr = append(arr, strs...)

	for i1 := 0; i1 < len(arr)-1; i1++ {
		for i2 := i1 + 1; i2 <= len(arr)-1; i2++ {
			if arr[i1][0] > arr[i2][0] {
				arr[i1], arr[i2] = arr[i2], arr[i1]
			}
		}
	}

	for _, w := range arr {
		for _, r := range w {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
