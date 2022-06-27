package piscine

func Any(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			return true
		}
	}
	return false
}

/*Écrivez une fonction Anyqui renvoie true,
pour une stringtranche :

si, lorsque cette stringtranche est passée à travers une
f fonction, au moins un élément renvoie true.*/
