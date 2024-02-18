package piscinego

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, r := range a {
		for _, i := range r {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
