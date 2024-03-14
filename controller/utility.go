package controller

import (
	"encoding/json"
	"f1/backend"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Pilotes() []backend.Pilote {
	var data backend.InfoPilotes
	var listpilotes []backend.Pilote

	for saison := 2023; saison >= 2011; saison-- {
		apiUrl := "http://ergast.com/api/f1/" + strconv.Itoa(saison) + "/last/results.json"
		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			fmt.Println("Erreur lors de la création de la requête:", err)
			return listpilotes
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de la requête:", err)
			return listpilotes
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse:", err)
			return listpilotes
		}

		json.Unmarshal(body, &data)

		for _, i := range data.MRData.RaceTable.Races {
			for _, j := range i.Results {
				var alreadyFoud bool
				for _, k := range listpilotes {
					if k.DriverID == j.Driver.DriverID {
						alreadyFoud = true
						for index, l := range listpilotes {
							if l.DriverID == j.Driver.DriverID {
								listpilotes[index].Saison = append(listpilotes[index].Saison, saison)
							}
						}
						break
					}
				}
				if !alreadyFoud {
					var tempData backend.Pilote
					tempData.DriverID = j.Driver.DriverID
					tempData.Name = j.Driver.GivenName
					tempData.FamilyName = j.Driver.FamilyName
					tempData.DateOfBirth = j.Driver.DateOfBirth
					tempData.Code, _ = strconv.Atoi(j.Driver.Code)
					tempData.Number, _ = strconv.Atoi(j.Driver.PermanentNumber)
					tempData.Nationality = j.Driver.Nationality
					tempData.Flag = Drapeaux(j.Driver.Nationality)
					tempData.Constructor = j.Constructor.Name
					tempData.ConstructorID = j.Constructor.ConstructorID
					tempData.Saison = append(tempData.Saison, saison)
					listpilotes = append(listpilotes, tempData)
				}
			}
		}
	}
	listpilotes = PaginationPilote(listpilotes)
	return listpilotes
}

func Drapeaux(nationality string) string {
	var flag string
	switch nationality {
	case "Dutch":
		flag = "Netherlands"
	case "Monegasque":
		flag = "Monaco"
	case "British":
		flag = "UK"
	case "Mexican":
		flag = "Mexico"
	case "Australian":
		flag = "Australia"
	case "Spanish":
		flag = "Spain"
	case "Japanese":
		flag = "Japan"
	case "Canadian":
		flag = "Canada"
	case "French":
		flag = "France"
	case "Thai":
		flag = "Thailand"
	case "German":
		flag = "Germany"
	case "American":
		flag = "USA"
	case "Chinese":
		flag = "China"
	case "Finnish":
		flag = "Finland"
	case "Danish":
		flag = "Denmark"
	case "Brazilian":
		flag = "Brazil"
	case "Polish":
		flag = "Poland"
	case "Indian":
		flag = "India"
	case "Italian":
		flag = "Italy"
	case "Belgian":
		flag = "Belgium"
	case "Russian":
		flag = "Russia"
	case "Swedish":
		flag = "Sweden"
	case "Venezuelan":
		flag = "Venezuela"
	case "Swiss":
		flag = "Switzerland"
	case "New Zealander":
		flag = "New_Zealand"
	}
	return flag
}

func Textify() {
	listpilotes := Pilotes()

	// Ouvrir le fichier en écriture
	file, err := os.Create("nationality.txt")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return
	}
	defer file.Close()

	// Parcourir la liste de pilotes et écrire les ID dans le fichier
	for _, pilote := range listpilotes {
		_, err := file.WriteString(pilote.Nationality + "\n")
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
			return
		}
	}

	fmt.Println("Les ID des pilotes ont été écrits dans le fichier nationality.txt")
}

func Circuits() []backend.Circuit {
	var data backend.InfoCircuits
	var listcircuits []backend.Circuit
	for i := 1950; i <= 2023; i++ {
		apiUrl := "http://ergast.com/api/f1/" + strconv.Itoa(i) + ".json"

		req, err := http.NewRequest("GET", apiUrl, nil)
		if err != nil {
			fmt.Println("Erreur lors de la création de la requête:", err)
			return listcircuits
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de la requête:", err)
			return listcircuits
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erreur lors de la lecture de la réponse:", err)
			return listcircuits
		}
		json.Unmarshal(body, &data)
		for _, j := range data.MRData.RaceTable.Races {
			var isInList bool
			var tempData backend.Circuit
			tempData.IDCircuit = j.Circuit.CircuitID
			tempData.GPName = j.RaceName
			tempData.Name = j.Circuit.CircuitName
			tempData.Pays = j.Circuit.Location.Country
			tempData.Ville = j.Circuit.Location.Locality
			tempData.Seasons = append(tempData.Seasons, j.Season)
			for index, l := range listcircuits {
				if l.IDCircuit == tempData.IDCircuit {
					isInList = true
					listcircuits[index].Seasons = append(listcircuits[index].Seasons, j.Season)
				}
			}
			if !isInList {
				listcircuits = append(listcircuits, tempData)
				isInList = false
			}
		}
	}
	listcircuits = PaginationCircuits(listcircuits)
	return listcircuits
}

