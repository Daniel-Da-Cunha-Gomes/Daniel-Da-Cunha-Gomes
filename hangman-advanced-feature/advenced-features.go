package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Random(a int) int { //chiffre random
	return rand.Intn(a)
}

func Verif(a int, sl []int) bool { //si c'est une lettre random choisie
	for i := 0; i < len(sl); i++ {
		if a == sl[i] {
			return true
		}
	}
	return false
}

func Initial(s string) string {
	var mot []string      // EX : salut
	var motcaché []string // EX : _ _ _ _ _
	for _, x := range s {
		mot = append(mot, string(x))
		motcaché = append(motcaché, "_")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (len(mot))/2-1; i++ {
		pos := rand.Intn(len(mot))
		motcaché[pos] = mot[pos]
	}
	res := ""
	for i := 0; i < len(motcaché); i++ {
		res = res + motcaché[i] + " "
	}
	return res
}

func Position(s string) [10]string { //répertorier les différentes étapes de hangman
	var slice [10]string
	count := 0
	str := ""
	data, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("failed opening file")
	}
	i := 0
	for _, r := range data {
		str += string(r)
		if r == '=' {
			count++
		}
		if r == '\n' && count == 9 {
			slice[i] = str
			str = ""
			count = 0
			i++
		}
	}
	return slice
}

func main() { //"rond point du progamme" prépare toutes les valeurs la structure
	var hangman HangManData //création de la variable pour la stucture
	var p [10]string
	word := ""
	ToFind := ""
	a := 0
	var slice []string
	data, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("failed opening file")
	}
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		slice = append(slice, (scanner.Text()))
	}
	max := len(slice)
	min := 0
	a = Random(max - min)
	word = slice[a]
	ToFind = Initial(word)
	p = Position("hangman.txt")
	hangman.hang(word, ToFind, 10, p) //envoie les valeurs à la structure
	hangman.Display()                 //lance la fonction Display
}

func (h *HangManData) hang(Word string, ToFind string, Attemps int, HangmanPosition [10]string) { //pointe les valeurs reçus
	h.Word = Word
	h.ToFind = ToFind
	h.Attempts = Attemps
	h.HangmanPositions = HangmanPosition
}

func IsHere(s string, toFind string) bool { //vérifie si la lettre est présente dans le mot finale
	z := []rune(toFind)
	for _, r := range s {
		if r == z[0] {
			return true
		}
		if r-32 == z[0] {
			return true
		}
	}
	return false
}

func IsUsed(s string, tf string, x string) bool { //on vérifie la longeur de la lettre reçu et vérifie si on l'a déjà utilisé dans le mot à completer
	count := 0
	count1 := 0
	rx := []rune(x)
	if len(x) != 1 {
		return false
	}
	for _, r := range s {
		if r == rx[0] {
			count++
		}
	}
	for _, r1 := range tf {
		if r1 == rx[0] {
			count1++
		}
	}
	if count == count1 {
		return false
	}
	return true
}

func Solution(s string) string { //on crée la réponse (hello --> h e l l o)
	str := ""
	for i, r := range s {
		if i < len(s) {
			str += string(r) + " "
		} else {
			str += string(r)
		}
	}
	return str
}

func Change(temp string, s string, add string) string { // change le str en le nv
	var tsl []rune
	radd := []rune(add)
	str := ""
	for _, r := range temp {
		tsl = append(tsl, r)
	}
	for i, r1 := range s {
		if r1 == '_' && tsl[i] == radd[0] {
			r1 = tsl[i]
		}
		str += string(r1)
	}
	return str
}

func ToUpper(s string) string {
	pomme := []rune(s)
	for x := range pomme {
		if pomme[x] >= 'a' && pomme[x] <= 'z' {
			pomme[x] = pomme[x] - 32
		}
	}
	return string(pomme)
}

func VerifWord(s string, toFind string) bool { //verifie si le mot proposé est correct
	if s == toFind {
		return true
	}
	return false
}

func Array(x string, s []string) bool { //verifie si la lettre n'a pas déjà été utilisée
	if len(s) == 0 {
		return true
	} else {
		for _, a := range s {
			if x == string(a) {
				return false
			}
		}
	}
	return true
}

func (h HangManData) Display() { //donne forme au jeu
	var sdf []string
	utf := h.ToFind
	uw := h.Word
	fmt.Println("Good Luck, you have", h.Attempts, "attempts.")
	fmt.Println(ToUpper(utf))
	sol := Solution(h.Word)
	for h.Attempts != 0 || h.ToFind == h.Word {
		utf = h.ToFind
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Choose :")
		scanner.Scan()
		x := scanner.Text()
		if len(x) != 0 {
			if len(x) == 1 {
				P := Array(x, sdf)
				if P == true {
					sdf = append(sdf, x)
				}
				I := IsHere(h.Word, x)
				J := IsUsed(h.Word, h.ToFind, x)
				if I == true && J == true && P == true {
					ch := Change(sol, h.ToFind, x)
					h.ToFind = ch
					utf = h.ToFind
					if sol == h.ToFind {
						fmt.Println("Congratulation !")
						fmt.Println("You find the word", ToUpper(uw), "!")
						break
					} else {
						fmt.Println(ToUpper(utf))
					}
				} else {
					if P == false {
						atp := h.Attempts
						fmt.Println("You already use this letter,", h.Attempts, "attempts remaining")
						fmt.Println(ToUpper(utf))
						fmt.Println(h.HangmanPositions[9-atp])
					} else {
						if h.Attempts == 1 {
							fmt.Println("Game Over")
							fmt.Println("The correct word was", ToUpper(uw))
							fmt.Println(h.HangmanPositions[9])
							break
						} else {
							h.Attempts -= 1
							atp := h.Attempts
							fmt.Println("Not present in the word,", h.Attempts, "attempts remaining")
							fmt.Println(ToUpper(utf))
							fmt.Println(h.HangmanPositions[9-atp])
						}
					}
				}
			}
			if len(x) > 1 {
				nj := VerifWord(h.Word, x)
				if nj == true {
					fmt.Println("Congratulation !")
					fmt.Println("You find the word", ToUpper(uw), "!")
					break
				} else {
					if h.Attempts <= 1 {
						fmt.Println("Game Over")
						fmt.Println("The correct word was", ToUpper(uw))
						fmt.Println(h.HangmanPositions[9])
						break
					} else {
						h.Attempts -= 2
						atp := h.Attempts
						fmt.Println("Not present in the word,", h.Attempts, "attempts remaining")
						fmt.Println(ToUpper(utf))
						fmt.Println(h.HangmanPositions[9-atp])
					}
				}
			}
		} else {
			fmt.Println("No letter detected")
			fmt.Println(ToUpper(utf))
		}
	}
}
