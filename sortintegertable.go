package piscinego

func SortIntegerTable(table []int) {
	for i1 := range table[:len(table)-1] {
		for i2 := i1 + 1; i2 <= len(table)-1; i2++ {
			if table[i2] < table[i1] {
				table[i1], table[i2] = table[i2], table[i1]
			}
		}
	}
}
