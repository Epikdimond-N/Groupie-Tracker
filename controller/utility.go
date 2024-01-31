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
	apiUrl := "http://ergast.com/api/f1/2023/last/results.json"

	var listpilotes []backend.Pilote

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
			var tempData backend.Pilote

			tempData.DriverID = j.Driver.DriverID
			tempData.Name = j.Driver.GivenName
			tempData.FamilyName = j.Driver.FamilyName
			tempData.DateOfBirth = j.Driver.DateOfBirth
			tempData.Code, _ = strconv.Atoi(j.Driver.Code)
			tempData.Number, _ = strconv.Atoi(j.Driver.PermanentNumber)
			tempData.Nationality = j.Driver.Nationality
			tempData.Constructor = j.Constructor.Name
			listpilotes = append(listpilotes, tempData)
		}
	}
	return listpilotes
}

func Textify() {
	var data backend.InfoPilotes
	listpilotes := Pilotes(data)

	// Ouvrir le fichier en écriture
	file, err := os.Create("ids.txt")
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier:", err)
		return
	}
	defer file.Close()

	// Parcourir la liste de pilotes et écrire les ID dans le fichier
	for _, pilote := range listpilotes {
		_, err := file.WriteString(pilote.DriverID + "\n")
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
			return
		}
	}

	fmt.Println("Les ID des pilotes ont été écrits dans le fichier ids.txt")
}
