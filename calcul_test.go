package main

import "testing"

func TestTraduction(t *testing.T) {
	tests := []struct {
		nom            string
		sens           string
		entree         string
		attendu        string
		erreurAttendue string
	}{
		{
			nom:     "normal texte vers Morse",
			sens:    "texte",
			entree:  "SOS",
			attendu: "... --- ...",
		},
		{
			nom:     "normal Morse vers texte",
			sens:    "morse",
			entree:  "... --- ...",
			attendu: "SOS",
		},
		{
			nom:     "normal deux mots vers Morse",
			sens:    "texte",
			entree:  "A B",
			attendu: ".- | -...",
		},
		{
			nom:     "normal deux mots vers texte",
			sens:    "morse",
			entree:  ".- | -...",
			attendu: "A B",
		},
		{
			nom:     "limite une lettre",
			sens:    "texte",
			entree:  "E",
			attendu: ".",
		},
		{
			nom:     "limite un symbole Morse",
			sens:    "morse",
			entree:  "-",
			attendu: "T",
		},
		{
			nom:            "erreur texte vide",
			sens:           "texte",
			entree:         "",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur uniquement des espaces",
			sens:           "texte",
			entree:         "   ",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur Morse vide",
			sens:           "morse",
			entree:         "",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur chiffre interdit",
			sens:           "texte",
			entree:         "ABC1",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur accent interdit",
			sens:           "texte",
			entree:         "école",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur ponctuation interdite",
			sens:           "texte",
			entree:         "Salut!",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur code Morse inconnu",
			sens:           "morse",
			entree:         "......",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "erreur caractère interdit en Morse",
			sens:           "morse",
			entree:         ".- / -...",
			erreurAttendue: "ERROR",
		},
		{
			nom:     "piège minuscules",
			sens:    "texte",
			entree:  "sos",
			attendu: "... --- ...",
		},
		{
			nom:     "piège mélange majuscules minuscules",
			sens:    "texte",
			entree:  "Sos",
			attendu: "... --- ...",
		},
		{
			nom:     "piège espaces multiples",
			sens:    "texte",
			entree:  "  A   B  ",
			attendu: ".- | -...",
		},
		{
			nom:     "piège barre sans espaces",
			sens:    "morse",
			entree:  ".-|-...",
			attendu: "A B",
		},
		{
			nom:            "piège deux barres",
			sens:           "morse",
			entree:         ".- || -...",
			erreurAttendue: "ERROR",
		},
		{
			nom:            "piège lettres Morse sans espaces",
			sens:           "morse",
			entree:         "...---...",
			erreurAttendue: "ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.nom, func(t *testing.T) {
			var obtenu string
			var err error

			switch tt.sens {
			case "texte":
				obtenu, err = texteToMorse(tt.entree)
			case "morse":
				obtenu, err = morseToTexte(tt.entree)
			default:
				t.Fatalf("sens de test inconnu : %q", tt.sens)
			}

			if tt.erreurAttendue != "" {
				if err == nil {
					t.Fatal("une erreur était attendue")
				}

				if err.Error() != tt.erreurAttendue {
					t.Errorf(
						"erreur obtenue %q, attendue %q",
						err.Error(),
						tt.erreurAttendue,
					)
				}

				return
			}

			if err != nil {
				t.Fatalf("erreur inattendue : %v", err)
			}

			if obtenu != tt.attendu {
				t.Errorf(
					"résultat obtenu %q, attendu %q",
					obtenu,
					tt.attendu,
				)
			}
		})
	}
}
