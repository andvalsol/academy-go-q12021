package main

type IRepository interface {
	GetPokemonByID(id string) (Pokemon, error)
}
