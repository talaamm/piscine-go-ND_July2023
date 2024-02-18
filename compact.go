package piscinego

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0
	writeIndex := 0

	for _, value := range slice {
		if value != "" {
			slice[writeIndex] = value
			writeIndex++
			count++
		}
	}

	*ptr = slice[:count]
	return count
}
