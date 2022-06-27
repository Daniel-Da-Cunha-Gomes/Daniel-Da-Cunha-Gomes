package piscine

func ToLower(s string) string {
	pomme := []rune(s)
	for x := range pomme {
		if pomme[x] >= 'A' && pomme[x] <= 'Z' {
			pomme[x] = pomme[x] + 32
		}
	}
	return string(pomme)
}
