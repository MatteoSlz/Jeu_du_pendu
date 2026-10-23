package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func score() {

	data_csv, Err_score_CSV := os.OpenFile("src/score.csv", os.O_RDONLY, 0644)
	if Err_score_CSV != nil {
		fmt.Printf("Erreur : %v", Err_score_CSV)
		fmt.Println()
		os.Exit(0)
	}
	defer data_csv.Close()

	texte := csv.NewReader(data_csv)

	lignes, erreur := texte.ReadAll()
	if erreur != nil {
		fmt.Printf("Erreur : %v", erreur)
	}

	if len(lignes) == 0 {
		afficherTexte50("Vous n'avez pas encore joué de partie.")
		os.Exit(0)
	}

	fmt.Printf("PSEUDO\t STATUT\t\t ESSAIS\t DIFFICULTE\t MOT\t LETTRES\t DATE\n")

	for i := 0; i < len(lignes); i++ {
		nom := lignes[i][0]
		status := lignes[i][1]
		essais := lignes[i][2]
		difficulte := lignes[i][3]
		mot := lignes[i][4]
		lettre := lignes[i][5]
		date := lignes[i][6]

		fmt.Printf("%s\t %s\t %s\t %s\t %s\t\t %s\t %s\n", nom, status, essais,difficulte, mot,lettre, date)
	}
fmt.Println()
afficherTexte50("Voulez vous afficher le classement des 5 meilleurs joueurs ?")
}
