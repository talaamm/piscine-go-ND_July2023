package piscinego

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune(111)
		} else if x != 1 && y != 1 {
			z01.PrintRune(111)
			for count := 1; count < (x - 1); count++ {
				z01.PrintRune(45)
			}
			z01.PrintRune(111)
			z01.PrintRune('\n')

			for county := 1; county < (y - 1); county++ {
				z01.PrintRune(124)

				for countS := 1; countS < (x - 1); countS++ {
					z01.PrintRune(' ')
				}

				z01.PrintRune(124)
				z01.PrintRune('\n')

			}
			z01.PrintRune(111)

			for count := 1; count < (x - 1); count++ {
				z01.PrintRune(45)
			}
			z01.PrintRune(111)
			z01.PrintRune('\n')

		} else if x == 1 && y != 1 {
			z01.PrintRune(111)

			for county := 1; county < (y - 1); county++ {
				z01.PrintRune('\n')
				z01.PrintRune(124)

			}
			z01.PrintRune('\n')
			z01.PrintRune(111)

		} else if x != 1 && y == 1 {
			z01.PrintRune(111)
			for count := 1; count < (x - 1); count++ {
				z01.PrintRune(45)
			}
			z01.PrintRune(111)
			z01.PrintRune('\n')

		}
	}
}
