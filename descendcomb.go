package piscinego

import "github.com/01-edu/z01"

func DescendComb() {
	for x := '9'; x >= '0'; x-- {
		for y := '9'; y >= '0'; y-- {
			for a := '9'; a >= '0'; a-- {
				for b := '9'; b >= '0'; b-- {
					if (x > a) || (x == a && y > b) {
						z01.PrintRune(x)
						z01.PrintRune(y)
						z01.PrintRune(' ')
						z01.PrintRune(a)
						z01.PrintRune(b)

						if !(x == '0' && y == '1' && a == '0' && b == '0') {

							z01.PrintRune(',')
							z01.PrintRune(' ')
						}

					}
				}
			}
		}
	}
	// z01.PrintRune('\n')
}
