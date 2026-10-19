package main

import (
	"fmt"
	"os"
)

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Veuillez entrer une commande valide : jouer / reprendre / score")
		return 
	}
	choix := os.Args[1]
	switch choix {
	case "jouer":
		jouer()
	case "reprendre":
		reprendre()
	case "score":
		score()
	default:
		fmt.Println("Erreur")
	}
}
