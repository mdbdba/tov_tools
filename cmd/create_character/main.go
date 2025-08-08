package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"tov_tools/pkg/types"
)

func main() {
	characterName := flag.String("name", "", "The name of the character to create")
	characterLevel := flag.Int("level", 1, "The level of the character to create")
	className := flag.String("class", "", "The class of the character to create")
	subclassName := flag.String("subclass", "", "The subclass of the character to create")
	lineageName := flag.String("lineage", "", "The lineage of the character to create")
	heritageName := flag.String("heritage", "", "The heritage of the character to create")
	sizeName := flag.String("size", "", "The size of the character to create")
	abilityGenMethod := flag.String("ability-generation", "standard", "The ability generation method")
	selectedTraitsJSON := flag.String("traits", "", "The traits of the character to create (JSON format)")
	talentsJSON := flag.String("talents", "", "The talents of the character (JSON array format)")
	languagesJSON := flag.String("languages", "", "The languages of the character (JSON array format)")
	apiBaseURL := flag.String("api-url", "http://localhost:8080", "Base URL for the API")

	flag.Parse()

	// Validate required fields
	if *characterName == "" {
		fmt.Printf("character name is required\n")
		os.Exit(2)
	}
	if *className == "" {
		fmt.Printf("class name is required\n")
		os.Exit(2)
	}
	if *lineageName == "" {
		fmt.Printf("lineage name is required\n")
		os.Exit(2)
	}
	if *heritageName == "" {
		fmt.Printf("heritage name is required\n")
		os.Exit(2)
	}

	// Parse the traits JSON string into a map[string]string
	var selectedTraits map[string]string
	if *selectedTraitsJSON != "" {
		if err := json.Unmarshal([]byte(*selectedTraitsJSON), &selectedTraits); err != nil {
			fmt.Printf("error parsing traits JSON: %v\n", err)
			os.Exit(2)
		}
	} else {
		selectedTraits = make(map[string]string)
	}

	// Parse the talents JSON string into a []string
	var talents []string
	if *talentsJSON != "" {
		if err := json.Unmarshal([]byte(*talentsJSON), &talents); err != nil {
			fmt.Printf("error parsing talents JSON: %v\n", err)
			os.Exit(2)
		}
	}

	// Parse the languages JSON string into a []string
	var languages []string
	if *languagesJSON != "" {
		if err := json.Unmarshal([]byte(*languagesJSON), &languages); err != nil {
			fmt.Printf("error parsing languages JSON: %v\n", err)
			os.Exit(2)
		}
	}

	// Create the character request using shared types
	createReq := types.CharacterCreateRequest{
		Name:             *characterName,
		Level:            characterLevel,
		Class:            *className,
		Subclass:         *subclassName,
		Lineage:          *lineageName,
		Heritage:         *heritageName,
		Size:             *sizeName,
		AbilityGenMethod: *abilityGenMethod,
		Traits:           selectedTraits,
		Talents:          talents,
		Languages:        languages,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(createReq)
	if err != nil {
		fmt.Printf("error marshaling character data: %v\n", err)
		os.Exit(2)
	}

	// Make HTTP request to API
	apiURL := *apiBaseURL + "/api/v1/character/create"
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("error making API request: %v\n", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading API response: %v\n", err)
		os.Exit(2)
	}

	// Handle different response status codes
	switch resp.StatusCode {
	case http.StatusCreated:
		var character types.CharacterResponse
		if err := json.Unmarshal(body, &character); err != nil {
			fmt.Printf("error parsing character response: %v\n", err)
			os.Exit(2)
		}

		fmt.Printf("Successfully created character: %s\n", character.Name)
		printCharacterDetails(character)

	case http.StatusBadRequest, http.StatusConflict:
		var errorResp types.ErrorResponse
		if err := json.Unmarshal(body, &errorResp); err != nil {
			fmt.Printf("error parsing error response: %v\n", err)
			os.Exit(2)
		}
		fmt.Printf("API Error: %s\n", errorResp.Error)
		os.Exit(2)

	default:
		fmt.Printf("API request failed with status %d: %s\n", resp.StatusCode, string(body))
		os.Exit(2)
	}
}

func printCharacterDetails(character types.CharacterResponse) {
	fmt.Printf("\n=== Character Details ===\n")
	fmt.Printf("ID: %d\n", character.ID)
	fmt.Printf("Name: %s\n", character.Name)
	fmt.Printf("Level: %d\n", character.Level)
	fmt.Printf("Class: %s\n", character.Class)
	if character.Subclass != "" {
		fmt.Printf("Subclass: %s\n", character.Subclass)
	}
	fmt.Printf("Lineage: %s\n", character.Lineage)
	fmt.Printf("Heritage: %s\n", character.Heritage)
	fmt.Printf("Size: %s\n", character.Size)

	fmt.Printf("\n--- Ability Scores ---\n")
	abilities := []string{"str", "dex", "con", "int", "wis", "cha"}
	for _, ability := range abilities {
		score := character.AbilityScores[ability]
		modifier := character.AbilityModifiers[ability]
		modifierStr := strconv.Itoa(modifier)
		if modifier >= 0 {
			modifierStr = "+" + modifierStr
		}
		fmt.Printf("%s: %d (%s)\n", ability, score, modifierStr)
	}

	if len(character.Traits) > 0 {
		fmt.Printf("\n--- Traits ---\n")
		for name, description := range character.Traits {
			fmt.Printf("%s: %s\n", name, description)
		}
	}

	if len(character.Talents) > 0 {
		fmt.Printf("\n--- Talents ---\n")
		for _, talent := range character.Talents {
			fmt.Printf("- %s\n", talent)
		}
	}

	if len(character.Languages) > 0 {
		fmt.Printf("\n--- Languages ---\n")
		for _, language := range character.Languages {
			fmt.Printf("- %s\n", language)
		}
	}

	fmt.Printf("\nCreated: %s\n", character.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated: %s\n", character.UpdatedAt.Format("2006-01-02 15:04:05"))
}
