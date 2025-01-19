package character

import (
	"fmt"
	"strings"
)

// Heritage represents upbringing and cultural traits
type Heritage struct {
	Name               string
	SkillProficiencies []string // e.g., ["Stealth", "Arcana"]
	Languages          []string
	CulturalTraits     map[string]string // e.g., "City Navigation": "Bonus to find your way in big cities"
}

// GetHeritageByName returns a Heritage by its name or an error if it doesn't exist
func GetHeritageByName(name string) (Heritage, error) {
	lowerName := strings.ToLower(name)
	heritage, exists := Heritages[lowerName]
	if !exists {
		return Heritage{}, fmt.Errorf("lineage '%s' does not exist", name)
	}
	return heritage, nil
}
