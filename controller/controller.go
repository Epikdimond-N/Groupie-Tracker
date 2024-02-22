package controller

import (
	"f1/backend"
	initTemplate "f1/temps"
	"net/http"
)

var page = 1
var pilotes []backend.Pilote
var filtresPilotes []backend.Pilote
var filtre bool
var toSend []backend.Pilote // liste des pilotes que je vais envoyer

func InitPilotes(w http.ResponseWriter, r *http.Request) { // Requete opti
	if pilotes != nil { // si on a deja fait la requete on se base sur pilote qui stock en local la requète globale
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	} else { // sinon on fait la requete
		pilotes = Pilotes()
		filtre = false
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	}
}

func NextPage(w http.ResponseWriter, r *http.Request) { // on augmente la page
	page++
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func PreviousPage(w http.ResponseWriter, r *http.Request) { // on décrément la page
	page--
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func DisplayAccueil(w http.ResponseWriter, r *http.Request) { //on affiche l'accueil
	initTemplate.Temp.ExecuteTemplate(w, "accueil", nil)
}

func DisplayPilotes(w http.ResponseWriter, r *http.Request) { // affichage opti des pilotes
	toSend = nil
	var pilotesFound int // var pour opti la recherche
	if !filtre {
		for _, i := range pilotes { // on s'arrete des qu'on trouve 10 qui ont l'ID correspondant à la page sur laquelle on se trouve
			if i.PageId == page {
				toSend = append(toSend, i)
				pilotesFound++
				if pilotesFound == 10 {
					break
				}
			}
		}
	} else {
		for _, i := range filtresPilotes {
			if i.PageId == page {
				toSend = append(toSend, i)
			}
		}
	}
	// else {
	// 	for _, i := range filtresPilotes {
	// 		if i.PageId == page {
	// 			toSend = append(toSend, i)
	// 			pilotesFound++
	// 			if pilotesFound == 10 {
	// 				break
	// 			}
	// 		}
	// 	}
	// }

	initTemplate.Temp.ExecuteTemplate(w, "pilotes", toSend)
}

func FiltresPilotes(w http.ResponseWriter, r *http.Request) {
	filtre = true
	filtresPilotes = nil
	r.ParseForm()
	query := r.Form["const"] // on récupère tout les filtres
	for _, i := range query {
		for _, j := range pilotes {
			if j.ConstructorID == i {
				filtresPilotes = append(filtresPilotes, j)
			}
		}
	}
	filtresPilotes = Pagination(filtresPilotes)
	page = 1
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func DisplayConstructeurs(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "constructeurs", nil)
}

func DisplayCircuits(w http.ResponseWriter, r *http.Request) {
	var data backend.InfoCircuits
	circuits := Circuits(data)
	initTemplate.Temp.ExecuteTemplate(w, "circuits", circuits)
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "login", nil)
}
