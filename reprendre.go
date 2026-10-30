package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var DifficulteReprendre string

func reprendre() {

	var liste_joueurs_json []Sauvegarde

	data, Err := os.ReadFile("src/save.json")
	if Err != nil {
		if errors.Is(Err, os.ErrNotExist) {
			fmt.Println("Erreur")
			os.WriteFile("src/save.json", []byte("[]"), 0644)
			data = []byte("[]")
		} else {
			fmt.Printf("Erreur : %v\n", Err)
			os.Exit(0)
		}
	}

	Err_json := json.Unmarshal(data, &liste_joueurs_json)
	if Err_json != nil {
		errStr := Err_json.Error()
		if errStr != "unexpected end of JSON input" && errStr != "EOF" {
			fmt.Printf("Erreur : %v", Err_json)
			os.Exit(0)
		}
	}
	if len(liste_joueurs_json) == 0 {
		afficherTexte50(Orange + "Aucune partie sauvegardée trouvée." + Reset)
		os.Exit(0)
	}
	var nom string
	choix_joueurs := true
	numero_joueur := -1

	for choix_joueurs {
		fmt.Println()
		afficherTexte50l(Magenta + "Quel partie souhaite-tu reprendre ? " + Reset)
		for i := 0; i < len(liste_joueurs_json); i++ {
			fmt.Println()
			afficherTexte50l(fmt.Sprintf(Orange+"%s"+Reset, liste_joueurs_json[i].Pseudo))
		}
		fmt.Println()
		afficherTexte50l(Magenta + ": " + Reset)
		fmt.Scan(&nom)
		for j := 0; j < len(liste_joueurs_json); j++ {
			if nom == liste_joueurs_json[j].Pseudo {
				numero_joueur = j
				choix_joueurs = false
				break
			}
		}
		if numero_joueur == -1 {
			fmt.Println()
			fmt.Println()
			afficherTexte50(Red + "Ce pseudo n'existe pas. Veuillez saisir un pseudo valide." + Reset)
		}
	}

	perdu := false
	Essais_restant := 0
	mot_a_trouver := ""
	tableau_mot := []mots{}
	defaite := true
	s := ""
	count_bonne_lettre := 0

	if numero_joueur < len(liste_joueurs_json) {
		fmt.Println()
		afficherTexte50l(fmt.Sprintf(Magenta+"D'accord re bonjour %s !"+Reset, liste_joueurs_json[numero_joueur].Pseudo))
		fmt.Println()

		DifficulteReprendre = liste_joueurs_json[numero_joueur].Difficultesave
		Essais_restant = liste_joueurs_json[numero_joueur].Essais
		mot_a_trouver = liste_joueurs_json[numero_joueur].Mot
		tableau_mot = liste_joueurs_json[numero_joueur].Lettres
		s = liste_joueurs_json[numero_joueur].Pseudo
		alphabet = liste_joueurs_json[numero_joueur].Alphabets
	}

	for defaite {
		stopSound()
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
				afficherTexte50l(fmt.Sprintf(Green+"%s "+Reset, tableau_mot[i].Lettre))
			} else if tableau_mot[i].Lettre == "-" {
				afficherTexte50l(couleur + "- " + Reset)
				tableau_mot[i].Lettre_trouve = true

			} else {
				afficherTexte50l("_ ")
			}
		}

		afficherTexte50l(fmt.Sprintf(Yellow+"\tIl vous reste %d essai(s)"+Reset, Essais_restant))
		fmt.Println()

