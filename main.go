package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	startRepl()
}

/*
return map[string]cliCommand{
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
}
*/
