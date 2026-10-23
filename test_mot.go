package main

import (
	"fmt"
	"os"
	"strings"
)

func ecrire() {
	texte := "mot.txt"
	contenu_byte, err := os.ReadFile(texte)
	if err != nil {
		fmt.Println("", err)
		return
	}
	contenu_string := string(contenu_byte)
	var tableau []string
	mot := ""

	for _, s := range contenu_string {
		if s != '\n' && s != '\t' && s != ' ' {
			mot = mot + string(s)
		} else {
			if mot != "" {
				tableau = append(tableau, mot)
				mot = ""
			}
		}

		if mot != "" {
			tableau = append(tableau, mot)
		}

		for i := 0; i < len(tableau); i++ {
			strings.ToLower(tableau[i])
			for j := 0; j < i; j++ {
				if !(j >= '0' && j <= '9') || (j >= 'à' && j <= 'ÿ') {
					tableau = append(tableau, tableau[:i]...)
				}
			}
		}

		nouveau_texte := ""
		for i, m := range tableau {
			if i > 0 {
				nouveau_texte += ""
			}
			nouveau_texte += m
		}

		nouveau := strings.Join(tableau, "\n")
		err = os.WriteFile("mot_correct.txt", []byte(nouveau), 0644)
		if err != nil {
			fmt.Println("Erreur écriture :", err)
			return
		}

		fmt.Println("ok")

	}
}
