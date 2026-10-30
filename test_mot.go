package main
 
import (
	"fmt"
	"os"
	"strings"
)

func readFileWords() []byte {
	texte := "src/mot.txt"
	contenu_byte, err := os.ReadFile(texte)
	if err != nil {
		fmt.Printf("Erreur : %v", err)
		fmt.Println()
		os.Exit(0)
	}
	return contenu_byte
}

func splitWordsFile(contenu_byte []byte) []string {
	var sentence []string
	var mot []byte

	for _, a := range contenu_byte {
		if a == '\n' {
			sentence = append(sentence, string(mot))
			mot = nil
		} else {
			mot = append(mot, a)
		}
	}
	if len(mot) > 0 {
		sentence = append(sentence, string(mot))
	}

	return sentence
}

func checkListWords(sentence []string) []string {
	listValideWords := []string{}
	for i := 0; i < len(sentence); i++ {
		isValide := true
		word := strings.TrimSpace(strings.ToLower(sentence[i]))
		for _, char := range word {
			if (char >= '0' && char <= '9') || (char >= 'à' && char <= 'ÿ') || char == ' ' {
				isValide = false
				break
			}
		}
		if isValide && word != "" {
			listValideWords = append(listValideWords, word)
		}
	}
	return listValideWords
}

func Sortie(texte string) {
	err := os.WriteFile("src/mot_correct.txt", []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture")
	}
}

func Recuperation() {
	fileContent := readFileWords()                           // string
	fileContentListwords := splitWordsFile(fileContent)      // []string
	listValidesWords := checkListWords(fileContentListwords) // []string
	Sortie(strings.Join(listValidesWords, "\n"))

}
