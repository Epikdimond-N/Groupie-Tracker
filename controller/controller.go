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
var constructeurs []backend.Constructeur

var filtresPilotes []backend.Pilote
var filtresCircuits []backend.Circuit
var filtresConstructeurs []backend.Constructeur

var filtre bool
var toSendPilotes []backend.Pilote             // liste des pilotes que je vais envoyer
var toSendCircuits []backend.Circuit           // liste des circuits que je vais envoyer
var toSendConstructeurs []backend.Constructeur // liste des circuits que je vais envoyer

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

func InitConstructeurs(w http.ResponseWriter, r *http.Request) {
	page = 1
	if constructeurs != nil {
		http.Redirect(w, r, "/constructeurs", http.StatusSeeOther)
	} else { // sinon on fait la requete
		constructeurs = Constructeurs()
		filtre = false
		http.Redirect(w, r, "/constructeurs", http.StatusSeeOther)
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

func NextPageConstructeur(w http.ResponseWriter, r *http.Request) {
	page++
	http.Redirect(w, r, "/constructeurs", http.StatusSeeOther)
}

func PreviousPageConstructeur(w http.ResponseWriter, r *http.Request) {
	page--
	http.Redirect(w, r, "/constructeurs", http.StatusSeeOther)
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
	} else {
		for _, i := range filtresCircuits {
			if i.PageId == page {
				toSendCircuits = append(toSendCircuits, i)
			}
		}
	}
	initTemplate.Temp.ExecuteTemplate(w, "circuits", toSendCircuits)
}

func DisplayConstructeurs(w http.ResponseWriter, r *http.Request) {
	toSendConstructeurs = nil
	var constructeursFound int
	if !filtre {
		for _, i := range constructeurs {
			if i.PageId == page {
				toSendConstructeurs = append(toSendConstructeurs, i)
				constructeursFound++
				if constructeursFound == 10 {
					break
				}
			}
		}
	} else {
		for _, i := range filtresConstructeurs {
			if i.PageId == page {
				toSendConstructeurs = append(toSendConstructeurs, i)
			}
		}
	}
	initTemplate.Temp.ExecuteTemplate(w, "constructeurs", toSendConstructeurs)
}

func DisplayLogin(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "login", nil)
}

func DisplayPiloteNotFound(w http.ResponseWriter, r *http.Request) {
	filtresPilotes = pilotes
	initTemplate.Temp.ExecuteTemplate(w, "pilote_not_found", nil)
}

func DisplayCircuitNotFound(w http.ResponseWriter, r *http.Request) {
	filtresCircuits = circuits
	initTemplate.Temp.ExecuteTemplate(w, "circuit_not_found", nil)
}

func DisplayConstructeurNotFound(w http.ResponseWriter, r *http.Request) {
	filtresConstructeurs = constructeurs
	initTemplate.Temp.ExecuteTemplate(w, "constructeur_not_found", nil)
}

func DisplayPiloteSearch(w http.ResponseWriter, r *http.Request) {
	nompilote := r.FormValue("search")
	filtresPilotes = SearchPilote(nompilote)
	filtresPilotes = PaginationPilote(filtresPilotes)
	if filtresPilotes != nil {
		initTemplate.Temp.ExecuteTemplate(w, "pilotes", filtresPilotes)
	} else {
		http.Redirect(w, r, "/pilote/notfound", http.StatusSeeOther)
	}
}

func DisplayCircuitSearch(w http.ResponseWriter, r *http.Request) {
	nomcircuit := r.FormValue("search")
	filtresCircuits = SearchCircuit(nomcircuit)
	filtresCircuits = PaginationCircuits(filtresCircuits)
	if filtresCircuits != nil {
		initTemplate.Temp.ExecuteTemplate(w, "circuits", filtresCircuits)
	} else {
		http.Redirect(w, r, "/circuit/notfound", http.StatusSeeOther)
	}
}

func DisplayConstructeursSearch(w http.ResponseWriter, r *http.Request) {
	nomconstructeur := r.FormValue("search")
	filtresConstructeurs = SearchConstructeur(nomconstructeur)
	filtresConstructeurs = PaginationConstructeurs(filtresConstructeurs)
	if filtresConstructeurs != nil {
		initTemplate.Temp.ExecuteTemplate(w, "constructeurs", filtresConstructeurs)
	} else {
		http.Redirect(w, r, "/constructeur/notfound", http.StatusSeeOther)
	}
}

