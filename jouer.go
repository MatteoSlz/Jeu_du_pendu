package main

import (
	"fmt"
	"os"
	"math/rand/v2"
)

func (p *Sauvegarde) jouer() {
	fmt.Println("jouer")
	liste_mot := "mot_correct.txt"
	contenu_byte, err := os.ReadFile(liste_mot)
	if err != nil {
		fmt.Println("", err)
		return
	}
	contenu_string := string(contenu_byte)
	var tableau []string
	mot := ""
	for _, s := range contenu_string {
		if  s != '\n' {
			mot = mot + string(s)
		} else {
			if mot != "" {
				tableau = append(tableau, mot)
				mot = ""
			}
		}
	}
}