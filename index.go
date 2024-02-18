package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func Index1(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		found := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func Index(s string, toFind string) int {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
