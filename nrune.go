package piscinego

func NRune(s string, n int) rune {
	if n > len(s) || n < 1 {
		return 0
	}

	return rune(s[n-1])
}
