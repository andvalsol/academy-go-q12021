package main

type GetPokemonUseCase struct {
	Repository Repository
}

/// Function that is used for getting a [Pokemon] object by its ID
func (useCase *GetPokemonUseCase) GetPokemonById(id int) (Pokemon, error) {
	pokemon, error := useCase.Repository.GetPokemonByID(id)

	if error == nil {
		return pokemon, error // TODO why can't we return nil here?
	}

	return pokemon, error
}
