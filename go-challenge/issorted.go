package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	res := 1
	res2 := 1
	res3 := 1
	for a, b := range tab {
		if a != len(tab)-1 {
			if f(b, tab[a+1]) < 0 {
				res++
			}
			if f(b, tab[a+1]) > 0 {
				res2++
			}
			if f(b, tab[a+1]) == 0 {
				res3++
			}
		}
	}
	return res == len(tab) || res2 == len(tab) || res3 == len(tab)
}
