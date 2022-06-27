package piscine

func ToUpper(s string) string {
	pomme := []rune(s)
	for x := range pomme {
		if pomme[x] >= 'a' && pomme[x] <= 'z' {
			pomme[x] = pomme[x] - 32
		}
	}
	return string(pomme)
}
