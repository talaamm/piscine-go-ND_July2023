package piscinego

func ToUpper(s string) string {
	result := []rune(s)
	for r := range result {
		if result[r] >= 97 && result[r] <= 122 {
			result[r] = result[r] - 32
		}
	}
	return string(result)
}
