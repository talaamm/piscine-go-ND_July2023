package piscinego

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for x := '0'; x <= '9'; x++ {
		for y := '0'; y <= '9'; y++ {
			for a := '0'; a <= '9'; a++ {
				for b := '0'; b <= '9'; b++ {
					//	if (x < a) || (x == a && y < b) {
					if (x*10 + y) < (a*10 + b) {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(a)
						z01.PrintRune(b)

						if x == '9' && y == '8' && a == '9' && b == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
}