if Essais_restant == 5 {
			fmt.Printf("%-40s%s", "", "\n")
			fmt.Printf("%-40s%s", "", "\n")
			fmt.Printf("%-40s%s", "", "\n")
			fmt.Printf("%-40s%s", "", "\n")
			fmt.Printf("%-40s%s", "", "\n")
			fmt.Printf("%-40s%s", "", "__|__ \n")
		}
		if Essais_restant == 4 {
			fmt.Printf("%-40s%s", "", "      \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "__|__ \n")
		}
		if Essais_restant == 3 {
			fmt.Printf("%-40s%s", "", "  _______\n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "__|__ \n")
		}
		if Essais_restant == 2 {
			fmt.Printf("%-40s%s", "", "  _______\n")
			fmt.Printf("%-40s%s", "", "  |/   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "__|__ \n")
		}
		if Essais_restant == 1 {
			fmt.Printf("%-40s%s", "", "  _______\n")
			fmt.Printf("%-40s%s", "", "  |/   |\n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "  |   \n")
			fmt.Printf("%-40s%s", "", "__|__ \n")
		}
		fmt.Println()


		afficherTexte50l(couleur + "Lettres proposées : " + Reset)
		for i := 0; i < len(alphabet); i++ {
			if alphabet[i].Lettre_trouve == true {
				dansLeMot := false
				for j := 0; j < len(tableau_mot); j++ {
					if alphabet[i].Lettre == tableau_mot[j].Lettre {
						dansLeMot = true
						break
					}
				}
				if dansLeMot {
					afficherTexte50l(fmt.Sprintf(Green+"%s "+Reset, alphabet[i].Lettre))
				} else {
					afficherTexte50l(fmt.Sprintf(Red+"%s "+Reset, alphabet[i].Lettre))
				}
			}
		}
		fmt.Println()
		fmt.Println()
		afficherTexte50l(couleur + "Que voulez-vous proposer : " + Reset)
		var action string
		fmt.Scan(&action)
		action = strings.ToLower(action)
		switch action {
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z":
			lettre_deja_proposer := false
			for k := 0; k < len(alphabet); k++ {
				if (action == alphabet[k].Lettre) && (alphabet[k].Lettre_trouve) {
					lettre_deja_proposer = true
				}
			}
			if lettre_deja_proposer == false {
				lettre_trouvee := false
				for i := 0; i < len(mot_a_trouver); i++ {
					if action == tableau_mot[i].Lettre {
						tableau_mot[i].Lettre_trouve = true
						lettre_trouvee = true
					}
				}
				for j := 0; j < len(alphabet); j++ {
					if action == alphabet[j].Lettre {
						alphabet[j].Lettre_trouve = true
					}
				}
				if lettre_trouvee {
					PlaySoundValidation()
					fmt.Println()
					afficherTexte50l(fmt.Sprintf(Green+"Bravo ! Vous avez trouvé la lettre : %s.\n"+Reset, action))
					oui = true
					afficherTexteSoundValide("  ")
				} else {
					PlaySoundFaux()
					fmt.Println()
					afficherTexte50l(fmt.Sprintf(Red+"Dommage la lettre '%s' n'est pas dans le mot."+Reset, action))
					fmt.Println()
					oui = false
					Essais_restant--
					afficherTexteSoundFaux(" ")
				}
			} else {
				fmt.Println()
				afficherTexte50l(fmt.Sprintf(Orange+"La lettre : %s a déjà été proposé."+Reset, action))
				fmt.Println()
			}

		case mot_a_trouver:
			if Essais_restant >= 2 {
				defaite = false
			} else {
				fmt.Println()
				afficherTexte50(Yellow + "Il ne vous reste pas assez d'essais pour effectuer cette action." + Reset)
				continue
			}
		case "!save":

			fmt.Println()
			afficherTexte250(Cyan + "Sauvegarde en cours . . . . . . . . . . ." + Reset)
			fmt.Println()
			afficherTexte50(Cyan + "Sauvegarde terminée. Nous avons hâte de vous revoir!" + Reset)
			fmt.Println()

			player := Sauvegarde{
				Pseudo:         s,
				Mot:            mot_a_trouver,
				Lettres:        tableau_mot,
				Essais:         Essais_restant,
				Alphabets:      alphabet,
				Difficultesave: DifficulteReprendre,
			}

			liste_joueurs_json[numero_joueur] = player

			encodeData, encodeErr := json.Marshal(liste_joueurs_json)

			if encodeErr != nil {
				fmt.Printf("Erreur : %v", encodeErr)
				return
			}
			fichierErr := os.WriteFile("src/save.json", encodeData, 0644)
			if fichierErr != nil {
				fmt.Printf("Erreur : %v", fichierErr)
				return
			}
			return
		default:
			if (len(mot_a_trouver) == len(action)) && (Essais_restant >= 2) {
				valeur := 0
				for i := 0; i < len(action); i++ {
					if string(action[i]) >= "a" && string(action[i]) <= "z" {
						valeur++
					}
				}
				if valeur == len(action) {
					fmt.Println()
					afficherTexte50(Red + "Le mot n'est pas bon, vous perdez 2 essais." + Reset)
					fmt.Println()
					Essais_restant -= 2
				}

			} else if (len(mot_a_trouver) == len(action)) && (Essais_restant < 2) {
				fmt.Println()
				afficherTexte50(Yellow + "Il ne vous reste pas assez d'essais pour effectuer cette action." + Yellow)
				fmt.Println()
				continue
			} else {
				fmt.Println()
				afficherTexte50(Red + "Entrée invalide, veuillez proposer une seule lettre ou un mot de la bonne longueur." + Reset)
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
		PlaySoundLooseMario()
		afficherTexte50l(fmt.Sprintf(Red+"Défaite, vous n'avez plus d'essais restants! Le mot à trouver était : "+Reset+couleur+"%s\n"+Reset, mot_a_trouver))
		fmt.Println()
		fmt.Printf("%-40s%s", "", "  _______\n")
		fmt.Printf("%-40s%s", "", "  |/   |\n")
		fmt.Printf("%-40s%s", "", "  |    O\n")
		fmt.Printf("%-40s%s", "", "  |   /|\\\n")
		fmt.Printf("%-40s%s", "", "  |   / \\\n")
		fmt.Printf("%-40s%s", "", "__|__ \n")
		Status = "defaite"
		afficherTexteSoundFaux(". . .")
		stopSound()
	} else {
		Status = "victoire"
		fmt.Println()
		afficherTexte50l(fmt.Sprintf(Green+"Bravo! Le mot à trouver était bien : %s\n"+Reset, mot_a_trouver))
		fmt.Println()
	}

	alphabet_string := ""
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i].Lettre_trouve {
			alphabet_string = alphabet_string + alphabet[i].Lettre
		}
	}
	data2, Err2 := os.ReadFile("src/save.json")
	if Err2 != nil {
		if errors.Is(Err2, os.ErrNotExist) {
			data2 = []byte("[]")
		} else {
			fmt.Printf("Erreur : %v\n", Err2)
			os.Exit(0)
		}
	}

	var liste_joueurs_json2 []Sauvegarde
	Err2_json := json.Unmarshal(data2, &liste_joueurs_json2)
	if len(data2) > 0 {
		if Err2_json != nil {
			fmt.Printf("Erreur : %v", Err2_json)
			fmt.Println()
			os.Exit(0)
		}
	}

	var n_joueur int
	joueur_trouve := false

	for i := 0; i < len(liste_joueurs_json2); i++ {
		if nom == liste_joueurs_json2[i].Pseudo {
			n_joueur = i
			joueur_trouve = true
			break
		}
	}
	if joueur_trouve {
		liste_joueurs_json2 = append(liste_joueurs_json2[:n_joueur], liste_joueurs_json2[n_joueur+1:]...)
		if len(liste_joueurs_json2) == 0 {
			liste_joueurs_json2 = []Sauvegarde{}
		}
		data2, Erreur2 := json.Marshal(liste_joueurs_json2)
		if Erreur2 != nil {
			fmt.Printf("Erreur %v : ", Erreur2)
		}
		os.WriteFile("src/save.json", data2, 0644)
	}

	data_csv, Err2_CSV := os.OpenFile("src/score.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if Err2_CSV != nil {
		fmt.Printf("Erreur : %v", Err2_CSV)
		fmt.Println()
		os.Exit(0)
	}
	defer data_csv.Close()

	texte := csv.NewWriter(data_csv)
	defer texte.Flush()

	date := time.Now().Format("2006-01-02 15:04")
	Essais_affiche := strconv.Itoa(Essais_restant)
	nouvelle_data := []string{nom, Status, Essais_affiche, DifficulteReprendre, mot_a_trouver, alphabet_string, date}

	erreur := texte.Write(nouvelle_data)
	if erreur != nil {
		fmt.Printf("Erreur : %v\n", erreur)
	}
}
