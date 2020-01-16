package piscine

func Swap(a *int, b *int) {
	one := *a
	two := *b
	*b = one
	*a = two
}
