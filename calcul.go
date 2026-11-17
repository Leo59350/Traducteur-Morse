package main

import (
	"errors"
	"strings"
)

var alphabetMorse = map[rune]string{
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..",
	'E': ".", 'F': "..-.", 'G': "--.", 'H': "....",
	'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
	'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",
}

func texteToMorse(texte string) (string, error) {
	if strings.TrimSpace(texte) == "" {
		return "", errors.New("ERROR")
	}

	// Vérifier les caractères avant de convertir en majuscules.
	for _, lettre := range texte {
		estMajuscule := lettre >= 'A' && lettre <= 'Z'
		estMinuscule := lettre >= 'a' && lettre <= 'z'

		if !estMajuscule && !estMinuscule && lettre != ' ' {
			return "", errors.New("ERROR")
		}
	}

	texte = strings.ToUpper(texte)
	mots := strings.Fields(texte)
	resultat := make([]string, 0, len(mots))

	for _, mot := range mots {
		lettres := make([]string, 0, len(mot))

		for _, lettre := range mot {
			lettres = append(lettres, alphabetMorse[lettre])
		}

		resultat = append(resultat, strings.Join(lettres, " "))
	}

	return strings.Join(resultat, " | "), nil
}

func morseToTexte(morse string) (string, error) {
	if strings.TrimSpace(morse) == "" {
		return "", errors.New("ERROR")
	}

	for _, caractere := range morse {
		if caractere != '.' &&
			caractere != '-' &&
			caractere != ' ' &&
			caractere != '|' {
			return "", errors.New("ERROR")
		}
	}

	alphabetTexte := make(map[string]rune, len(alphabetMorse))

	for lettre, code := range alphabetMorse {
		alphabetTexte[code] = lettre
	}

	mots := strings.Split(morse, "|")
	resultat := make([]string, 0, len(mots))

	for _, mot := range mots {
		codes := strings.Fields(mot)

		if len(codes) == 0 {
			return "", errors.New("ERROR")
		}

		var texte strings.Builder

		for _, code := range codes {
			lettre, existe := alphabetTexte[code]
			if !existe {
				return "", errors.New("ERROR")
			}

			texte.WriteRune(lettre)
		}

		resultat = append(resultat, texte.String())
	}

	return strings.Join(resultat, " "), nil
}
