package main
 
import (
<<<<<<< HEAD
    "fmt"
    "os"
    "strings"
)
 
func ecrire() {
=======
	"fmt"
	"os"
	"strings"
)

func readFileWords() []byte {
>>>>>>> 4bd73e9ea8f9d2999d01df66a605501535217bd0
	texte := "src/mot.txt"
	contenu_byte, err := os.ReadFile(texte)
	if err != nil {
		fmt.Printf("Erreur : %v", err)
		fmt.Println()
		os.Exit(0)
	}
	return contenu_byte
}

<<<<<<< HEAD
    

	for _, s := range contenu_string {
		if s != '\n' && s != '\t' && s != ' ' {
			mot = mot + string(s)
=======
func splitWordsFile(contenu_byte []byte) []string {
	var sentence []string
	var mot []byte

	for _, a := range contenu_byte {
		if a == '\n' {
			sentence = append(sentence, string(mot))
			mot = nil
>>>>>>> 4bd73e9ea8f9d2999d01df66a605501535217bd0
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

<<<<<<< HEAD
		for i := 0; i < len(tableau); i++ {
			for j := 0; j < i; j++ {
				if !('a' >= j && j <= 'z') {
					tableau = append(tableau, tableau[:i]...)
				}
			}
		}

		nouveau_texte := ""
		for i, m := range tableau {
			if !strings.Contains(m, "0"){
				mot = append(mot, mot[:i]...)
			}
			nouveau_texte += m
		}

		nouveau_contenu := []byte(nouveau_texte)
		err = os.WriteFile("src/mot_correct.txt", nouveau_contenu, 0644)
		if err != nil {
			fmt.Printf("Erreur : %v", err)
			fmt.Println()
			os.Exit(0)
		}
		fmt.Println("ok")
	}
}
 
=======
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
>>>>>>> 4bd73e9ea8f9d2999d01df66a605501535217bd0
