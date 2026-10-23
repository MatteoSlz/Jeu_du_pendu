package main

import (
	"fmt"
	"os"
	"strings"
)

func ChoixDifficulte() string {
	difficulteChoix := true
	for difficulteChoix {
		afficherTexte50(Magenta + "Quelle difficulté voulez-vous choisir ? ("+ Reset + Green + "Facile" + Reset + Magenta + "/" + Reset + Yellow +"Normal" + Reset + Magenta + "/" + Reset + Orange + "Difficile"+ Reset + Magenta + "/" + Reset + Red +"Hardcore"+ Reset + Magenta + ")" + Reset)
		afficherTexte50l(Magenta +": " + Reset)
		fmt.Scan(&Difficulte)
		Difficulte = strings.ToLower(Difficulte)
		switch Difficulte {
		case "facile":
			return "facile"
		case "normal":
			return "normal"
		case "difficile":
			return "difficile"
		case "hardcore":
			return "hardcore"
		default:
			afficherTexte50(Yellow +"Veuillez choisir entre Facile / Normal / Difficile / Hardcore" + Yellow)
		}
	}
	return ""
}

func ecriture_difficulte() string {
	difficulte := ChoixDifficulte()
	difficulteFinal := difficulte

	if difficulte == "hardcore" {
		difficulteFinal = "difficile"
	}
	liste_mot_final := "src/mot_" + difficulteFinal + ".txt"
	contenu_byte, err := os.ReadFile("src/mot_correct.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return ""
	}

	tableau_final := strings.Fields(string(contenu_byte))

	if len(tableau_final) == 0 {
		fmt.Println("Erreur : Le fichier est vide ou ne contient aucun mot valide .")
		return ""
	}

	nouveau_texte := ""

	for i, m := range tableau_final {
		if ((len(tableau_final[i]) <= 7) && (difficulteFinal == "facile")) {
			if i > 0 {
				nouveau_texte += ""
			}
			nouveau_texte += m + "\n"
		}
		if ((len(tableau_final[i]) > 7) && (len(tableau_final[i]) <= 11) && (difficulteFinal == "normal")) {
			if i > 0 {
				nouveau_texte += ""
			}
			nouveau_texte += m + "\n"
		}
		if ((len(tableau_final[i]) > 11) && (difficulteFinal == "difficile")) {
			if i > 0 {
				nouveau_texte += ""
			}
			nouveau_texte += m + "\n"
		}
	}

	nouveau_contenu := []byte(nouveau_texte)
	err = os.WriteFile(liste_mot_final, nouveau_contenu, 0644)
	if err != nil {
		fmt.Printf("Erreur : %v", err)
		fmt.Println()
		os.Exit(0)
	}
	return difficulte
}
