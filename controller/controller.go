package controller

import (
	"f1/backend"
	initTemplate "f1/temps"
	"fmt"
	"net/http"
)

var page = 1
var pilotes []backend.Pilote

// func InitPagePilotes(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("page"))
// 	if err != nil {
// 		return
// 	}
// 	page = id
// 	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
// }

func NextPage(w http.ResponseWriter, r *http.Request) {
	page++
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func PreviousPage(w http.ResponseWriter, r *http.Request) {
	page--
	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func InitPilotes(w http.ResponseWriter, r *http.Request) {
	if pilotes != nil {
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	} else {
		var data backend.InfoPilotes
		pilotes = Pilotes(data)
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	}
}

func DisplayAccueil(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "accueil", nil)
}

func DisplayPilotes(w http.ResponseWriter, r *http.Request) {
	var toSend []backend.Pilote
	var pilotesFound int
	for _, i := range pilotes {
		if i.PageId == page {
			toSend = append(toSend, i)
			pilotesFound++
			if pilotesFound == 10 {
				break
			}
		}
	}
	// for _, j := range toSend {
	// 	fmt.Println(j.PageId, j.DriverID)
	// }
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", toSend)
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

func FiltresPilotes(w http.ResponseWriter, r *http.Request) {
	var data backend.InfoPilotes
	pilotes = Pilotes(data)
	r.ParseForm()
	query := r.Form["const"]
	fmt.Println(data)
	var filtresPilotes []backend.Pilote
	for _, i := range query {
		for _, j := range pilotes {
			if j.ConstructorID == i {
				filtresPilotes = append(filtresPilotes, j)
			}
		}
	}
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", filtresPilotes)
}
