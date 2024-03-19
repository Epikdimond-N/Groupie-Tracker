package main

import (
	"f1/routeur"
	initTemplate "f1/temps"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServe()
}
