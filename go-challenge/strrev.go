package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/*func StrRev(s string) string {
    L1 := []rune(s)
    L2 := []rune(s)
    for i := range s {
        L1[i] = L2[len(L1)-1-i]
    }
    return string(L1)
}*/
// -1 important pour commencer le string de 0
