package piscinego

func ToLower(s string) string {
	result := []rune(s)
	for r := range result {
		if result[r] >= 65 && result[r] <= 90 {
			result[r] = result[r] + 32
		}
	}
	return string(result)
}
