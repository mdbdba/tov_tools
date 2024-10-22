// cmd/get_table/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"tov_tools/pkg/character"
)

func main() {
	// Define the -type command-line flag
	dataType := flag.String("type", "", "The type of data to retrieve (e.g., 'class')")
	flag.Parse()

	if *dataType == "" {
		fmt.Println("Error: -type argument is required")
		os.Exit(1)
	}

	switch *dataType {
	case "class":
		printClassTable()
	default:
		fmt.Printf("Error: unsupported type '%s'\n", *dataType)
		os.Exit(1)
	}
}

// printClassTable prints out the class data in a text table
func printClassTable() {
	fmt.Println(character.ToStringTable())
}
