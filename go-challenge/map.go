package piscine

func Map(f func(int) bool, a []int) []bool {
	l := len(a)
	s := make([]bool, l)
	for i := 0; i < l; i++ {
		s[i] = f(a[i])
	}
	return s
}

/*Écrivez une fonction Mapqui, pour une inttranche,
applique une fonction de ce type func(int) bool
sur chaque élément de cette tranche et renvoie une tranche
de toutes les valeurs de retour.*/
