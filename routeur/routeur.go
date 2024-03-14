package routeur

import (
	controller "f1/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {
	http.HandleFunc("/next/pilotes", controller.NextPagePilote)
	http.HandleFunc("/previous/pilotes", controller.PreviousPagePilote)
	http.HandleFunc("/next/circuits", controller.NextPageCircuit)
	http.HandleFunc("/previous/circuits", controller.PreviousPageCircuit)

	http.HandleFunc("/init/pilotes", controller.InitPilotes)
	http.HandleFunc("/init/circuits", controller.InitCircuits)

	http.HandleFunc("/", controller.DisplayAccueil)
	http.HandleFunc("/pilotes", controller.DisplayPilotes)
	http.HandleFunc("/constructeurs", controller.DisplayConstructeurs)
	http.HandleFunc("/circuits", controller.DisplayCircuits)
	http.HandleFunc("/login", controller.DisplayLogin)
	
	http.HandleFunc("/pilote/notfound", controller.DisplayPiloteNotFound)
	http.HandleFunc("/circuit/notfound", controller.DisplayCircuitNotFound)

	http.HandleFunc("/filtres/pilotes", controller.FiltresPilotes)
	http.HandleFunc("/filtres/circuits", controller.FiltresCircuits)

	http.HandleFunc("/search/pilotes", controller.DisplayPiloteSearch)
	http.HandleFunc("/search/circuits", controller.DisplayCircuitSearch)

	http.HandleFunc("/action/pilote_to_favoris", controller.InitFavoris)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
