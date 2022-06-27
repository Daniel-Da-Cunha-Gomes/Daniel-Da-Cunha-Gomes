package piscine

/*import (
	"fmt"
	"piscine"
)*/

/*func AlphaCount(s string) int {
	count := 0
	for i := range s{
		for i := 96 && i <= 123 ||  i >=65 && i <= 90 {
			count++
			z01.PrintRune(rune (count))
		}
	}
	return count
}*/

/*func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}*/

func AlphaCount(s string) int {
	var lettre string = ("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	count := 0
	for i := range s {
		for n := range lettre {
			if s[i] == lettre[n] {
				count++
			}
		}
	}
	return count
}
