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
	listpilotes = Pagination(listpilotes)
	// for _, m := range listpilotes {
	// 	fmt.Println(m.PageId, m.Name, m.FamilyName, m.Nationality, m.Flag)
	// }
	return listpilotes
}

func Drapeaux(nationality string) string {
	var flag string
	switch nationality {
	case "Dutch":
		flag = "Netherland"
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

func Circuits(data backend.InfoCircuits) []backend.Circuit {
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
	return listcircuits
}

func Pagination(data []backend.Pilote) []backend.Pilote {
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

func SearchPilote(nompilote string) []backend.Pilote {
	var found []backend.Pilote
	nompilote = strings.ToUpper(nompilote)
	for _, i := range pilotes {
		if strings.ToUpper(i.FamilyName) == nompilote || strings.ToUpper(i.Name) == nompilote || strings.ToUpper(i.Name)+" "+strings.ToUpper(i.FamilyName) == nompilote {
			found = append(found, i)
			break
		}
	}
	return found
}
