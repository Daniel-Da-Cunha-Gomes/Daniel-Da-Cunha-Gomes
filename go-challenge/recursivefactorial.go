package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 99 {
		return 0
	}
	if nb > 0 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 1
}
