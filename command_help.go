package main

import "fmt"

func commandHelp() {
	fmt.Println("")
	fmt.Printf("Available commands: \n")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf(" >> %s : %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")

}
