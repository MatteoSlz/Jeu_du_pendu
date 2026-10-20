package main

import "fmt"

func (p *Sauvegarde) Start() {
	var Pseudo string
	fmt.Println("Comment vous appelez vous ?")
	fmt.Scan(&Pseudo)
	fmt.Printf("Vous vous appelez %s", Pseudo)
}
