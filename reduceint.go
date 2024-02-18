package piscinego

import "github.com/01-edu/z01"

func ReduceInt(a []int, f func(int, int) int) {
	res := f(a[0], a[1])

	for i := 2; i < len(a); i++ {
		res = f(res, a[i])
	}

	printNum := itoa(res)
	for _, r := range printNum {
		z01.PrintRune(r)
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
