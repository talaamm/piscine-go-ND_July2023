package piscinego

import "github.com/01-edu/z01"

func Hello() {
	s := "Hello world"
	for _, r := range s {
		z01.PrintRune(r)
	}
}
