// cmd/create_character/main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"tov_tools/pkg/character"
)

func main() {
	// Define the -class command-line flag
	className := flag.String("class", "", "The class of the character to create")
	flag.Parse()

	var selectedClass character.Class
	var err error

	if *className == "" {
		selectedClass = character.RandomClass()
	} else {
		selectedClass, err = character.GetClassByName(*className)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Printf("Created character with class:\n%s", selectedClass.ToString())
}
