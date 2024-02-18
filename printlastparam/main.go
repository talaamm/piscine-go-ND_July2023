package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[len(os.Args)-1]

	if len(os.Args) == 1 {
		return
	} else {
		for _, i := range arr {
			z01.PrintRune(i)
		}
	}

}
