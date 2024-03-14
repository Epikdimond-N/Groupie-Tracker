package controller

import (
	"encoding/json"
	"f1/backend"
	initTemplate "f1/temps"
	"fmt"
	"net/http"
	"os"
)

var page = 1
var pilotes []backend.Pilote
var circuits []backend.Circuit

var filtresPilotes []backend.Pilote
var filtresCircuits []backend.Circuit

var filtre bool
var toSendPilotes []backend.Pilote   // liste des pilotes que je vais envoyer
var toSendCircuits []backend.Circuit // liste des circuits que je vais envoyer

func InitPilotes(w http.ResponseWriter, r *http.Request) { // Requete opti
	page = 1
	if pilotes != nil { // si on a deja fait la requete on se base sur pilote qui stock en local la requète globale
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	} else { // sinon on fait la requete
		pilotes = Pilotes()
		filtre = false
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	}
}

func InitCircuits(w http.ResponseWriter, r *http.Request) {
	page = 1
	if circuits != nil {
		http.Redirect(w, r, "/circuits", http.StatusSeeOther)
	} else { // sinon on fait la requete
		circuits = Circuits()
		filtre = false
		http.Redirect(w, r, "/circuits", http.StatusSeeOther)
	}
}

func NextPagePilote(w http.ResponseWriter, r *http.Request) { // on augmente la page
	page++
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func PreviousPagePilote(w http.ResponseWriter, r *http.Request) { // on décrément la page
	page--
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func NextPageCircuit(w http.ResponseWriter, r *http.Request) {
	page++
	http.Redirect(w, r, "/circuits", http.StatusSeeOther)
}

func PreviousPageCircuit(w http.ResponseWriter, r *http.Request) {
	page--
	http.Redirect(w, r, "/circuits", http.StatusSeeOther)
}

func DisplayAccueil(w http.ResponseWriter, r *http.Request) { //on affiche l'accueil
	filtre = false
	initTemplate.Temp.ExecuteTemplate(w, "accueil", nil)
}

func DisplayPilotes(w http.ResponseWriter, r *http.Request) { // affichage opti des pilotes
	toSendPilotes = nil
	var pilotesFound int // var pour opti la recherche
	if !filtre {
		for _, i := range pilotes { // on s'arrete des qu'on trouve 10 qui ont l'ID correspondant à la page sur laquelle on se trouve
			if i.PageId == page {
				toSendPilotes = append(toSendPilotes, i)
				pilotesFound++
				if pilotesFound == 10 {
					break
				}
			}
		}
	} else {
		for _, i := range filtresPilotes {
			if i.PageId == page {
				toSendPilotes = append(toSendPilotes, i)
			}
		}
	}
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", toSendPilotes)
}

func DisplayCircuits(w http.ResponseWriter, r *http.Request) { // on affiche le template circuits
	toSendCircuits = nil
	var circuitsFound int
	if !filtre {
		for _, i := range circuits {
			if i.PageId == page {
				toSendCircuits = append(toSendCircuits, i)
				circuitsFound++
				if circuitsFound == 10 {
					break
				}
			}
		}
	}
	initTemplate.Temp.ExecuteTemplate(w, "circuits", toSendCircuits)
}

func DisplayConstructeurs(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "constructeurs", nil)
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "login", nil)
}

func DisplayPiloteNotFound(w http.ResponseWriter, r *http.Request) {
	filtresPilotes = pilotes
	initTemplate.Temp.ExecuteTemplate(w, "pilotenotfound", nil)
}

func DisplayCircuitNotFound(w http.ResponseWriter, r *http.Request) {
	filtresCircuits = circuits
	initTemplate.Temp.ExecuteTemplate(w, "circuitnotfound", nil)
}

func DisplayPiloteSearch(w http.ResponseWriter, r *http.Request) {
	nompilote := r.FormValue("search")
	filtresPilotes = SearchPilote(nompilote)
	filtresPilotes = PaginationPilote(filtresPilotes)
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", filtresPilotes)
}

func DisplayCircuitSearch(w http.ResponseWriter, r *http.Request) {
	nomcircuit := r.FormValue("search")
	filtresCircuits = SearchCircuit(nomcircuit)
	filtresCircuits = PaginationCircuits(filtresCircuits)
	initTemplate.Temp.ExecuteTemplate(w, "circuits", filtresCircuits)
}

func InitFavoris(w http.ResponseWriter, r *http.Request) {
	toAdd := r.URL.Query().Get("id")
	fmt.Println("to add :", toAdd)
	for _, i := range pilotes {
		if i.DriverID == toAdd {
			fichierJSON, err := os.Open("fav.json")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier JSON :", err)
				return
			}
			defer fichierJSON.Close()
			encodeur := json.NewEncoder(fichierJSON)
			if err := encodeur.Encode(&i); err != nil {
				fmt.Println("Erreur lors de l'encodage en JSON :", err)
				return
			}
		}
	}
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}
