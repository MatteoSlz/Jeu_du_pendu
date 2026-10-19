package main

import (
	"os"
)

func Entree() ([]byte, error) {
	contenu, err := os.ReadFile("./mot.txt")
	if err != nil {
		return nil, err
	}
	return contenu, nil
}

func decoupage(contenu_byte []byte) []string {
	var mots []string // la liste des mots terminés
	var mot []byte    // le mot en cours de construction

	for _, r := range contenu_byte {
		if r == '\n' {
			mots = append(mots, string(mot) )
			mot = nil
		} else {
			mot = append(mot, r)
		}
	}
		if len(mot) > 0 {
		mots = append(mots, string(mot))
		}

	return mots
}

func Verification(contenu []string) {
	decoupage()

}