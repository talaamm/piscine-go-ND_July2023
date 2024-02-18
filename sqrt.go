package piscinego

func Sqrt(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	}

	for i := 1; i*i <= nb; i++ {
		result = i
	}
	if result*result == nb {
		return result
	} else {
		return 0
	}
}
