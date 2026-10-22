package main

import (
	"fmt"
	"os"
)

func ecrire() {
	texte := "mot.txt"
	contenu_byte, err := os.ReadFile(texte)
	if err != nil {
		fmt.Printf("Erreur : %v", err)
		fmt.Println()
		os.Exit(0)
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
			for j := 0; j < i; j++ {
				if !('a' >= j && j <= 'z') {
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

		nouveau_contenu := []byte(nouveau_texte)
		err = os.WriteFile("mot_correct.txt", nouveau_contenu, 0644)
		if err != nil {
			fmt.Println("Erreur : %v", err)
			fmt.Println()
			os.Exit(0)
		}
		fmt.Println("ok")
	}
}
