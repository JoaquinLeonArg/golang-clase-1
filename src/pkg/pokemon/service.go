package pokemon

type PokemonService interface {
	GetPokemonByName(name string) (*Pokemon, error)
	AddPokemon(Pokemon) error
}
