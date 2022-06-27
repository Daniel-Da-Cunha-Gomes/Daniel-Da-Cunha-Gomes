package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	*b = *a % *b
	*a = c
}
