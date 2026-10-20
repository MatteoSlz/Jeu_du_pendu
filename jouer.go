package main

import (
	"fmt"
	"os"
	"math/rand/v2"
)

func (p *Sauvegarde) jouer() {
	defaite := true 
	count_bonne_lettre := 0 
	type mots struct {
		lettre string
		lettre_trouve bool 
	} 

	fmt.Println("jouer")
	liste_mot := "mot_correct.txt"
	contenu_byte, err := os.ReadFile(liste_mot)
	if err != nil {
		fmt.Println("", err)
		return
	}
	contenu_string := string(contenu_byte)
	var tableau_final []string
	mot := ""
	for _, s := range contenu_string {
		if  s != '\n' {
			mot = mot + string(s)
		} else {
			if mot != "" {
				tableau_final = append(tableau_final, mot)
				mot = ""
			}
		}
	}
	index_mots_random := rand.IntN(len(tableau_final))
	mot_a_trouver := tableau_final[index_mots_random]
	tableau_mot := []mots{}
	fmt.Println("Bienvenue au jeu du pendu !")
	fmt.Printf("Votre mot à trouver contient %s lettres.",len(mot_a_trouver))
	fmt.Println("A toi de jouer !")
	fmt.Println()

	for i := 0 ; i<len(mot_a_trouver);i++{
		tableau_mot[i].lettre_trouve = false 
	}
	index_lettre_donnee_debut := rand.IntN(len(mot_a_trouver))
	lettre_donnee_debut := tableau_mot[index_lettre_donnee_debut].lettre
	for i := 0 ; i < len(tableau_mot);i++{
		if lettre_donnee_debut == tableau_mot[i].lettre {
			tableau_mot[i].lettre_trouve = true
		}
	}
	for defaite {
		count_bonne_lettre = 0 
			for i := 0 ; i<len(mot_a_trouver);i++{
				if tableau_mot[i].lettre_trouve{
					count_bonne_lettre++
				}
			} 
			if count_bonne_lettre == len(mot_a_trouver){
				defaite = false
				break
			}

		for i := 0 ; i<len(mot_a_trouver);i++{
			if tableau_mot[i].lettre_trouve{
				fmt.Printf("%s ",tableau_mot[i].lettre)
			} else {
				fmt.Printf("_ ")
			}
		}
		fmt.Println()
		fmt.Println("Que voulez proposer : ")
		var action string 
		fmt.Scan(&action)
		switch action {
		case "a", "b", "c","d", "e", "f","g", "h", "i","j", "k", "l","m", "n", "o","p", "q", "r","s", "t", "u","v", "w", "x","y", "z" :
			for i := 0 ; i<len(mot_a_trouver);i++ {
				if action == tableau_mot[i].lettre{
					tableau_mot[i].lettre_trouve = true
				}
				fmt.Println()
			}
			fmt.Printf("Vous avez trouvé la lettre :%s.\n",tableau_mot[i].lettre)
		case mot_a_trouver : 
			defaite = false 
		default :
		fmt.Println("Entree invalide, veuillez réessayer.") 
		continue
	}
}
fmt.Println()
fmt.Printf("Bravo ! le mot à trouver était : %s\n",mot_a_trouver)
}
