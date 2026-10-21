package main 

import (
	"fmt"
	"encoding/csv"
	"os"
)

func score() {

data_csv, _ := os.OpenFile("score.csv", os.O_RDONLY, 0644)
defer data_csv.Close()

texte := csv.NewReader(data_csv)

lignes, erreur := texte.ReadAll()
if erreur != nil {
	fmt.Printf("csv %v", erreur)
}

fmt.Printf("PSEUDO\t STATUT\t\t ESSAIS\t MOT\t\t DATE\n")

for i := 0; i < len(lignes); i++ {
	nom := lignes[i][0]
	status := lignes[i][1]
	essais := lignes[i][2]
	mot := lignes[i][3]
	date := lignes[i][5]

	fmt.Printf("%s\t %s\t %s\t %s\t %s\n",nom,status,essais,mot,date)
}

}