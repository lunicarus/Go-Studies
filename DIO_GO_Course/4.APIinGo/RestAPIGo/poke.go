// https://pokeapi.co/api/v2/pokedex/kanto/
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	println(string(responseData))
}