// func Constructeurs() []backend.Constructeur {
// 	var data backend.InfoConstructeurs
// 	var listconstructeurs []backend.Constructeur

// 	for saison := 2023; saison >= 2011; saison-- {
// 		apiUrl := "http://ergast.com/api/f1/" + strconv.Itoa(saison) + "/last/results.json"
// 		req, err := http.NewRequest("GET", apiUrl, nil)
// 		if err != nil {
// 			fmt.Println("Erreur lors de la création de la requête:", err)
// 			return listconstructeurs
// 		}

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			fmt.Println("Erreur lors de l'envoi de la requête:", err)
// 			return listconstructeurs
// 		}
// 		defer resp.Body.Close()

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println("Erreur lors de la lecture de la réponse:", err)
// 			return listconstructeurs
// 		}

// 		json.Unmarshal(body, &data)

// 		for _, i := range data.MRData.ConstructorTable.Constructors {
// 			var alreadyFoud bool
// 			for _, j := range listconstructeurs {
// 				if i.ConstructorID == j.ConstructorId {
// 					alreadyFoud = true
// 					break
// 				}
// 			}
// 			if !alreadyFoud {
// 				listconstructeurs = append(listconstructeurs, tempData)
// 			}
// 		}
// 	}

// 	listconstructeurs = PaginationConstructeurs(listconstructeurs)
// 	return listconstructeurs
// }

func PaginationPilote(data []backend.Pilote) []backend.Pilote {
	var listpilotes []backend.Pilote
	var id = 1
	var attributedId int
	for _, i := range data {
		i.PageId = id
		listpilotes = append(listpilotes, i)
		attributedId++
		if attributedId == 10 {
			id++
			attributedId = 0
		}
	}

	var maxpage int
	if len(listpilotes)%10 != 0 {
		maxpage = (len(listpilotes) / 10) + 1
		for index, _ := range listpilotes {
			listpilotes[index].MaxPage = maxpage
		}
	} else {
		maxpage = (len(listpilotes) / 10)
		for index, _ := range listpilotes {
			listpilotes[index].MaxPage = maxpage
		}
	}
	return listpilotes
}

func PaginationCircuits(data []backend.Circuit) []backend.Circuit {
	var listpilotes []backend.Circuit
	var id = 1
	var attributedId int
	for _, i := range data {
		i.PageId = id
		listpilotes = append(listpilotes, i)
		attributedId++
		if attributedId == 10 {
			id++
			attributedId = 0
		}
	}

	var maxpage int
	if len(listpilotes)%10 != 0 {
		maxpage = (len(listpilotes) / 10) + 1
		for index, _ := range listpilotes {
			listpilotes[index].MaxPage = maxpage
		}
	} else {
		maxpage = (len(listpilotes) / 10)
		for index, _ := range listpilotes {
			listpilotes[index].MaxPage = maxpage
		}
	}
	return listpilotes
}

func PaginationConstructeurs(data []backend.Constructeur) []backend.Constructeur { //ne marche pas
	var listconstructeurs []backend.Constructeur
	var id = 1
	var attributedId int
	for _, i := range data {
		i.PageId = id
		listconstructeurs = append(listconstructeurs, i)
		attributedId++
		if attributedId == 10 {
			id++
			attributedId = 0
		}
	}

	var maxpage int
	if len(listconstructeurs)%10 != 0 {
		maxpage = (len(listconstructeurs) / 10) + 1
		for index, _ := range listconstructeurs {
			listconstructeurs[index].MaxPage = maxpage
		}
	} else {
		maxpage = (len(listconstructeurs) / 10)
		for index, _ := range listconstructeurs {
			listconstructeurs[index].MaxPage = maxpage
		}
	}
	return listconstructeurs
}

func SearchPilote(nompilote string) []backend.Pilote {
	var found []backend.Pilote
	nompilote = strings.ToUpper(nompilote)
	for _, i := range pilotes {
		if strings.ToUpper(i.FamilyName) == nompilote || strings.ToUpper(i.Name) == nompilote || strings.ToUpper(i.Name)+" "+strings.ToUpper(i.FamilyName) == nompilote {
			found = append(found, i)
		}
	}
	return found
}

func SearchCircuit(nomcircuit string) []backend.Circuit {
	var found []backend.Circuit
	for _, i := range circuits {
		if Search(i.Name, nomcircuit) {
			found = append(found, i)
		}
	}
	return found
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
	filtresPilotes = PaginationPilote(filtresPilotes)
	page = 1
	if filtresPilotes == nil {
		http.Redirect(w, r, "/pilote/notfound", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/pilotes", http.StatusSeeOther)
	}
}

func FiltresCircuits(w http.ResponseWriter, r *http.Request) {
	filtre = true
	filtresCircuits = nil
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

	filtresPilotes = PaginationPilote(filtresPilotes)
	page = 1
	if filtresPilotes == nil {
		http.Redirect(w, r, "/circuit/notfound", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/circuits", http.StatusSeeOther)
	}
}

func Search(word string, s string) bool {
	return strings.Contains(strings.ToLower(word), strings.ToLower(s))
}
