package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
		if (len(os.Args) != 2) {
			fmt.Println("Veuillez entrer une commande valide : jouer / reprendre / score")
			return
		}
		 
		switch strings.ToLower(os.Args[1])  {
		case "jouer":
			jouer()
		case "reprendre":
			reprendre()
		case "score":
			score()
		default:
			fmt.Println("Erreur: commande indiqué incorrect")
			os.Exit(1)
		}
}
