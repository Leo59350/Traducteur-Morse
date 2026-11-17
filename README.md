# Traducteur Morse

## Présentation
Ce programme traduit du texte en Morse et du Morse en texte.

En Morse :
- Les lettres sont séparées par un espace.
- Les mots sont séparés par une barre verticale `|`.

## Utilisation
Depuis le dossier du projet :

```powershell
go run main.go calcul.go
```

Choisir :
- 1 : Texte vers Morse.
- 2 : Morse vers texte.

Saisir ensuite le contenu à traduire.

Si le module Go a été initialisé, on peut aussi utiliser :

```powershell
go run .
```

## Exemples

Texte :
sos sos

Morse :
... --- ... | ... --- ...

Texte :
A B

Morse :
.- | -...

Pour décoder SOS, saisir `... --- ...` avec des espaces
entre les lettres Morse.

## Choix et limites
- Les lettres A à Z sont acceptées en majuscules ou minuscules.
- Les minuscules sont converties en majuscules.
- Les accents, chiffres et signes de ponctuation provoquent ERROR.
- Une entrée vide provoque ERROR.
- Un code Morse inconnu provoque ERROR.
- Les espaces multiples du texte sont regroupés.
- Les tabulations sont refusées.
- Le résultat du décodage est en majuscules.

## Vérifications
Après initialisation du module :

```powershell
gofmt -w .
go test ./...
go vet ./...
```
