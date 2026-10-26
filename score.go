package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type classements struct {
	Pseudo           string
	difficulte       string
	Essais           int
	difficulteValeur int
}

var tableau = make([]classements, 5)

func classementScore(pseudo string, difficulte string, essais string) []classements {

	d := 0
	switch difficulte {
	case "facile":
		d = 1
	case "normal":
		d = 2
	case "difficile":
		d = 3
	case "hardcore":
		d = 4
	default:
		fmt.Println("Erreur")
	}

	essaisInt, _ := strconv.Atoi(essais)

	var player = classements{
		Pseudo:           pseudo,
		difficulte:       difficulte,
		Essais:           essaisInt,
		difficulteValeur: d,
	}

	for i := 0; i < 5; i++ {
		placeEstMoinsBonne := tableau[i].difficulteValeur == 0 || tableau[i].Essais == 0
		joueurEstMeilleur := player.difficulteValeur > tableau[i].difficulteValeur || (player.difficulteValeur == tableau[i].difficulteValeur && player.Essais > tableau[i].Essais)

		if player.Essais > 0 && (placeEstMoinsBonne || joueurEstMeilleur) {
			for j := 4; j > i; j-- {
				tableau[j] = tableau[j-1]
			}
			tableau[i] = player
			break
		} else if player.Essais == 0 && tableau[i].difficulteValeur == 0 {
			for j := 4; j > i; j-- {
				tableau[j] = tableau[j-1]
			}
			tableau[i] = player
			break
		}
	}
	return tableau
}

func score() {

	data_csv, Err_score_CSV := os.OpenFile("src/score.csv", os.O_RDONLY, 0644)
	if Err_score_CSV != nil {
		fmt.Printf("Erreur : %v\n", Err_score_CSV)
		os.Exit(0)
	}
	defer data_csv.Close()

	texte := csv.NewReader(data_csv)

	lignes, erreur := texte.ReadAll()
	if erreur != nil {
		fmt.Printf("Erreur : %v\n", erreur)
	}

	if len(lignes) == 0 {
		afficherTexte50("Vous n'avez pas encore joué de partie.")
		os.Exit(0)
	}

	formatLigne := "%-14s %-11s %-8s %-11s %-24s %-20s %s\n"

	fmt.Printf(formatLigne, "PSEUDO", "STATUT", "ESSAIS", "DIFFICULTE", "MOT", "LETTRES", "DATE")

	for i := 0; i < len(lignes); i++ {
		nom := lignes[i][0]
		status := lignes[i][1]
		essais := lignes[i][2]
		difficulte := lignes[i][3]
		mot := lignes[i][4]
		lettre := lignes[i][5]
		date := lignes[i][6]

		tableau = classementScore(nom, difficulte, essais)

		fmt.Printf(formatLigne, nom, status, essais, difficulte, mot, lettre, date)
	}

	fmt.Println()
	var reponseClassement = true

	for reponseClassement {
		afficherTexte50("Voulez vous afficher le classement des 5 meilleurs joueurs ?")
		afficherTexte50l("oui/non : ")

		var reponse string
		fmt.Scan(&reponse)
		switch reponse {
		case "oui":
			fmt.Println()
			formatEntete := "%-14s %-14s %-12s %-8s\n"
			formatLigne2 := "%-14s %-14s %-12s %-8d\n"

			fmt.Printf(formatEntete, "Classement", "PSEUDO", "DIFFICULTE", "ESSAIS")
			fmt.Println("")
			fmt.Printf(formatLigne2, "#1", tableau[0].Pseudo, tableau[0].difficulte, tableau[0].Essais)
			fmt.Printf(formatLigne2, "#2", tableau[1].Pseudo, tableau[1].difficulte, tableau[1].Essais)
			fmt.Printf(formatLigne2, "#3", tableau[2].Pseudo, tableau[2].difficulte, tableau[2].Essais)
			fmt.Printf(formatLigne2, "#4", tableau[3].Pseudo, tableau[3].difficulte, tableau[3].Essais)
			fmt.Printf(formatLigne2, "#5", tableau[4].Pseudo, tableau[4].difficulte, tableau[4].Essais)
			return
		case "non":
			afficherTexte50("Daccord à bientôt !")
			return
		default:
			afficherTexte50("Veuillez saisir une entrée valide. (oui/non)")
		}
	}
}
