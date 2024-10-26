package character

import "fmt"

// TraitChoices represents the trait options and the number to select.
type TraitChoices struct {
	NumberToSelect int
	Options        []string
}

// Lineage represents blood ties and hereditary traits
type Lineage struct {
	Name          string
	AgeDiceRoll   string
	SizeOptions   []string
	Speed         int
	TraitOptions  map[string]TraitChoices
	LineageSource string // Store where the lineage information came from
}

// PrintDetails prints detailed information about the lineage
func (l *Lineage) PrintDetails() {
	fmt.Printf("Lineage: %s\n", l.Name)
	fmt.Printf("Source: %s\n", l.LineageSource)
	fmt.Printf("Age Dice Roll: %s\n", l.AgeDiceRoll)
	fmt.Printf("Size Options: %v\n", l.SizeOptions)
	fmt.Printf("Speed: %d\n", l.Speed)
	fmt.Printf("Trait Options:\n")
	for trait, choices := range l.TraitOptions {
		fmt.Printf("  %s (%d to select): %v\n", trait, choices.NumberToSelect, choices.Options)
	}
}
