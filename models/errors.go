package models

import "errors"

var (
	ErrNotFound    = errors.New("pokémon introuvable")
	ErrBadRequest  = errors.New("requête invalide")
	ErrTimeout     = errors.New("l'API met trop de temps à répondre")
	ErrUnavailable = errors.New("l'API est indisponible")
)
