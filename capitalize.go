package piscinego

/*/////////////////////////////////////////////////
						solution 1
/////////////////////////////////////////////////*/

func Capitalize(s string) string {
	newWord := true
	word := ""
	final := ""
	for _, r := range s {

		if !(r >= 'A' && r <= 'Z') && !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9') {
			newWord = true
			word += string(r)
			final += word
			word = ""
			continue

		}
		if !newWord {
			if r >= 'A' && r <= 'Z' {
				word += string(r + 32)
			} else {
				word += string(r)
			}
		}
		if newWord {
			if r >= 'a' && r <= 'z' {
				word += string(r - 32)
			} else {
				word += string(r)
			}
			newWord = false
		}

	}
	final += word
	return final
}

/*/////////////////////////////////////////////////
						solution 2
/////////////////////////////////////////////////*/

func CapitalizeCOPY(s string) string {
	str := []rune(s)
	isReset := true
	for ind, val := range str {
		if isReset {
			if ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9') {
				isReset = false
				continue
			} else if 'a' <= val && val <= 'z' {
				str[ind] -= 'a' - 'A'
				isReset = false
			}
		} else if !isReset && 'A' <= val && val <= 'Z' {
			str[ind] -= 'A' - 'a'
		} else if !(('a' <= val && val <= 'z') || ('A' <= val && val <= 'Z') || ('0' <= val && val <= '9')) {
			isReset = true
		}
	}
	return string(str)
}
