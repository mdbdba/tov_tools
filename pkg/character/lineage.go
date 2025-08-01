package character

import (
	"fmt"
	"strings"
)

// LineagePreDefinedTraits holds descriptions for predefined Traits of a Lineage.
type LineagePreDefinedTraits struct {
	Lineage string
	Traits  map[string]string
}

type LineageNaturalAdaptationTraitDescriptions struct {
	Lineage string
	Traits  map[string]map[string]string
}

// Lineage represents blood ties and hereditary Traits
type Lineage struct {
	Name          string
	MaturityAge   int // The age at which the character reaches maturity
	AgeDiceSides  int // The number of dice sides to roll for age determination
	AgeDiceRolls  int // The number of dice rolls for determining the additional age
	SizeOptions   []string
	Speed         int
	Traits        []string // Predefined
	TraitOptions  map[string]ChoiceOptions
	LineageSource string // Store where the Lineage information came from
}

// GetLineageByName returns a Lineage by its Name or an error if it doesn't exist
func GetLineageByName(name string) (Lineage, error) {
	lowerName := strings.ToLower(name)
	// fmt.Printf("Looking for lineage %s (lower: %s)\n", name, lowerName)
	lineage, exists := Lineages[lowerName]
	if !exists {
		return Lineage{}, fmt.Errorf("Lineage '%s' does not exist", name)
	}
	return lineage, nil
}

// PrintDetails prints detailed information about the Lineage
func (l *Lineage) PrintDetails() {
	fmt.Printf("Lineage: %s\n", l.Name)
	fmt.Printf("Source: %s\n", l.LineageSource)
	fmt.Printf("Maturity Age: %d\n", l.MaturityAge)
	fmt.Printf("Age Dice Sides: %d\n", l.AgeDiceSides)
	fmt.Printf("Age Dice Rolls: %d\n", l.AgeDiceRolls)
	fmt.Printf("Size Options: %v\n", l.SizeOptions)
	fmt.Printf("Speed: %d\n", l.Speed)
	fmt.Printf("Traits:\n")
	for _, trait := range l.Traits {
		if lineageTraits, ok := PredefinedTraitsData[strings.ToLower(l.Name)]; ok {
			if description, exists := lineageTraits.Traits[trait]; exists {
				fmt.Printf("  %s: %s\n", trait, description)
			} else {
				fmt.Printf("  %s\n", trait)
			}
		} else {
			fmt.Printf("  %s\n", trait)
		}
	}
	fmt.Printf("Trait Options:\n")
	for trait, choices := range l.TraitOptions {
		fmt.Printf("  %s (%d to select): %v\n", trait, choices.NumberToSelect, choices.Options)
	}
}
