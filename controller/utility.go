package controller

import (
	"encoding/json"
	"f1/backend"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Pilotes(data backend.InfoPilotes) []backend.Pilote {

	var listpilotes []backend.Pilote
	var PageId = 1
	var PageIdUsed int

	for i := 2023; i >= 2011; i-- {
		apiUrl := "http://ergast.com/api/f1/" + strconv.Itoa(i) + "/last/results.json"
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
						break
					}
				}
				if !alreadyFoud {
					var tempData backend.Pilote

					tempData.PageId = PageId
					PageIdUsed++
					if PageIdUsed >= 10 {
						PageId++
						PageIdUsed = 0
					}
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
					listpilotes = append(listpilotes, tempData)
					// fmt.Println(tempData.DriverID, tempData.Nationality)
				}
			}
		}
	}

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
	case "Venezuelan":
		flag = "Venezuela"
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
	case "Venezuelian":
		flag = "Venezuela"
	case "Swiss":
		flag = "Switzerland"
	case "New Zealander":
		flag = "New_Zealand"
	}
	return flag
}

func Textify() {
	var data backend.InfoPilotes
	listpilotes := Pilotes(data)

	// Ouvrir le fichier en écriture
	file, err := os.Create("pays.txt")
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

	fmt.Println("Les ID des pilotes ont été écrits dans le fichier ids.txt")
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
