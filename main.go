package main

import (
	"fmt"
	"os"
)

func  Start() {
	var nom string
	fmt.Println()
	fmt.Println("Comment vous appelez vous ?")
	fmt.Scan(&nom)
	fmt.Printf("D'accord bienvenue %s !", nom)
	fmt.Println()
}

func main() {
	
	Start()

	if (len(os.Args) != 2) {
		fmt.Println("Veuillez entrer une commande valide : jouer / reprendre / score")
		return 
	}
	choix := os.Args[1]
	switch choix {
	case "jouer":
		jouer()
	default:
		fmt.Println("Erreur")
	}
}


