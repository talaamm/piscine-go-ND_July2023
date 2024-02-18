package piscinego

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if r >= 65 && r <= 90 {
			count += 1
		} else if r >= 97 && r <= 122 {
			count += 1
		}
	}
	return count
}
