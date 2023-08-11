package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		fmt.Println("You wrote", text)
	}
}

func cleanInput(str string) []string {
	lower := strings.ToLower(str)
	words := strings.Fields(lower)
	return words
}
