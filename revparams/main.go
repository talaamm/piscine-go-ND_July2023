package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args
	ArgsCounter := len(Args) - 1
	/*for index := range Args {
		ArgsCounter = index
	}*/
	for i := ArgsCounter; i >= 1; i-- {
		for _, PrintInRune := range Args[i] {
			z01.PrintRune(PrintInRune)
		}
		z01.PrintRune('\n')
	}
}
