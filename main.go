package main

import (
	// "f1/controller"
	"f1/routeur"
	initTemplate "f1/temps"
)

func main() {
	// controller.Textify()
	initTemplate.InitTemplate()
	routeur.InitServe()
}
