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

func Split(s []byte) []string { //mettre les mots possibles dans une slice
	var slice []string
	word := ""
	for i, r := range s {
		if r != '\n' && i != len(s)-1 {
			word += string(r)
		} else {
			if r == '\n' && i != len(s)-1 {
				slice = append(slice, word)
				word = ""
			}
			if i == len(s)-1 {
				word += string(r)
				slice = append(slice, word)
			}
		}
	}
	return slice
}

func Verif(a int, sl []int) bool { //si c'est une lettre random choisie
	for i := 0; i < len(sl); i++ {
		if a == sl[i] {
			return true
		}
	}
	return false
}

func Words(s string) string { //créer le mot affiché (pour hello ça va afficer _ _ l _ o)
	str := ""
	var slice []int
	a := len(s)/2 - 1
	ran := 0
	for j := 0; j < a; j++ {
		ran = Random(len(s) - 1)
		slice = append(slice, ran)
	}
	for i, r := range s {
		x := Verif(i, slice)
		if x == true {
			str += string(r)
		} else {
			str += "_"
		}
	}
	return str
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
	data, err := os.ReadFile("words.txt")
	if err != nil {
		fmt.Println("failed opening file")
	}
	slice = Split(data)
	max := len(slice)
	min := 0
	a = Random(max - min)
	word = slice[a]
	ToFind = Words(word)
	p = Position("hangman.txt")
	hangman.hang(word, ToFind, 10, p) //envoie les valeurs à la structure
	hangman.Display()                 //lance la fonction Jose
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

/*
func Solution(s string) string { //on crée la réponse (hello --> h e l l o)
	str := ""
	for _, r := range s {
		str += string(r) + " "
	}
	return str
}
*/

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

func Letter(s string) []string { //répertorier les lettres en ASCII
	var line []string
	file, err := os.Open(s)
	if err != nil {
		fmt.Println("failed opening file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	return line
}

func AsciiConv(s string, ascii []string) []string {
	var AllLetters []string
	var Used []int
	var UsedAscii []string
	for i := 32; i < 127; i++ {
		AllLetters = append(AllLetters, string(rune(i)))
	}
	for _, r := range s {
		for i1, r1 := range AllLetters {
			if string(r) == r1 {
				Used = append(Used, i1)
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j := range Used {
			UsedAscii = append(UsedAscii, ascii[Used[j]*9+i])
		}
	}
	return UsedAscii
}

func (h HangManData) Display() { //donne forme au jeu
	var z []string
	count := 0
	utf := h.ToFind
	uw := h.Word
	ascii := Letter("standard.txt")
	fmt.Println("Good Luck , you have", h.Attempts, "attempts.")
	fmt.Println(utf)
	y := AsciiConv(utf, ascii)
	for i := range y {
		fmt.Printf(y[i])
		if i%len(h.ToFind) == len(h.ToFind)-1 && i >= len(h.ToFind) {
			fmt.Printf(string('\n'))
		}
	}
	fmt.Printf(string('\n'))
	/*
		sol := Solution(h.Word)
	*/
	for h.Attempts != -1 || h.ToFind == h.Word {
		utf = h.ToFind
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Choose :")
		scanner.Scan()
		x := scanner.Text()
		if len(x) == 1 {
			I := IsHere(h.Word, x)
			J := IsUsed(h.Word, h.ToFind, x)
			if I == true && J == true {
				ch := Change(h.Word, h.ToFind, x)
				h.ToFind = ch
				utf = h.ToFind
				if h.Word == h.ToFind {
					fmt.Println("Congratulation !")
					fmt.Println("You find the word :")
					z = AsciiConv(uw, ascii)
					for i := range z {
						fmt.Printf(z[i])
						if i%len(h.ToFind) == len(h.ToFind)-1 && i >= len(h.ToFind) {
							fmt.Printf(string('\n'))
						}
					}
					break
				} else {
					fmt.Println(utf)
					z = AsciiConv(utf, ascii)
					for i := range z {
						fmt.Printf(z[i])
						if i%len(h.ToFind) == len(h.ToFind)-1 && i >= len(h.ToFind) {
							fmt.Printf(string('\n'))
						}
					}
				}
			} else {
				if h.Attempts == 1 {
					fmt.Println("Game Over")
					fmt.Println("The correct word was :")
					z = AsciiConv(uw, ascii)
					for i := range z {
						fmt.Printf(z[i])
						if i%len(h.ToFind) == len(h.ToFind)-1 && i >= len(h.ToFind) {
							fmt.Printf(string('\n'))
						}
					}
					fmt.Println(h.HangmanPositions[9])
					break
				} else {
					h.Attempts -= 1
					atp := h.Attempts
					fmt.Println("Not present in the word,", h.Attempts, "attempts remaining")
					fmt.Println(utf)
					z = AsciiConv(utf, ascii)
					for i := range z {
						fmt.Printf(z[i])
						if i%len(h.ToFind) == len(h.ToFind)-1 && i >= len(h.ToFind) {
							fmt.Printf(string('\n'))
						}
					}
					fmt.Println(h.HangmanPositions[9-atp])
					count++
				}
			}
		} else {
			fmt.Println("No letter detected")
			fmt.Println(utf)
			z = AsciiConv(utf, ascii)
			for i := range z {
				fmt.Printf(z[i])
				if i%len(h.ToFind) == len(h.ToFind)-1 && i >= len(h.ToFind) {
					fmt.Printf(string('\n'))
				}
			}
		}
	}
}
