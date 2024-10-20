// cmd/roll/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"tov_tools/pkg/dice"
)

func main() {
	sides := flag.Int("sides", 20, "number of sides on the dice")
	times := flag.Int("times", 1, "number of times to roll the dice")
	options := flag.String("options", "",
		"comma-separated options for the roll (e.g. 'advantage,disadvantage,drop lowest')")
	ctxRef := flag.String("ctx", "default", "context reference for the roll")

	flag.Parse()

	opts := []string{}
	if *options != "" {
		opts = strings.Split(*options, ",")
	}

	roll, err := dice.Perform(*sides, *times, *ctxRef, opts...)
	if err != nil {
		log.Fatalf("Error performing roll: %v", err)
	}

	fmt.Println(roll.ToPrettyString())
}
