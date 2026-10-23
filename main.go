package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		afficherTexte50(Red + "Veuillez entrer une commande valide : jouer / reprendre / score" + Reset)
		return
	}
	Recuperation()
	choix := os.Args[1]
	switch choix {
	case "jouer":
		jouer()
	case "reprendre":
		reprendre()
	case "score":
		score()
	default:
		afficherTexte50(Red + "Choix invalide, veuillez choisir entre jouer/reprendre/score " + Reset)
	}
}
