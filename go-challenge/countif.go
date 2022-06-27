package piscine

func CountIf(f func(string) bool, tab []string) int {
	pomme := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			pomme++
		}
	}
	return pomme
}

/*Écrivez une fonction CountIf qui renvoie le nombre d'éléments
d'une stringtranche pour laquelle la ffonction renvoie true.*/
