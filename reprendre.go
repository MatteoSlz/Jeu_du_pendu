package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
)


func reprendre() {

data, _ := os.ReadFile("save.json")
	var liste_joueurs_json []Sauvegarde
 json.Unmarshal(data,&liste_joueurs_json) //erreur a traiter

var nom string
choix_joueurs := true
numero_joueur := -1

for choix_joueurs {
	fmt.Println()
	fmt.Printf("Quel partie souhaite-tu reprendre : ")
	fmt.Println()
	for i := 0 ; i< len(liste_joueurs_json);i++{
		fmt.Println()
		fmt.Printf("%s",liste_joueurs_json[i].Pseudo)
	}
	fmt.Println()
	fmt.Printf(": ")
	fmt.Scan(&nom)
	for j := 0 ; j<len(liste_joueurs_json);j++{
		if  nom == liste_joueurs_json[j].Pseudo{
			numero_joueur = j
			choix_joueurs = false 
			break
		} 
	}
	if numero_joueur == -1 {
		fmt.Println()
		fmt.Println()
		fmt.Println("Ce pseudo n'existe pas. Veuillez saisir un pseudo valide.")
	}
}

perdu := false 
	Essais_restant := 0
	mot_a_trouver := ""
	tableau_mot := []mots{}
	defaite := true
	s := ""
	count_bonne_lettre := 0

	if numero_joueur < len(liste_joueurs_json){
	fmt.Println()
	fmt.Printf("D'accord re bonjour %s !", liste_joueurs_json[numero_joueur].Pseudo)
	fmt.Println()

	Essais_restant = liste_joueurs_json[numero_joueur].Essais
	mot_a_trouver = liste_joueurs_json[numero_joueur].Mot
	tableau_mot = liste_joueurs_json[numero_joueur].Lettres
	s = liste_joueurs_json[numero_joueur].Pseudo
	}


	for defaite {
		count_bonne_lettre = 0
		for i := 0; i < len(mot_a_trouver); i++ {
			if tableau_mot[i].Lettre_trouve {
				count_bonne_lettre++
			}
		}
		if count_bonne_lettre == len(mot_a_trouver) {
			defaite = false
			break
		}
		fmt.Println()
		for i := 0; i < len(mot_a_trouver); i++ {
			if tableau_mot[i].Lettre_trouve {
				fmt.Printf("%s ", tableau_mot[i].Lettre)
			} else if tableau_mot[i].Lettre == "-" {
				fmt.Printf("- ")
				tableau_mot[i].Lettre_trouve = true 

			}else{
				fmt.Printf("_ ")
			}
		}
		fmt.Println()
		fmt.Printf("Il vous reste %d essai(s)",Essais_restant)
		fmt.Println()
		fmt.Print("Que voulez-vous faire : ")
		var action string
		fmt.Scan(&action)
		action = strings.ToLower(action)
		switch action {
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			lettre_trouvee := false
			for i := 0; i < len(mot_a_trouver); i++ {
				if action == tableau_mot[i].Lettre {
					tableau_mot[i].Lettre_trouve = true
					lettre_trouvee = true
				}
			}
			if lettre_trouvee {
				fmt.Printf("Bravo ! Vous avez trouvé la lettre : %s.\n", action)
			} else {
				fmt.Printf("Dommage. La lettre '%s' n'est pas dans le mot.", action)
				Essais_restant -- 
			}
			fmt.Println()
		case mot_a_trouver:
			if Essais_restant >= 2 {
			defaite = false
		}else{
			fmt.Println()
			fmt.Println("Il ne vous reste pas assez d'essais pour effectuer cette action.")
			continue
		}
	case "!save" :
		fmt.Println()
		fmt.Println("Sauvegarde de votre partie en cours . . . . . . . . . . .")
		fmt.Println()
		fmt.Println("Sauvegarde terminée. Nous avons hâte de vous revoir!")
		fmt.Println()

	player := Sauvegarde{
		Pseudo : s,
		Mot : mot_a_trouver,
		Lettres : tableau_mot,
		Essais : Essais_restant,
	}
	
	data , _ := os.ReadFile("save.json")
	var liste_joueurs_json []Sauvegarde
   json.Unmarshal(data,&liste_joueurs_json)

	liste_joueurs = append(liste_joueurs_json, player )

	encodeData, encodeErr := json.Marshal(liste_joueurs)

	if encodeErr != nil {
		fmt.Printf("Erreur encodage données : %v", encodeErr)
		return
	}
	fileErr := os.WriteFile("save.json", encodeData, 0644)
	if fileErr != nil {
		fmt.Printf("Erreur encodage données : %v", fileErr)
		return
	}
	return
		default:
			if (len(mot_a_trouver) == len(action)) && (Essais_restant >=2){
				valeur := 0
				for i := 0 ; i <len(action);i++{
					if string(action[i])>="a" && string(action[i])<="z"{
						valeur ++
					} 
				}
				if valeur == len(action){
					fmt.Println()
					fmt.Println("Le mot n'est pas bon, vous perdez 2 essais.")
					fmt.Println()
					Essais_restant -= 2
				}

			} else if (len(mot_a_trouver) == len(action)) && (Essais_restant <2){
				fmt.Println("Il ne vous reste pas assez d'essais pour effectuer cette action.")
			fmt.Println()
			continue
			} else {
			fmt.Println("Entrée invalide, veuillez proposer une seule lettre ou un mot de la bonne longueur.")
			fmt.Println()
			continue
			}
		}
		if Essais_restant == 0 {
			perdu = true
			defaite = false 
		}
	}
	if perdu {
	fmt.Printf("Dommage c'est perdu. Le mot à trouver était : %s\n", mot_a_trouver) 
	} else {
	fmt.Println()
	fmt.Printf("Bravo! Le mot à trouver était bien : %s\n", mot_a_trouver) 
}
}