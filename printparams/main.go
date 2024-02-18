package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	for _, strings := range Args {
		for _, PrintInRune := range strings {
			z01.PrintRune(PrintInRune)
		}
		z01.PrintRune('\n')
	}
}
