package character

import "fmt"

// Character represents a character in the game
type Character struct {
	Name         string
	Lineage      Lineage
	Heritage     Heritage
	ChosenSize   string
	ChosenTraits map[string]string
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
