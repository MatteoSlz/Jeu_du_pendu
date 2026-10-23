package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		afficherTexte50("Veuillez entrer une commande valide : jouer / reprendre / score")
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
		afficherTexte50("Choix invalide, veuillez choisir entre jouer/reprendre/score ")
	}
}
