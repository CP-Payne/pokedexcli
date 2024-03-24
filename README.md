# PokedexCLI

PokedexCLI is a command-line interface tool built with Go that allows users to interact with the [PokeAPI](https://pokeapi.co/docs/v2). Designed for Pokémon enthusiasts, it offers a variety of features including searching for areas on a map, viewing Pokémon in an area, catching Pokémon, and more. Leveraging the PokeAPI, this tool brings the world of Pokémon to your terminal, complete with caching for faster subsequent requests.

## Features

- **Map Navigation**: Easily navigate through different locations on the map and explore areas where Pokémon can be found.
- **Catch Pokémon**: Encounter and catch Pokémon to add them to your Pokedex.
- **Pokedex**: View detailed information about each Pokémon you've caught, including stats, abilities, and more.
- **Caching**: API requests are cached to speed up access to previously retrieved information.
- **Help System**: Unsure of what commands are available? Just type `help` to get detailed information about all available commands.

## Commands

- `help`: Displays a help message with information about all commands.
- `exit`: Exits the PokedexCLI application.
- `map`: Retrieves a list of locations on the map. Calling this repeatedly will fetch the next set of locations.
- `mapb`: Fetches the previous list of locations on the map.
- `explore`: Lists all Pokémon in a specific area.
- `catch`: Attempts to catch a Pokémon.
- `inspect`: Inspects detailed information about a caught Pokémon.
- `pokedex`: Displays all Pokémon that have been caught and are present in your Pokedex.

## Installation Instructions

To run PokedexCLI, you'll need Go installed on your machine. This project also makes use of an external package, which will be automatically handled by Go Modules. Follow these steps to get started:

1. **Clone the repository to your local machine.**

```bash
git clone https://github.com/yourusername/pokedexCLI.git
```

2. **Navigate to the cloned repository.**

```bash
cd pokedexCLI
```

3. **Build the project.**

```bash
go build
```

4. **Run the compiled binary.**

```bash
./pokedexcli
```

## External Packages

This project depends on the following external package(s):

- [github.com/chzyer/readline](https://github.com/chzyer/readline): A Go package that provides an interface for line editing and history capabilities. It is used in this project to enhance the command-line interface, allowing for features like navigable input history and editable command lines, which improve the user experience by making it easier to edit and repeat commands.

## Future Enhancements

- Simulate battles between pokemon.
- Add more unit tests
- Allow for pokemon that are caught to evolve after a set amount of time
- Persist a user's Pokedex to disk so they can save progress between sessions
- Introduce random encounters with wild Pokémon.
- Add support for various types of Pokéballs, each with different catch probabilities.
