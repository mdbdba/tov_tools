package character

import (
	"fmt"
	"go.uber.org/zap"
)

// Character represents a character in the game
type Character struct {
	Name           string
	Level          int
	CharacterClass string
	Lineage        Lineage
	Heritage       Heritage
	ChosenSize     string
	ChosenTraits   map[string]string
	Abilities      AbilityArray
	RollingOption  string
}

// NewCharacter Method to create a new character with default properties
func NewCharacter(
	name string,
	level int,
	characterClass string,
	lineage Lineage,
	heritage Heritage,
	chosenSize string,
	rollingOption string,
	chosenTraits map[string]string,
	ctxRef string,
	logger *zap.SugaredLogger) *Character {

	useClass := Class{}
	err := error(nil)

	if characterClass != "" {
		useClass, err = GetClassByName(characterClass)
		if err != nil {
			fmt.Printf("Error getting class '%s': %v\n", characterClass, err)
			fmt.Println("Using random selection instead.")
			useClass = RandomClass()
		}
	} else {
		fmt.Println("No class specified. Using random selection instead.")
		useClass = RandomClass()
	}

	// TODO: Implement these
	LevelChangeIncrease := AbilityArrayTemplate()
	AdditionalBonus := AbilityArrayTemplate()

	a, err := GetAbilityArray(rollingOption, useClass.AbilityScoreOrderPreference, LevelChangeIncrease,
		AdditionalBonus, ctxRef, false, logger)

	if err != nil {
		panic(err)
	}

	character := &Character{
		Name:           name,
		Level:          level,
		CharacterClass: characterClass,
		Lineage:        lineage,
		Heritage:       heritage,
		ChosenSize:     chosenSize,
		ChosenTraits:   chosenTraits,
		Abilities:      *a,
	}
	return character
}

// Initializes a default AbilityArray,
// Modify rolling options and other parameters based on your RPG rules.
func createDefaultAbilities() AbilityArray {
	// Consider default parameters, adjust or prompt user/UI for input
	rollingOption := "standard"                                     // Default to standard or based on user choice
	sortOrder := []string{"str", "dex", "con", "int", "wis", "cha"} // Customize order as needed

	// Creating a default AbilityArray; passing empty values as per your initial setup
	abilityArray, _ := GetAbilityArray(rollingOption, sortOrder, map[string]int{},
		map[string]int{}, "initial character creation", false, nil)
	return *abilityArray
}

// PrintDetails prints detailed information about the character
func (c *Character) PrintDetails() {
	fmt.Printf("Character: %s\n", c.Name)
	c.Lineage.PrintDetails()
	fmt.Printf("Chosen Size: %s\n", c.ChosenSize)
	fmt.Println("Chosen Traits:")
	for traitType, trait := range c.ChosenTraits {
		fmt.Printf("  %s: %s\n", traitType, trait)
	}
	fmt.Printf("Heritage: %s, Languages: %v, Cultural Trait: %s\n",
		c.Heritage.Name, c.Heritage.Languages, c.Heritage.CulturalTraits["City Navigation"])
}
