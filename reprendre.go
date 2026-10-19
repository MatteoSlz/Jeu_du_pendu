package main

import (
	"fmt"
)

type  Sauvegarde struct {
	Pseudo string  `json:"pseudo"`
	Mot string 		`json:"mot"`
	Lettres []string `json:"lettres"`
	Essais int `json:"essais"`
}

func reprendre() {
	fmt.Println("reprendre")
}