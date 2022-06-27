package piscine

import "github.com/01-edu/z01"

func main() {
	PrintNbrInOrder(32132)
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var tableau []int
		valeur := 0

		tabcompteur := 0
		var valeurmini int
		for n != 0 {
			valeur = n % 10
			n /= 10
			tableau = append(tableau, valeur)
		}

		for compteur := range tableau {
			tabcompteur = compteur + 1
		}
		for i := 0; i < tabcompteur-1; i++ {
			for j := 0; j < tabcompteur-i-1; j++ {
				if tableau[j] > tableau[j+1] {
					valeurmini = tableau[j]
					tableau[j] = tableau[j+1]
					tableau[j+1] = valeurmini
				}
			}
		}
		for i := 0; i < tabcompteur; i++ {
			z01.PrintRune(rune(tableau[i] + 48))
		}
	}
}
