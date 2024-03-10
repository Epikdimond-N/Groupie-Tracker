package controller

import (
	"f1/backend"
	initTemplate "f1/temps"
	"net/http"
	"strconv"
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
	filtre = false
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
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", toSend)
}

func FiltresPilotes(w http.ResponseWriter, r *http.Request) {
	filtre = true
	filtresPilotes = nil
	r.ParseForm()
	queryconst := r.Form["const"]                                    // on récupère tout les filtres constructeurs
	queryflag := r.Form["flag"]                                      // on récupère tout les filtres drapeaux
	querysaison := r.Form["saison"]                                  // on récupère tout les filtres saisons
	if queryconst == nil && queryflag == nil && querysaison == nil { // si il n'y en a aucun
		filtresPilotes = pilotes
	} else {
		if queryconst != nil && queryflag == nil && querysaison == nil { // si ya que constructeur de choisi
			for _, i := range queryconst {
				for _, j := range pilotes {
					if j.ConstructorID == i {
						filtresPilotes = append(filtresPilotes, j)
					}
				}
			}
		} else if queryconst != nil && queryflag != nil && querysaison == nil { // si ya constructeur + flag
			for _, constructeur := range queryconst {
				for _, flag := range queryflag {
					for _, pilote := range pilotes {
						if pilote.Nationality == flag && pilote.ConstructorID == constructeur {
							var piloteAlreadyFound bool
							for _, n := range filtresPilotes {
								if n.DriverID == pilote.DriverID {
									piloteAlreadyFound = true
								}
							}
							if !piloteAlreadyFound {
								filtresPilotes = append(filtresPilotes, pilote)
							}

						}
					}
				}
			}
		} else if queryconst != nil && queryflag == nil && querysaison != nil { // si ya constructeur + saison
			for _, constructeur := range queryconst {
				for _, saison := range querysaison {
					for _, pilote := range pilotes {
						for _, l := range pilote.Saison {
							if strconv.Itoa(l) == saison && pilote.ConstructorID == constructeur {
								var piloteAlreadyFound bool
								for _, n := range filtresPilotes {
									if n.DriverID == pilote.DriverID {
										piloteAlreadyFound = true
									}
								}
								if !piloteAlreadyFound {
									filtresPilotes = append(filtresPilotes, pilote)
								}
							}
						}

					}
				}
			}
		} else if queryconst == nil && queryflag != nil && querysaison != nil { // si ya flag + saison
			for _, flag := range queryflag {
				for _, saison := range querysaison {
					for _, pilote := range pilotes {
						for _, l := range pilote.Saison {
							if strconv.Itoa(l) == saison && pilote.Nationality == flag {
								var piloteAlreadyFound bool
								for _, n := range filtresPilotes {
									if n.DriverID == pilote.DriverID {
										piloteAlreadyFound = true
									}
								}
								if !piloteAlreadyFound {
									filtresPilotes = append(filtresPilotes, pilote)
								}

							}
						}

					}
				}
			}
		} else if queryconst == nil && queryflag != nil && querysaison == nil { // si ya que flag
			for _, flag := range queryflag {
				for _, pilote := range pilotes {
					if pilote.Nationality == flag {
						var piloteAlreadyFound bool
						for _, i := range filtresPilotes {
							if i.DriverID == pilote.DriverID {
								piloteAlreadyFound = true
							}
						}
						if !piloteAlreadyFound {
							filtresPilotes = append(filtresPilotes, pilote)
						}

					}
				}
			}
		} else if queryconst == nil && queryflag == nil && querysaison != nil { // si ya que saison
			for _, i := range querysaison { // on range les filtres
				for _, j := range pilotes { // on range les pilotes
					for _, k := range j.Saison { //on range les saisons d'un pilote
						if strconv.Itoa(k) == i {
							var piloteAlreadyFound bool
							for _, l := range filtresPilotes {
								if l.DriverID == j.DriverID {
									piloteAlreadyFound = true
								}
							}
							if !piloteAlreadyFound {
								filtresPilotes = append(filtresPilotes, j)
							}
						}
					}
				}
			}
		} else if queryconst != nil && queryflag != nil && querysaison != nil { // si il y sont tous
			for _, flag := range queryflag { // drapeaux
				for _, constructeur := range queryconst { // constructeurs
					for _, saison := range querysaison { //saisons
						for _, pilote := range pilotes { // on parcour les pilotes
							for _, m := range pilote.Saison { // on parcour les saisons du pilotes
								if strconv.Itoa(m) == saison && flag == pilote.Nationality && constructeur == pilote.ConstructorID {
									var piloteAlreadyFound bool
									for _, n := range filtresPilotes {
										if n.DriverID == pilote.DriverID {
											piloteAlreadyFound = true
										}
									}
									if !piloteAlreadyFound {
										filtresPilotes = append(filtresPilotes, pilote)
									}

								}
							}

						}
					}
				}
			}

		}
	}

	filtresPilotes = Pagination(filtresPilotes)
	page = 1
	if filtresPilotes == nil {
		http.Redirect(w, r, "/notfound", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	}
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

func DisplayNotFound(w http.ResponseWriter, r *http.Request) {
	filtresPilotes = pilotes
	initTemplate.Temp.ExecuteTemplate(w, "notfound", nil)
}

func DisplayPiloteSearch(w http.ResponseWriter, r *http.Request) {
	nompilote := r.FormValue("search")
	filtresPilotes = SearchPilote(nompilote)
	filtresPilotes = Pagination(filtresPilotes)
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", filtresPilotes)
}
