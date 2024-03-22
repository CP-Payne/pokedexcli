package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CP-Payne/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type Config struct {
	NextUrl string
	PrevUrl string
	Cache   *pokecache.Cache
}

var commands map[string]cliCommand

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	c := &Config{
		NextUrl: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		PrevUrl: "",
		Cache:   cache,
	}

	scanner := bufio.NewScanner(os.Stdin)
	commands = c.getCommands()

	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()

		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		input := scanner.Text()
		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command. See 'help' for valid commands.")
			continue
		}

		err = command.callback()
		if err != nil {
			fmt.Printf("Error executing %s command: %s", input, err)
		}

		// fmt.Printf("Your input is: %s\n", scanner.Text())
	}
}
