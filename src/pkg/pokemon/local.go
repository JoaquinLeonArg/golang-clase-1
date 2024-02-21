package pokemon

import "fmt"

type PokemonLocalService struct {
	Pokemons []Pokemon
}

type Pokemon struct {
	Name string      `json:"name"`
	Type PokemonType `json:"pokemon_type"` // parse type format to specific http request/response body
}

// go pseudo enum
type PokemonType string

const (
	PokemonTypeNormal PokemonType = "normal"
	PokemonTypeFire   PokemonType = "fire"
	PokemonTypeWater  PokemonType = "water"
	PokemonTypeGrass  PokemonType = "grass"
)

func NewPokemonLocalService() PokemonLocalService {
	return PokemonLocalService{
		Pokemons: []Pokemon{
			{Name: "Bulbasaur", Type: PokemonTypeGrass},
			{Name: "Charmander", Type: PokemonTypeFire},
			{Name: "Squirtle", Type: PokemonTypeWater},
		},
	}
}

func (ps *PokemonLocalService) GetPokemonByName(name string) (*Pokemon, error) {
	for _, pokemon := range ps.Pokemons {
		if pokemon.Name == name {
			return &pokemon, nil
		}
	}
	return nil, fmt.Errorf("Pokemon not found")
}

func (ps *PokemonLocalService) AddPokemon(poke Pokemon) error {
	if _, err := ps.GetPokemonByName(poke.Name); err == nil {
		return fmt.Errorf("Pokemon already exists")
	}
	ps.Pokemons = append(ps.Pokemons, poke)
	return nil
}