func AddCircuitToFavoris(w http.ResponseWriter, r *http.Request) {
	var circuitToAdd backend.Circuit
	toAdd := r.URL.Query().Get("idcircuit")

	for _, i := range circuits {
		if i.IDCircuit == toAdd {
			circuitToAdd = i
			break
		}
	}

	fichier, err := os.OpenFile("fav_circuits.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer fichier.Close()

	var favoris []backend.Circuit

	if err := json.NewDecoder(fichier).Decode(&favoris); err != nil && err.Error() != "EOF" {
		fmt.Println("Erreur lors de la lecture du fichier JSON :", err)
		return
	}
	circuitToAdd = TexteCircuit(circuitToAdd)
	favoris = append(favoris, circuitToAdd)

	updatedJSON, err := json.MarshalIndent(favoris, "", "  ")
	if err != nil {
		fmt.Println("il y a une erreur", err)
		return
	}

	err = os.WriteFile("fav_circuits.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println("il y a une erreur", err)
		return
	}

	http.Redirect(w, r, "/circuits", http.StatusSeeOther)
}

func AddPiloteToFavoris(w http.ResponseWriter, r *http.Request) {
	var piloteToAdd backend.Pilote
	toAdd := r.URL.Query().Get("idpilote")

	for _, i := range pilotes {
		if i.DriverID == toAdd {
			piloteToAdd = i
			break
		}
	}

	fichier, err := os.OpenFile("fav_pilotes.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer fichier.Close()

	var favoris []backend.Pilote

	if err := json.NewDecoder(fichier).Decode(&favoris); err != nil && err.Error() != "EOF" {
		fmt.Println("Erreur lors de la lecture du fichier JSON :", err)
		return
	}
	piloteToAdd = TextePilote(piloteToAdd)
	favoris = append(favoris, piloteToAdd)

	updatedJSON, err := json.MarshalIndent(favoris, "", "  ")
	if err != nil {
		fmt.Println("Il y a une erreur", err)
		return
	}

	err = os.WriteFile("fav_pilotes.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println("Il y a une erreur", err)
		return
	}

	http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
}

func AddConstructeurToFavoris(w http.ResponseWriter, r *http.Request) {
	var constructeurToAdd backend.Constructeur
	toAdd := r.URL.Query().Get("idconstructeur")
	fmt.Println(toAdd)
	for _, i := range constructeurs {
		fmt.Println(i.ConstructorId)
		if i.ConstructorId == toAdd {

			constructeurToAdd = i
			break
		}
	}

	fichier, err := os.OpenFile("fav_constructeurs.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer fichier.Close()

	var favoris []backend.Constructeur

	if err := json.NewDecoder(fichier).Decode(&favoris); err != nil && err.Error() != "EOF" {
		fmt.Println("Erreur lors de la lecture du fichier JSON :", err)
		return
	}
	constructeurToAdd = TexteConstructeur(constructeurToAdd)
	favoris = append(favoris, constructeurToAdd)

	updatedJSON, err := json.MarshalIndent(favoris, "", "  ")
	if err != nil {
		fmt.Println("il y a une erreur", err)
		return
	}

	err = os.WriteFile("fav_constructeurs.json", updatedJSON, 0644)
	if err != nil {
		fmt.Println("il y a une erreur", err)
		return
	}

	http.Redirect(w, r, "/constructeurs", http.StatusSeeOther)
}

func DisplayDetailPilote(w http.ResponseWriter, r *http.Request) {
	var PiloteDetail backend.Pilote
	toAdd := r.URL.Query().Get("idpilote")
	for _, i := range pilotes {
		if i.DriverID == toAdd {
			PiloteDetail = i
			break
		}
	}
	PiloteDetail = TextePilote(PiloteDetail)
	initTemplate.Temp.ExecuteTemplate(w, "detail_pilote", PiloteDetail)
}

func DisplayDetailCircuit(w http.ResponseWriter, r *http.Request) {
	var CircuitDetail backend.Circuit
	toAdd := r.URL.Query().Get("idcircuit")
	for _, i := range circuits {
		if i.IDCircuit == toAdd {
			CircuitDetail = i
			break
		}
	}
	CircuitDetail = TexteCircuit(CircuitDetail)
	initTemplate.Temp.ExecuteTemplate(w, "detail_circuit", CircuitDetail)
}

func DisplayDetailConstructeur(w http.ResponseWriter, r *http.Request) {
	var ConstructeurDetail backend.Constructeur
	toAdd := r.URL.Query().Get("idconstructeur")
	for _, i := range constructeurs {
		if i.ConstructorId == toAdd {
			ConstructeurDetail = i
			break
		}
	}
	ConstructeurDetail = TexteConstructeur(ConstructeurDetail)
	initTemplate.Temp.ExecuteTemplate(w, "detail_constructeur", ConstructeurDetail)
}

func BackToCircuits(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "circuits", toSendCircuits)
}

func BackToPilotes(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", toSendPilotes)
}

func BackToConstructeurs(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "constructeurs", toSendConstructeurs)
}
