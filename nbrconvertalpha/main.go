package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nums := os.Args[1:]
	ascii := []int{}
	digs := 0
	res := ""
	up := false
	if len(os.Args) <= 1 || os.Args[1] == "" {
		return
	}
	for i, num := range nums {
		if nums[i] == "--upper" {
			up = true
			continue
		} else {
			for _, dig := range num {
				digs = int(dig-48) + digs*10
			}
			ascii = append(ascii, digs)
			digs = 0
		}
	}

	for _, k := range ascii {
		if k >= 1 && k <= 26 {
			if up {
				res += string(rune(96 + k - 32))
			} else {
				res += string(rune(96 + k))
			}
		} else {
			res += " "
		}
	}

	for _, s := range res {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
