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

const (
	BleuCiel256 = "\033[38;5;117m"
	Orange      = "\033[38;5;208m"
	Reset       = "\033[0m"
	Black       = "\033[30m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	Brown       = "\033[38;5;94m"
	LightGray   = "\033[37m"
	Gray        = "\033[90m"
	Turquoise   = "\033[38;5;43m"
)

var couleur = BleuCiel256

type mots struct {
	Lettre        string `json:"lettre"`
	Lettre_trouve bool   `json:"lettre_trouve"`
}

type Sauvegarde struct {
	Pseudo         string `json:"pseudo"`
	Mot            string `json:"mot"`
	Lettres        []mots `json:"lettres"`
	Essais         int    `json:"essais"`
	Alphabets      []mots `json:"alphabet"`
	Difficultesave string `json:difficulte`
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
var oui = true
var nom string
var Status string
var Essais_restant int
var Difficulte string

func afficherTexte50l(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(5 * time.Millisecond)
	}
}

func afficherTexteSoundValide(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(1000 * time.Millisecond)
	}
}
func afficherTexteSoundFaux(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(600 * time.Millisecond)
	}
}

func afficherTexte50(texte string) {
	for _, caractere := range texte {
		fmt.Print(string(caractere))
		time.Sleep(5 * time.Millisecond)
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

	data, Err_lecture_json := os.ReadFile("src/save.json")
	if Err_lecture_json != nil {
		if errors.Is(Err_lecture_json, os.ErrNotExist) {
			fmt.Println("Erreur")
			os.WriteFile("src/save.json", []byte("[]"), 0644)
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
		afficherTexte50l(Gray + "Comment vous appelez vous : " + Reset)
		fmt.Scan(&nom)
		for i := 0; i < len(liste_joueurs_json); i++ {
			if nom == liste_joueurs_json[i].Pseudo {
				for saisie_valide {
					n_joueur = i
					afficherTexte50(fmt.Sprintf(Yellow+"Le pseudo %s existe déjà "+Reset, nom))
					afficherTexte50(Magenta + "Voulez vous écraser la sauvegarde précédente effectué sur ce pseudo ?" + Reset)
					afficherTexte50(Magenta + "oui/non : " + Reset)
					fmt.Scan(&ecrase)
					ecrase = strings.ToLower(ecrase)
					switch ecrase {
					case "oui":
						liste_joueurs_json = append(liste_joueurs_json[:n_joueur], liste_joueurs_json[n_joueur+1:]...)
						encodeData, encodeErr := json.Marshal(liste_joueurs_json)

						if encodeErr != nil {
							fmt.Printf("Erreur : %v", encodeErr)
						}
						fileErr := os.WriteFile("src/save.json", encodeData, 0644)
						if fileErr != nil {
							fmt.Printf("Erreur : %v", fileErr)
						}
						pseudo_choisis = false
						saisie_valide = false
					case "non":
						fmt.Println()
						afficherTexte50(Magenta + "Pas de problème. Au revoir !" + Reset)
						fmt.Println()
						os.Exit(0)
					default:
						fmt.Println()
						afficherTexte50(Red + "Veuillez entrer une saisie valide." + Reset)
						fmt.Println()
					}
				}
			}
		}
		fmt.Println()
		afficherTexte50(fmt.Sprintf(Gray+"D'accord bienvenue %s !"+Reset, nom))
		pseudo_choisis = false
	}
	return nom
}

func jouer() {
	s := Start()
	difficulte := ecriture_difficulte()
	difficulteFinal := difficulte
	perdu := false
	Essais_restant := 6
	if difficulte == "hardcore" {
		Essais_restant = 2
		difficulteFinal = "difficile"
	}
	defaite := true
	count_bonne_lettre := 0

	liste_mot := "src/mot_" + difficulteFinal + ".txt"
	contenu_byte, err := os.ReadFile(liste_mot)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	tableau_final := strings.Fields(string(contenu_byte))
	if len(tableau_final) == 0 {
		fmt.Println("Erreur : Le fichier est vide.")
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
	longueur_mot := len(mot_a_trouver)
	for i := 0; i < len(mot_a_trouver); i++ {
		if string(mot_a_trouver[i]) == "-" {
			longueur_mot--
		}
	}
	afficherTexte50(couleur + "Bonne chance !" + Reset)
	fmt.Println()
	afficherTexte50(fmt.Sprintf(couleur+"Votre mot à trouver contient"+Reset+Red+" %d "+Reset+couleur+"lettres. "+Reset, longueur_mot))
	fmt.Println()
	afficherTexte50(couleur + "À toi de jouer !" + Reset)
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
				afficherTexte50l(fmt.Sprintf("%s ", tableau_mot[i].Lettre))
			} else if tableau_mot[i].Lettre == "-" {
				afficherTexte50l(couleur + "- " + Reset)
				tableau_mot[i].Lettre_trouve = true

			} else {
				afficherTexte50l("_ ")
			}
		}

		afficherTexte50l(fmt.Sprintf(Yellow+"\tIl vous reste %d essai(s)\n"+Reset, Essais_restant))

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
					for j := 0; j < len(alphabet); j++ {
						if action == alphabet[j].Lettre {
							alphabet[j].Lettre_trouve = true
						}
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
				Difficultesave: Difficulte,
			}

			data, Err2 := os.ReadFile("src/save.json")
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
					afficherTexte50(Red + "Le mot n'est pas bon, vous perdez 2 essais." + Red)
					fmt.Println()
					Essais_restant -= 2
				}

			} else if (len(mot_a_trouver) == len(action)) && (Essais_restant < 2) {
				fmt.Println()
				afficherTexte50(Yellow + "Il ne vous reste pas assez d'essais pour effectuer cette action." + Reset)
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
		afficherTexteSoundValide(". . .")
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

	nouvelle_data := []string{nom, Status, Essais_affiche, difficulte, mot_a_trouver, alphabet_string, date}

	erreur := texte.Write(nouvelle_data)
	if erreur != nil {
		fmt.Printf("Erreur : %v\n", erreur)
	}
}
