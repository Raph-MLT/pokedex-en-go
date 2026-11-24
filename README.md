# pokedex-en-go

Site web en Go qui affiche un Pokédex en français avec les données de l'API [Tyradex] (https://tyradex.app). Il utilise seulement la bibliothèque standard de Go.

Le site affiche la liste des Pokémon et la fiche de chacun, et permet de chercher un Pokémon par son nom. En cas de problème (numéro invalide, Pokémon inexistant, page inconnue, API indisponible), il affiche une page d'erreur avec le code HTTP qui correspond.

## Lancement

Avec Go installé, lancer depuis la racine du projet :

```
go run .
```

puis ouvrir http://localhost:8080. Lancé depuis un autre dossier, le programme ne trouve pas `templates/` et s'arrête.

## Architecture

```
pokedex/
├── main.go         lancement du serveur
├── api/            appels HTTP vers Tyradex
├── models/         structures de données et erreurs connues
├── services/       vérification des paramètres, recherche
├── controllers/    lecture de la requête, choix de la page
├── routes/         routeur
├── helpers/        chargement et rendu des templates
├── templates/      list.html, pokemon.html, error.html
└── assets/css/     style.css
```
