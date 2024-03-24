package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CP-Payne/pokedexcli/internal/pokeapi"
	"github.com/CP-Payne/pokedexcli/internal/pokecache"
	"github.com/chzyer/readline"
)

type Config struct {
	NextUrl  string
	PrevUrl  string
	Cache    *pokecache.Cache
	Location *pokeapi.LocationPokemonApiResponse
	Pokedex  *pokeapi.Pokedex
}

var commands map[string]cliCommand

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	c := &Config{
		NextUrl:  "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PrevUrl:  "",
		Cache:    cache,
		Location: &pokeapi.LocationPokemonApiResponse{},
		Pokedex: &pokeapi.Pokedex{
			Pokedex: make(map[string]pokeapi.PokemonDetails),
		},
	}

	// Using chzyer/readline package for better commandline experience

	rl, err := readline.New("Pokedex > ")
	if err != nil {
		log.Fatalf("readline.New failed: %v", err)
	}

	defer rl.Close()
	commands = c.getCommands()

	for {

		// Print a newline at the start to ensure the prompt starts with a new line
		fmt.Println()
		line, err := rl.Readline()
		if err != nil { // Ctrl-D now also returns EOF
			break
		}

		// Splitting the input line into command and parameters
		inputSlice := strings.Split(line, " ")
		commandInput := inputSlice[0]
		params := inputSlice[1:]

		// Looking up the command and executing it
		command, ok := commands[commandInput]
		if !ok {
			fmt.Println("Unknown command. See 'help' for valid commands.")
			continue
		}

		err = command.callback(params...)
		if err != nil {
			// fmt.Printf("Error executing %s command: %s\n", commandInput, err)
			fmt.Printf("%s\n", err)
		}
	}

	// THE BELOW USES SCANNER TO READ INPUT.
	// THIS DOES NOT PROVIDE A GOOD USER EXPERIENCE.
	//
	// scanner := bufio.NewScanner(os.Stdin)
	// commands = c.getCommands()
	//
	// for {
	// 	fmt.Print("\nPokedex > ")
	// 	scanner.Scan()
	//
	// 	err := scanner.Err()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	input := scanner.Text()
	// 	inputSlice := strings.Split(input, " ")
	// 	commandInput := inputSlice[0]
	// 	params := inputSlice[1:]
	//
	// 	command, ok := commands[commandInput]
	// 	if !ok {
	// 		fmt.Println("Unknown command. See 'help' for valid commands.")
	// 		continue
	// 	}
	//
	// 	err = command.callback(params...)
	// 	if err != nil {
	// 		fmt.Printf("Error executing %s command: %s", commandInput, err)
	// 	}
	//
	// }
}
