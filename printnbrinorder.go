package piscinego

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	nb := n
	digits := []int{}
	for nb != 0 {
		digits = append(digits, nb%10)
		nb = nb / 10
	}
	for i := 0; i < len(digits)-1; i++ {
		for k := i + 1; k <= len(digits)-1; k++ {
			if digits[i] > digits[k] {
				digits[i], digits[k] = digits[k], digits[i]
			}
		}
	}
	for _, r := range digits {
		z01.PrintRune(rune(r + '0'))
	}
}
