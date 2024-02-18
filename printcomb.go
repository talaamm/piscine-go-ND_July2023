package piscinego

import "github.com/01-edu/z01"

func PrintComb() {
	for x := '0'; x <= '9'; x++ {
		for y := '0'; y <= '9'; y++ {
			for z := '0'; z <= '9'; z++ {
				if x < y && y < z {
					z01.PrintRune(x)
					z01.PrintRune(y)
					z01.PrintRune(z)
					if x <= '6' {
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}

				}
			}
		}
	}
	z01.PrintRune('\n')
}
