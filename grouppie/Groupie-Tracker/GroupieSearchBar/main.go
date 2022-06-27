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

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate rune
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Local struct {
	relationsss map[string]interface{}
}

func main() {

	var artists []Artist
	var artists1 Local
	var relations map[string]interface{}
	type M map[string]interface{}

	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	artistsJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(artistsJSON, &artists)
	tmpl := template.Must(template.ParseGlob("html/*"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input := r.FormValue("x")
		for taille, contenu := range artists {
			for _, membres := range contenu.Members {
				if input == contenu.Name || input == membres || input == contenu.FirstAlbum {
					rr := strconv.Itoa(artists[taille].Id)
					http.Redirect(w, r, "http://localhost:5550/artists/"+rr, http.StatusSeeOther)
				}
			}
		}
		tmpl.ExecuteTemplate(w, "index.html", M{
			"Artists":   artists,
			"Relations": artists1.relationsss["datesLocations"],
		})
	})

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
		a_, _ := http.Get(artists[number-1].Relations)
		Data, _ := ioutil.ReadAll(a_.Body)
		json.Unmarshal(Data, &relations)
		artists1.relationsss = relations
		tmpl.ExecuteTemplate(w, "artists.html", M{
			"Artists":   artists[number-1],
			"Relations": artists1.relationsss["datesLocations"],
		})
	})

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe("localhost:5550", nil)
}
