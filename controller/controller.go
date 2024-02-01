package controller

import (
	"f1/backend"
	initTemplate "f1/temps"
	"net/http"
)

func DisplayPilotes(w http.ResponseWriter, r *http.Request) {
	var data backend.InfoPilotes
	pilotes := Pilotes(data)
	initTemplate.Temp.ExecuteTemplate(w, "pilotes", pilotes)
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
