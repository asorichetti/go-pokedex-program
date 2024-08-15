package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, this will be a GoLang Pokedex Application to serve up pokemon based on user input, this statement is temporary")
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/hoenn/")
	fmt.Println("Please enter a pokedex number of any pokemon you want information on (Range from 1 to 202)")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var user_num int
	fmt.Scan(&user_num)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	fmt.Println(responseObject.Pokemon[user_num-1].Species.Name)
	fmt.Println(responseObject.Pokemon[user_num-1].Ptype)

}

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
	Ptype   string    `json:"type_name"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
	Ptype   string         `json:"type_name"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}
