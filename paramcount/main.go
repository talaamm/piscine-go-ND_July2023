package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) < 10 {
		z01.PrintRune(rune(len(arg) + 48))
	} else {
		length := itoa(len(arg))
		for _, r := range length {
			z01.PrintRune(r)
		}
	}

}

func itoa(num int) string {
	b := num
	digits := []int{}
	res := ""

	for b != 0 {
		digits = append(digits, b%10)
		b /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		res += string(rune(digits[i] + 48))
	}
	return res
}
