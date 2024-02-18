package piscinego

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 10000000000000000 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}

	result := nb * RecursiveFactorial(nb-1)
	return result
}
