package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

type mots struct {
	Lettre        string `json:"lettre"`
	Lettre_trouve bool   `json:"lettre_trouve"`
}

type Sauvegarde struct {
	Pseudo    string `json:"pseudo"`
	Mot       string `json:"mot"`
	Lettres   []mots `json:"lettres"`
	Essais    int    `json:"essais"`
	Alphabets []mots `json:"alphabet"`
}

var liste_joueurs = []Sauvegarde{}

var alphabet = []mots{
	{Lettre: "a", Lettre_trouve: false},
	{Lettre: "b", Lettre_trouve: false},
	{Lettre: "c", Lettre_trouve: false},
	{Lettre: "d", Lettre_trouve: false},
	{Lettre: "e", Lettre_trouve: false},
	{Lettre: "f", Lettre_trouve: false},
	{Lettre: "g", Lettre_trouve: false},
	{Lettre: "h", Lettre_trouve: false},
	{Lettre: "i", Lettre_trouve: false},
	{Lettre: "j", Lettre_trouve: false},
	{Lettre: "k", Lettre_trouve: false},
	{Lettre: "l", Lettre_trouve: false},
	{Lettre: "m", Lettre_trouve: false},
	{Lettre: "n", Lettre_trouve: false},
	{Lettre: "o", Lettre_trouve: false},
	{Lettre: "p", Lettre_trouve: false},
	{Lettre: "q", Lettre_trouve: false},
	{Lettre: "r", Lettre_trouve: false},
	{Lettre: "s", Lettre_trouve: false},
	{Lettre: "t", Lettre_trouve: false},
	{Lettre: "u", Lettre_trouve: false},
	{Lettre: "v", Lettre_trouve: false},
	{Lettre: "w", Lettre_trouve: false},
	{Lettre: "x", Lettre_trouve: false},
	{Lettre: "y", Lettre_trouve: false},
	{Lettre: "z", Lettre_trouve: false},
}

var nom string
var Status string
var Essais_restant int

func afficherTexte50l(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(50 * time.Millisecond)
	}
}
func afficherTexte50(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}
func afficherTexte250(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(250 * time.Millisecond)
	}
	fmt.Println()
}

func Start() string {

	data, Err_lecture_json := os.ReadFile("save.json")
	if Err_lecture_json != nil {
		if errors.Is(Err_lecture_json, os.ErrNotExist) {
			fmt.Println("Erreur")
			os.WriteFile("save.json", []byte("[]"), 0644)
			data = []byte("[]")
		} else {
			fmt.Printf("Erreur : %v\n", Err_lecture_json)
			os.Exit(0)
		}
	}
	var liste_joueurs_json []Sauvegarde
	Err_decode := json.Unmarshal(data, &liste_joueurs_json)
	if len(data) > 0 {
		if Err_decode != nil {
			fmt.Printf("Erreur : %v", Err_decode)
			fmt.Println()
			os.Exit(0)
		}
	}

	pseudo_choisis := true
	var ecrase string
	var n_joueur int
	saisie_valide := true

	for pseudo_choisis {
		fmt.Println()
		afficherTexte50l("Comment vous appelez vous : ")
		fmt.Scan(&nom)
		for i := 0; i < len(liste_joueurs_json); i++ {
			if nom == liste_joueurs_json[i].Pseudo {
				for saisie_valide {
					n_joueur = i
					afficherTexte50(fmt.Sprintf("Le pseudo %s existe déjà ", nom))
					afficherTexte50("Voulez vous écraser la sauvegarde précédente effectué sur ce pseudo ?")
					afficherTexte50("oui/non : ")
					fmt.Scan(&ecrase)
					ecrase = strings.ToLower(ecrase)
					switch ecrase {
					case "oui":
						liste_joueurs_json = append(liste_joueurs_json[:n_joueur], liste_joueurs_json[n_joueur+1:]...)
						encodeData, encodeErr := json.Marshal(liste_joueurs_json)

						if encodeErr != nil {
							fmt.Printf("Erreur encodage données : %v", encodeErr)
						}
						fileErr := os.WriteFile("save.json", encodeData, 0644)
						if fileErr != nil {
							fmt.Printf("Erreur encodage données : %v", fileErr)
						}
						pseudo_choisis = false
						saisie_valide = false
					case "non":
						fmt.Println()
						afficherTexte50("Pas de problème. Au revoir !")
						fmt.Println()
						os.Exit(0)
					default:
						fmt.Println()
						afficherTexte50("Veuillez entrer une saisie valide.")
						fmt.Println()
					}
				}
			}
		}
		fmt.Println()
		afficherTexte50(fmt.Sprintf("D'accord bienvenue %s !", nom))
		pseudo_choisis = false
	}
	return nom
}

