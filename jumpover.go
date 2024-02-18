package piscinego

func JumpOver(str string) string {
	result := ""
	if str == "" || len(str) <= 2 {
		return string('\n')
	}

	for indx, r := range str {
		if (indx+1)%3 == 0 {
			result += string(r)
		}
	}

	result += string('\n')
	return result
}
