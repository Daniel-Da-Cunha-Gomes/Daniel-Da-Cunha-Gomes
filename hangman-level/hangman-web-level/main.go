package main

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
}

func (h *HangManData) hang(Word string, ToFind string, Attemps int, HangmanPosition [10]string) { //pointe les valeurs reçus
	h.Word = Word
	h.ToFind = ToFind
	h.Attempts = Attemps
	h.HangmanPositions = HangmanPosition
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//chiffre random
func Random(a int) int {
	return rand.Intn(a)
}

//crée le mot à compléter
func Initial(s string) string {
	var word []string   // EX : salut
	var ToFind []string // EX : _ _ _ _ _
	for _, x := range s {
		word = append(word, string(x))
		ToFind = append(ToFind, "_")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (len(word))/2-1; i++ {
		pos := rand.Intn(len(word))
		ToFind[pos] = word[pos]
	}
	res := ""
	for i := 0; i < len(ToFind); i++ {
		res = res + ToFind[i] + " "
	}
	return res
}

//répertorier les différentes étapes de hangman
func Position(s string) [10]string {
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

//vérifie si la lettre est présente dans le mot final
func IsHere(s string, lettre string) bool {
	z := []rune(lettre)
	for _, r := range s {
		if r == z[0] {
			return true
		}
	}
	return false
}

//on vérifie la longeur de la lettre reçu et si elle est déjà dans le mot à compléter
func IsUsed(s string, tf string, x string) bool {
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

//on crée la réponse (hello --> h e l l o)
func Solution(s string) string {
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

//met à jour le mot à compléter
func Change(temp string, s string, add string) string {
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

//prépare toutes les valeurs la structure
func main() {
	var hangman HangManData //création de la variable pour la stucture
	//définition des variables
	var p [10]string
	word := ""
	ToFind := ""
	a := 0
	var slice []string
	var slice2 []string
	var slice3 []string

	data, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("failed opening file")
	}
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		slice = append(slice, (scanner.Text()))
	}
	data2, err := os.Open("words2.txt")
	if err != nil {
		fmt.Println("failed opening file")
	}
	scanner2 := bufio.NewScanner(data2)
	scanner2.Split(bufio.ScanLines)
	for scanner2.Scan() {
		slice2 = append(slice2, (scanner2.Text()))
	}
	data3, err := os.Open("words3.txt")
	if err != nil {
		fmt.Println("failed opening file")
	}
	scanner3 := bufio.NewScanner(data3)
	scanner3.Split(bufio.ScanLines)
	for scanner3.Scan() {
		slice3 = append(slice3, (scanner3.Text()))
	}

	p = Position("hangman.txt")
	FinalToFind := ""
	FinalWord := ""
	menu := true

	tmpl := template.Must(template.ParseGlob("html/*"))

	//création de la template
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if menu == true {
			switch r.Method {
			case "GET":
				t, _ := template.ParseFiles("html/difficulte.html")
				t.Execute(w, "difficulte.html")
				return
			case "POST":
				button := r.FormValue("button")
				switch button {
				case "EASY":
					max := len(slice)
					min := 0
					a = Random(max - min)
					word = slice[a]
					UpperWord := Solution(word)
					ToFind = Initial(word)
					hangman.hang(UpperWord, ToFind, 10, p) //envoie les valeurs à la structure
					FinalToFind = hangman.ToFind
					FinalWord = hangman.Word

				case "NORMAL":
					max := len(slice2)
					min := 0
					a = Random(max - min)
					word = slice2[a]
					UpperWord := Solution(word)
					ToFind = Initial(word)
					hangman.hang(UpperWord, ToFind, 10, p)
					FinalToFind = hangman.ToFind
					FinalWord = hangman.Word

				case "HARD":
					max := len(slice3)
					min := 0
					a = Random(max - min)
					word = slice3[a]
					UpperWord := Solution(word)
					ToFind = Initial(word)
					hangman.hang(UpperWord, ToFind, 10, p)
					FinalToFind = hangman.ToFind
					FinalWord = hangman.Word
				}
				http.Redirect(w, r, "http://localhost:8080/hangman/play", http.StatusSeeOther)
			}
			menu = false
		} else {
			scanner_input := r.FormValue("x")

			if len(scanner_input) == 1 {
				I := IsHere(hangman.Word, scanner_input)
				J := IsUsed(hangman.Word, hangman.ToFind, scanner_input)
				//condition qui vérifie si la lettre entrée par le joueur est déjà utilisée
				if I == true && J == true {
					ch := Change(hangman.Word, hangman.ToFind, scanner_input)
					hangman.ToFind = ch
					FinalToFind = hangman.ToFind
					if hangman.Word == hangman.ToFind {
						tmpl.ExecuteTemplate(w, "end.html", hangman)
					} else {
						hangman.hang(FinalWord, FinalToFind, hangman.Attempts, p)
						tmpl.ExecuteTemplate(w, "index.html", hangman)
					}
				} else {
					//condition qui vérifie si le joueur a perdu
					if hangman.Attempts == 1 {
						hangman.hang(FinalWord, FinalToFind, hangman.Attempts, p)
						tmpl.ExecuteTemplate(w, "end.html", hangman)
						//actions si on appuie sur rejouer
					} else {
						hangman.Attempts -= 1
						hangman.hang(FinalWord, FinalToFind, hangman.Attempts, p)
						tmpl.ExecuteTemplate(w, "index.html", hangman)
					}

				}
			} else {
				hangman.hang(FinalWord, FinalToFind, hangman.Attempts, p)
				tmpl.ExecuteTemplate(w, "index.html", hangman)
			}
		}
	})

	//envoie les données au localhost choisis
	http.ListenAndServe(":8080", nil)
}
