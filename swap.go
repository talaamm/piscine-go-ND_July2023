package piscinego

func Swap(a *int, b *int) {
	c := *a

	*a = *b
	*b = c
}
