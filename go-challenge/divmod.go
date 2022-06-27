package piscine

func DivMod(a int, b int, div *int, mod *int) {
	divi := a / b
	modo := a % b
	*div = divi
	*mod = modo
}
