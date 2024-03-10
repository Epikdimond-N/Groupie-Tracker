package routeur

import (
	controller "f1/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {
	http.HandleFunc("/next", controller.NextPage)
	http.HandleFunc("/previous", controller.PreviousPage)

	http.HandleFunc("/init/pilotes", controller.InitPilotes)

	http.HandleFunc("/", controller.DisplayAccueil)
	http.HandleFunc("/pilotes", controller.DisplayPilotes)
	http.HandleFunc("/constructeurs", controller.DisplayConstructeurs)
	http.HandleFunc("/circuits", controller.DisplayCircuits)
	http.HandleFunc("/login", controller.DisplayLogin)
	http.HandleFunc("/notfound", controller.DisplayNotFound)

	http.HandleFunc("/filtres/pilotes", controller.FiltresPilotes)

	http.HandleFunc("/search/pilotes", controller.DisplayPiloteSearch)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8082 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8082", nil)
}
