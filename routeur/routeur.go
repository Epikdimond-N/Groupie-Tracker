package routeur

import (
	controller "f1/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServe() {

	http.HandleFunc("/next/pilotes", controller.NextPagePilote) // routes pour changer de pages dans pilotes, circuits et constructeurs
	http.HandleFunc("/previous/pilotes", controller.PreviousPagePilote)
	http.HandleFunc("/next/circuits", controller.NextPageCircuit)
	http.HandleFunc("/previous/circuits", controller.PreviousPageCircuit)
	http.HandleFunc("/next/constructeurs", controller.NextPageConstructeur)
	http.HandleFunc("/previous/constructeurs", controller.PreviousPageConstructeur)

	http.HandleFunc("/init/pilotes", controller.InitPilotes) // routes pour initialiser les pilotes, circuits et constructeurs
	http.HandleFunc("/init/circuits", controller.InitCircuits)
	http.HandleFunc("/init/constructeurs", controller.InitConstructeurs)

	http.HandleFunc("/", controller.DisplayAccueil) // route pour aller à l'accueil

	http.HandleFunc("/pilotes", controller.DisplayPilotes) // routes pour voir l'ensemble des pilotes, circuits ou constructeurs
	http.HandleFunc("/constructeurs", controller.DisplayConstructeurs)
	http.HandleFunc("/circuits", controller.DisplayCircuits)

	http.HandleFunc("/login", controller.DisplayLogin) // route pour se log

	http.HandleFunc("/pilote/notfound", controller.DisplayPiloteNotFound)   // routes quand les filtres appliqués ne correspondent pas a un pilote,
	http.HandleFunc("/circuit/notfound", controller.DisplayCircuitNotFound) // un circuit ou un constructeur
	http.HandleFunc("/constructeur/notfound", controller.DisplayConstructeurNotFound)

	http.HandleFunc("/filtres/pilotes", controller.FiltresPilotes) // route sur laquelle les filtres sont appliqués
	http.HandleFunc("/filtres/circuits", controller.FiltresCircuits)
	http.HandleFunc("/filtres/constructeurs", controller.FiltresConstructeurs)

	http.HandleFunc("/search/pilotes", controller.DisplayPiloteSearch) // routes sur lesquels la saisie est traitée et comparé pour renvoyer le constructeur, le pilote ou le circuit correspondant
	http.HandleFunc("/search/circuits", controller.DisplayCircuitSearch)
	http.HandleFunc("/search/constructeurs", controller.DisplayConstructeursSearch)

	http.HandleFunc("/detail/pilote", controller.DisplayDetailPilote) // route sur lesquelles les templates detail de chaques catégories vont etres affichées
	http.HandleFunc("/detail/circuit", controller.DisplayDetailCircuit)
	http.HandleFunc("/detail/constructeur", controller.DisplayDetailConstructeur)

	http.HandleFunc("/back/pilotes", controller.BackToPilotes)   // routes qui va servir a retourner en arrire quand on est sur
	http.HandleFunc("/back/circuits", controller.BackToCircuits) // un template detail sans changer la page sur laquelle on se trouvait précedemment
	http.HandleFunc("/back/constructeurs", controller.BackToConstructeurs)

	http.HandleFunc("/add/pilote_to_favoris", controller.AddPiloteToFavoris)
	http.HandleFunc("/add/circuit_to_favoris", controller.AddCircuitToFavoris)
	http.HandleFunc("/add/constructeur_to_favoris", controller.AddConstructeurToFavoris)

	http.HandleFunc("/remove/pilote_of_favoris", controller.RemovePiloteOfFavoris)
	http.HandleFunc("/remove/circuit_of_favoris", controller.RemoveCircuitOfFavoris)
	http.HandleFunc("/remove/constructeur_of_favoris", controller.RemoveConstructorOfFavoris)

	http.HandleFunc("/favoris", controller.DisplayFavoris)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("\nLien vers le site : http://localhost:8080 (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080", nil)
}
