package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lecteur := bufio.NewScanner(os.Stdin)

	fmt.Println("=== Traducteur Morse ===")
	fmt.Println("1 : Texte vers Morse")
	fmt.Println("2 : Morse vers texte")
	fmt.Print("Votre choix : ")

	if !lecteur.Scan() {
		fmt.Println("ERROR")
		return
	}

	choix := strings.TrimSpace(lecteur.Text())

	if choix != "1" && choix != "2" {
		fmt.Println("ERROR")
		return
	}

	if choix == "1" {
		fmt.Print("Entrez votre texte : ")
	} else {
		fmt.Print("Entrez votre Morse (lettres séparées par un espace, mots par |) : ")
	}

	if !lecteur.Scan() {
		fmt.Println("ERROR")
		return
	}

	entree := lecteur.Text()

	var resultat string
	var err error

	if choix == "1" {
		resultat, err = texteToMorse(entree)
	} else {
		resultat, err = morseToTexte(entree)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Résultat :", resultat)
}
