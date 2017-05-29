package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"time"
)

// Un objet représentant une nouvelle
type News struct {
	XMLName      xml.Name  `json:"-" xml:"nouvelle"`
	Titre        string    `json:"titre" xml:"titre"`
	Contenu      string    `json:"contenu" xml:"contenu"`
	DateCreation time.Time `json:"date" xml:"date,attr"`
}

var (
	// Notre liste de nouvelles
	listNews = []News{
		News{Titre: "Titre 1", Contenu: "Contenu 1", DateCreation: time.Now()},
		News{Titre: "Titre 2", Contenu: "Contenu 2", DateCreation: time.Now().Add(50 * time.Minute)},
	}
)

func main() {
	http.HandleFunc("/nouvelles.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(listNews); err != nil {
			http.Error(w, "Impossible de créer le JSON", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/nouvelles.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		if err := xml.NewEncoder(w).Encode(listNews); err != nil {
			http.Error(w, "Impossible de créer le XML", http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":8080", nil)
}