func jouer() {
	s := Start()
	perdu := false
	Essais_restant := 6
	defaite := true
	count_bonne_lettre := 0
	liste_mot := "mot_correct.txt"
	contenu_byte, err := os.ReadFile(liste_mot)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	tableau_final := strings.Fields(string(contenu_byte))
	if len(tableau_final) == 0 {
		fmt.Println("Erreur : Le fichier est vide ou ne contient aucun mot valide.")
		return
	}
	index_mots_random := rand.IntN(len(tableau_final))
	mot_a_trouver := strings.ToLower(tableau_final[index_mots_random])

	tableau_mot := make([]mots, len(mot_a_trouver))
	for i, s := range mot_a_trouver {
		tableau_mot[i].Lettre = string(s)
		tableau_mot[i].Lettre_trouve = false
	}
	fmt.Println()
	fmt.Println()
	afficherTexte50("Bienvenue au jeu du pendu !")
	afficherTexte50(fmt.Sprintf("Votre mot à trouver contient %d lettres. ", len(mot_a_trouver)))
	fmt.Println("À toi de jouer !")
	index_lettre_donnee_debut := rand.IntN(len(mot_a_trouver))
	lettre_donnee_debut := tableau_mot[index_lettre_donnee_debut].Lettre
	for i := 0; i < len(tableau_mot); i++ {
		if lettre_donnee_debut == tableau_mot[i].Lettre {
			tableau_mot[i].Lettre_trouve = true
		}
	}
	for j := 0; j < len(alphabet); j++ {
		if lettre_donnee_debut == alphabet[j].Lettre {
			alphabet[j].Lettre_trouve = true
		}
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
				afficherTexte50l(fmt.Sprintf("%s ", tableau_mot[i].Lettre))
			} else if tableau_mot[i].Lettre == "-" {
				afficherTexte50l("- ")
				tableau_mot[i].Lettre_trouve = true

			} else {
				afficherTexte50l("_ ")
			}
		}

		afficherTexte50l(fmt.Sprintf("\tIl vous reste %d essai(s)", Essais_restant))
		fmt.Println()
		afficherTexte50l("Lettres proposées : ")
		for i := 0; i < len(alphabet); i++ {
			if alphabet[i].Lettre_trouve == true {
				afficherTexte50l(fmt.Sprintf("%s ", alphabet[i].Lettre))
			}
		}
		fmt.Println()
		fmt.Println()
		afficherTexte50("Que voulez-vous proposer : ")
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
					for j := 0; j < len(alphabet); j++ {
						if action == alphabet[j].Lettre {
							alphabet[j].Lettre_trouve = true
						}
					}
				}
				if lettre_trouvee {
					afficherTexte50l(fmt.Sprintf("Bravo ! Vous avez trouvé la lettre : %s.\n", action))
				} else {
					fmt.Println()
					afficherTexte50l(fmt.Sprintf("Dommage la lettre '%s' n'est pas dans le mot.", action))
					fmt.Println()
					Essais_restant--
				}
			} else {
				fmt.Println()
				afficherTexte50l(fmt.Sprintf("La lettre : %s a déjà été proposé.", action))
				fmt.Println()
			}

		case mot_a_trouver:
			if Essais_restant >= 2 {
				defaite = false
			} else {
				fmt.Println()
				afficherTexte50("Il ne vous reste pas assez d'essais pour effectuer cette action.")
				continue
			}
		case "!save":
			fmt.Println()
			afficherTexte250("Sauvegarde en cours . . . . . . . . . . .")
			fmt.Println()
			afficherTexte50("Sauvegarde terminée. Nous avons hâte de vous revoir!")
			fmt.Println()

			player := Sauvegarde{
				Pseudo:    s,
				Mot:       mot_a_trouver,
				Lettres:   tableau_mot,
				Essais:    Essais_restant,
				Alphabets: alphabet,
			}

			data, Err2 := os.ReadFile("save.json")
			if Err2 != nil {
				if errors.Is(Err2, os.ErrNotExist) {
					data = []byte("[]")
				} else {
					fmt.Printf("Erreur : %v\n", Err2)
					os.Exit(0)
				}
			}
			var liste_joueurs_json []Sauvegarde
			Err2_json := json.Unmarshal(data, &liste_joueurs_json)
			if len(data) > 0 {
				if Err2_json != nil {
					fmt.Printf("Erreur : %v", Err2_json)
					fmt.Println()
					os.Exit(0)
				}
			}

			liste_joueurs = append(liste_joueurs_json, player)

			encodeData, encodeErr := json.Marshal(liste_joueurs)

			if encodeErr != nil {
				fmt.Printf("Erreur : %v", encodeErr)
				return
			}
			fichierErr := os.WriteFile("save.json", encodeData, 0644)
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
					afficherTexte50("Le mot n'est pas bon, vous perdez 2 essais.")
					fmt.Println()
					Essais_restant -= 2
				}

			} else if (len(mot_a_trouver) == len(action)) && (Essais_restant < 2) {
				fmt.Println()
				afficherTexte50("Il ne vous reste pas assez d'essais pour effectuer cette action.")
				fmt.Println()
				continue
			} else {
				fmt.Println()
				afficherTexte50("Entrée invalide, veuillez proposer une seule lettre ou un mot de la bonne longueur.")
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
		afficherTexte50l(fmt.Sprintf("Dommage c'est perdu. Le mot à trouver était : %s\n", mot_a_trouver))
		fmt.Println()
		Status = "defaite"
	} else {
		Status = "victoire"
		fmt.Println()
		afficherTexte50l(fmt.Sprintf("Bravo! Le mot à trouver était bien : %s\n", mot_a_trouver))
		fmt.Println()
	}

	alphabet_string := ""
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i].Lettre_trouve {
			alphabet_string = alphabet_string + alphabet[i].Lettre
		}
	}

	data_csv, Err2_CSV := os.OpenFile("score.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	nouvelle_data := []string{nom, Status, Essais_affiche, mot_a_trouver, alphabet_string, date}

	erreur := texte.Write(nouvelle_data)
	if erreur != nil {
		fmt.Printf("Erreur : %v\n", erreur)
	}
}
