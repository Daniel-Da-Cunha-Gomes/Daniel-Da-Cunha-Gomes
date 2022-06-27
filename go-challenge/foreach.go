package piscine

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}

/*Écrivez une fonction ForEachqui, pour une int tranche,
applique une fonction sur chaque élément de cette tranche.*/
