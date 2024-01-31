package controller

import (
	"f1/backend"
	initTemplate "f1/temps"
	"net/http"
)

func DisplayHome(w http.ResponseWriter, r *http.Request) {
	var data backend.InfoPilotes
	pilotes := Pilotes(data)
	initTemplate.Temp.ExecuteTemplate(w, "index", pilotes)
}
