package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//La structure de base où on aura toutes les infos des artistes
type Artist struct {
	Id             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	Locations      string
	ConcertDates   string
	Relations      string
	AlbumYear      []int
	FirstAlbumYear []int
}

//Copie de la structure Artist dont les valeurs vont être réadapté pour correspodre aux filtres
type FiltrerArtists struct {
	Id             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	Locations      string
	ConcertDates   string
	Relations      string
	AlbumYear      []int
	FirstAlbumYear []int
}

//Structure pour les relations
type Local struct {
	relationsss map[string]interface{}
}

//Définition des variables
var Artists []Artist                    //Pour récupérer les infos de l'API
var INFO Artist                         //Pour stoquer et utiliser des informations complémentaires de la structure
var Filter = make([]FiltrerArtists, 52) //Pour stoquer les informations correspondentes selon le filtre actif
var artists1 Local                      //Pour stoquer des valeurs dans la structure Local
var relations map[string]interface{}    //Pour récupérer les relations selon les artistes

//Applique le filte choisis par l'utilisateur
func ApplyFilter(DateMin int, DateMax int, TheFirstAlbum string, GroupSize []int) {
	var NewId []int
	var NewAlbumYear []int
	var NewMembers [][]string

	var BackupId []int
	var BackupAlbumYear []int
	var BackupMembers [][]string

	//Filtres les dates minimum et maximum
	for i := range Artists {
		if DateMin > 0 && DateMax > 0 {
			if Artists[i].CreationDate >= DateMin && Artists[i].CreationDate <= DateMax {
				NewId = append(NewId, Artists[i].Id)
				NewAlbumYear = append(NewAlbumYear, INFO.AlbumYear[i])
				NewMembers = append(NewMembers, Artists[i].Members)

			}
		} else {
			NewId = append(NewId, Artists[i].Id)
			NewAlbumYear = append(NewAlbumYear, INFO.AlbumYear[i])
			NewMembers = append(NewMembers, Artists[i].Members)
		}
	}
	BackupId = NewId
	BackupAlbumYear = NewAlbumYear
	BackupMembers = NewMembers
	NewId = nil
	NewAlbumYear = nil
	NewMembers = nil

	//Filtre selon l'année du premier Album
	IntAlbum, _ := strconv.Atoi(TheFirstAlbum)
	for i, r := range BackupAlbumYear {
		if TheFirstAlbum != "Toutes" {
			if r == IntAlbum {
				NewId = append(NewId, BackupId[i])
				NewMembers = append(NewMembers, BackupMembers[i])
			}
		} else {
			NewId = append(NewId, BackupId[i])
			NewMembers = append(NewMembers, BackupMembers[i])
		}
	}
	BackupId = NewId
	BackupMembers = NewMembers
	NewId = nil
	NewAlbumYear = nil
	NewMembers = nil

	//Filtre selon la taille du groupe
	for i := range BackupMembers {
		for j := range GroupSize {
			if GroupSize[j] < 6 {
				if len(BackupMembers[i]) == GroupSize[j] {
					NewId = append(NewId, BackupId[i])
				}
			} else {
				if len(BackupMembers[i]) >= GroupSize[j] {
					NewId = append(NewId, BackupId[i])
				}
			}
		}
	}
	BackupId = NewId
	NewId = nil

	//Recherche de corespondences des filtres appliqués
	position := 0
	for _, r := range BackupId {
		Filter[position] = FiltrerArtists(Artists[r-1])
		position++
	}

}

//Tri des dates dans l'ordre pour les afficher dans le filtre
func Year(ArtistYear []int) []int {
	var AllYears []int
	for i := 1900; i < 2100; i++ {
		for _, r := range ArtistYear {
			if len(AllYears) == 0 {
				if i == r {
					AllYears = append(AllYears, r)
				}
			} else {
				if i == r && AllYears[len(AllYears)-1] != r {
					AllYears = append(AllYears, r)
				}
			}

		}
	}
	return AllYears
}

func main() {
	//Pour envoyer plusieurs structure à la template
	type M map[string]interface{}

	//Récupération des données de l'API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	artistsJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(artistsJSON, &Artists)

	//toutes les années dans l'ordre avec répétitions des années
	var year string
	var Album = make([]int, len(Artists))
	for j, artist := range Artists {
		for i, r := range artist.FirstAlbum {
			if i > 5 {
				year += string(r)
			}
		}
		AlbumInt, err := strconv.Atoi(year)
		if err != nil {
			log.Fatal(err)
		}
		Album[j] = AlbumInt
		year = ""
	}

	INFO.AlbumYear = Album
	INFO.FirstAlbumYear = Year(INFO.AlbumYear)

	//definition de la template
	tmpl := template.Must(template.ParseGlob("html/*"))

	//Page d'acceuil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "vitrine.html", Artists)
	})

	//Page avec la liste des artistes
	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {

		NomberOfMember := make([]int, 6)

		datesMin, _ := strconv.Atoi(r.FormValue("DatesMin"))
		datesMax, _ := strconv.Atoi(r.FormValue("DatesMax"))
		theFirstAlbum := r.FormValue("FirstAlbum")
		for i := 0; i < 6; i++ {
			MemeberSize, _ := strconv.Atoi(r.FormValue(strconv.Itoa(i + 1)))
			NomberOfMember[i] = MemeberSize
		}

		ApplyFilter(datesMin, datesMax, theFirstAlbum, NomberOfMember)

		switch r.Method {
		case "POST":
			http.Redirect(w, r, "http://localhost:5550/artists/filter", http.StatusSeeOther)
		}

		tmpl.ExecuteTemplate(w, "index.html", M{
			"Artists": Artists,
			"INFO":    INFO,
		})

	})

	//Page des filtre lorsqu'ils sont appliqués
	http.HandleFunc("/artists/filter", func(w http.ResponseWriter, r *http.Request) {
		NomberOfMember := make([]int, 6)

		datesMin, _ := strconv.Atoi(r.FormValue("DatesMin"))
		datesMax, _ := strconv.Atoi(r.FormValue("DatesMax"))
		theFirstAlbum := r.FormValue("FirstAlbum")
		for i := 0; i < 6; i++ {
			MemeberSize, _ := strconv.Atoi(r.FormValue(strconv.Itoa(i + 1)))
			NomberOfMember[i] = MemeberSize
		}

		ApplyFilter(datesMin, datesMax, theFirstAlbum, NomberOfMember)

		switch r.Method {
		case "POST":
			http.Redirect(w, r, "http://localhost:5550/artists/filter", http.StatusSeeOther)
		}

		tmpl.ExecuteTemplate(w, "filter.html", M{
			"Filter": Filter,
			"INFO":   INFO,
		})
	})

	//Page du service client
	http.HandleFunc("/service-client", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "service-client.html", Artists)
	})

	//Page d'un artiste défini
	http.HandleFunc("/artists/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.RequestURI()
		var num string
		for _, rn := range url {
			if rn >= '0' && rn <= '9' {
				num += string(rn)
			}
		}
		number, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		a_, _ := http.Get(Artists[number-1].Relations)
		Data, _ := ioutil.ReadAll(a_.Body)
		json.Unmarshal(Data, &relations)
		artists1.relationsss = relations

		tmpl.ExecuteTemplate(w, "artists.html", M{
			"Artists":   Artists[number-1],
			"Relations": artists1.relationsss["datesLocations"],
		})
	})

	//Prise en charge du CSS par la template
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe("localhost:5550", nil)
}
