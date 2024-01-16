# Verano

Package pour gérer les activités dans un projet.
Ce projet est inspiré  du logiciel 'Primavera' ou P6 (Printemps en Espagnol), d'où le nom 'Verano' (Été en Espagnol).

# Motivation

- Comprendre comment Primavera fonctionne sous le capot.
- Ajouter des activités une par une dans Primavera prend beaucoup de temps.
- Présenter une alternative gratuite et open source a Primavera.
- Améliorer mes compétences en Go.

# Installation

```sh
# dans un module Go, excutez la commande suivante:
$ go get github.com/vanillaiice/verano
```

# Fonctionnalités

- Classer les activités dans un projet en fonction de leurs relations (seules les relations début à fin sont actuellement supportés).
- Calculer les dates de début et de fin pour chaque activité.
- Générer un graph (avec graphviz) montrant les activités et leurs relations.
- Analyser et traiter des listes d'activités au format JSON, CSV et XLSX.
- Stockage des activités dans une base de données SQLite.

> Veuillez consulter le dossier 'examples' dans ce repertoire pour voir ces fonctionnalités en action.

# Structure d'une activité

Voici la structure des activités utilisée par ce package:

```go
type Activity struct {
	Id             int           // Identifiant unique de l'activité
	Description    string        // Description de l'activité
	Duration       time.Duration // Durée de l'activité
	Start          time.Time     // Date de début de l'activité
	Finish         time.Time     // Date de fin de l'activité
	PredecessorsId []int         // ID des activités qui précèdent
	SuccessorsId   []int         // ID des activités qui suivent
	Cost           float64       // Coût de l'activité
}
```

# Auteur

Vanillaiice

# Licence

GPLv3
