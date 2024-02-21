package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaquinleonarg/go-pokemon/src/pkg/pokemon"
)

var pokemonService pokemon.PokemonService

func StartServer(port int, serviceType string) {
	// TODO: Decide kind of service, switch
	pokemonService = pokemon.NewPokemonLocalService()

	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", GetPokemonHandler).Methods(http.MethodGet)
	r.HandleFunc("/pokemon", AddPokemonHandler).Methods(http.MethodPost)

	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

func GetPokemonHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name, ok := pathParams["name"]
	if !ok {
		http.Error(w, "missing name param", http.StatusBadRequest)
		return
	}

	pokemon, err := pokemonService.GetPokemonByName(name)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(pokemon); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AddPokemonHandler(w http.ResponseWriter, r *http.Request) {
	var pokemon pokemon.Pokemon
	err := json.NewDecoder(r.Body).Decode(&pokemon)
	if err != nil {
		http.Error(w, "bad data", http.StatusBadRequest)
		return
	}

	err = pokemonService.AddPokemon(pokemon)
	if err != nil {
		http.Error(w, "already exists", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
