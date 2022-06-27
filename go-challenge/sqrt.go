package piscine

func Sqrt(nb int) int {
	x := 0
	for Puissance := 10; Puissance > 0; Puissance-- {
		x = Puissance * Puissance
		if nb == x {
			return Puissance
		}
	}
	return 0
}
