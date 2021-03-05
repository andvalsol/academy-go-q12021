package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Repository struct {
}

func (*Repository) GetPokemonByID(id int) (Pokemon, error) {
	var pokemon Pokemon

	// Open the CSV file
	csvFile, err := os.Open("pokemons.csv")

	// Early exit in case of an error
	if err != nil {
		return pokemon, err
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()

		if err != nil {
			// Check if we have reached the end of the CSV file
			if err == io.EOF {
				break
			} else {
				if _, ok := err.(*csv.ParseError); ok {
					// There was an error parsing the CSV file
					continue
				}
			}
		}

		// record[0] represents the id of the pokemon
		intID, errorID := strconv.Atoi(record[0])

		if errorID != nil {
			continue
		}

		if intID == id {
			pokemon, err = transformCSVRecordIntoPokemon(record)

			if err != nil {
				return pokemon, err
			}
		}
	}

	return pokemon, nil
}

func transformCSVRecordIntoPokemon(record []string) (Pokemon, error) {
	// From the record we receive i.e => [1, "pikachu", 10]
	// which record[0] is an integer that represents the id
	//		 record[1] is a string that represents the name
	//		 record[2] is an integer that represents the height

	pokemon := new(Pokemon)

	// Convert the ID
	id, errorID := strconv.Atoi(record[0])

	if errorID != nil {
		return *pokemon, errorID
	}

	// Get the name from the records
	name := record[1]

	// Convert the height into an integer
	height, errorHeight := strconv.Atoi(record[2])

	if errorHeight != nil {
		return *pokemon, errorHeight
	}

	pokemon.ID = id
	pokemon.Name = name
	pokemon.Height = height

	return *pokemon, nil
}
