package piscinego

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= nb; i++ {
		if nb > 10000000000000000 {
			return 0
		}
		result = result * i
	}
	return result
}
