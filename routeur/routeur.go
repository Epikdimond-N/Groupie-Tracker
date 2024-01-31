package routeur

import (
	"fmt"
	"net/http"
	"os"
	controller "f1/controller"
)

func InitServe() {
	http.HandleFunc("/", controller.DisplayHome)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
