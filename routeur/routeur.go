package routeur

import (
	controller "f1/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {
	http.HandleFunc("/pilotes", controller.DisplayPilotes)
	http.HandleFunc("/constructeurs", controller.DisplayConstructeurs)
	http.HandleFunc("/circuits", controller.DisplayCircuits)
	http.HandleFunc("/login", controller.DisplayLogin)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
