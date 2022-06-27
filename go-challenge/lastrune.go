package piscine

func LastRune(s string) rune {
	l := 0
	for range s {
		l++
	}
	lettre := rune(s[l-1])
	return lettre
}
