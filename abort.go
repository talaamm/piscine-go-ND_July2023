package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func Abort(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}

	swapped := true
	n := len(arr)

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				// Swap the elements
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
		n-- // Reduce the length to avoid re-checking the last element (already in place)
	}
	return arr[2]
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func Abort2(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}
	for i := 0; i < len(arr)-1; i++ {
		for k := 1 + i; k <= len(arr)-1; k++ {
			if arr[i] > arr[k] {
				arr[i], arr[k] = arr[k], arr[i]
			}
		}
	}
	return arr[2]
}
