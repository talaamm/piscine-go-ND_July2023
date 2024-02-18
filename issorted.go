package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func IsSorted1(f func(a, b int) int, a []int) bool {
	length := len(a)
	asciending := true
	descending := true

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) >= 0) {
			descending = false
		}
	}

	for i := 1; i < length; i++ {
		if !(f(a[i-1], a[i]) <= 0) {
			asciending = false
		}
	}

	return asciending || descending
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func IsSorted(f func(a, b int) int, a []int) bool {
	for r := 0; r < len(a)-1; r++ {
		if f(a[r], a[r+1]) >= 0 {
			return false
		}
	}
	return true
}
