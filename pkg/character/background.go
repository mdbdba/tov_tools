package character

import (
	"fmt"
	"sort"
	"strings"
	"tov_tools/pkg/static_data"
)

type Background struct {
	Name                         string
	Description                  string
	SkillProficiencies           []string                 // predefined
	SkillProficiencyOptions      map[string]ChoiceOptions // choose x from c1, c2, ...
	AdditionalProficiencies      []string                 // predefined
	AdditionalProficiencyOptions map[string]ChoiceOptions
	Equipment                    []static_data.EquipmentPackContent // predefined
	Money                        Money
	EquipmentOptions             map[string]ChoiceOptions
	TalentOptions                map[string]ChoiceOptions
	Motivations                  map[string]map[int]string // adventuring motivation, artistic expression, secret, etc.
	BackgroundSource             string
}

// GetBackgroundByName returns a Background by its Name or an error if it doesn't exist
func GetBackgroundByName(name string) (Background, error) {
	lowerName := strings.ToLower(name)

	background, exists := Backgrounds[lowerName]
	if !exists {
		return Background{}, fmt.Errorf("background '%s' does not exist", name)
	}
	return background, nil
}

// PrintDetails prints detailed information about the Lineage
func (b *Background) PrintDetails() {

	fmt.Printf("Background: %s\n", b.Name)
	fmt.Printf("Source: %s\n", b.BackgroundSource)
	fmt.Printf("Description: %s\n", b.Description)
	fmt.Printf("Skill Proficiencies:\n")
	for _, sp := range b.SkillProficiencies {
		fmt.Printf("  %s\n", sp)
	}
	fmt.Printf("Skill Proficiency Options:\n")
	for spName, choices := range b.SkillProficiencyOptions {
		fmt.Printf("  %s (%d to select): %v\n", spName, choices.NumberToSelect, choices.Options)
	}
	fmt.Printf("Additional Proficiencies:\n")
	for _, ap := range b.AdditionalProficiencies {
		fmt.Printf("  %s\n", ap)
	}
	fmt.Printf("Additional Proficiency Options:\n")
	for apName, choices := range b.AdditionalProficiencyOptions {
		fmt.Printf("  %s (%d to select): %v\n", apName, choices.NumberToSelect, choices.Options)
	}
	fmt.Printf("Equipment:\n")
	for _, eq := range b.Equipment {
		fmt.Printf("  %s\n", eq.Name)
	}
	fmt.Printf("Equipment Options:\n")
	for eqName, choices := range b.EquipmentOptions {
		fmt.Printf("  %s (%d to select): %v\n", eqName, choices.NumberToSelect, choices.Options)
	}
	fmt.Printf("Money:\n  ")
	if b.Money.GoldPieces > 0 {
		fmt.Printf("GP: %d - ", b.Money.GoldPieces)
	}
	if b.Money.SilverPieces > 0 {
		fmt.Printf("SP: %d - ", b.Money.SilverPieces)
	}
	if b.Money.CopperPieces > 0 {
		fmt.Printf("SP: %d - ", b.Money.CopperPieces)
	}

	fmt.Printf("\nTalent Options:\n")
	for talent, choices := range b.TalentOptions {
		fmt.Printf("  %s (%d to select): %v\n", talent, choices.NumberToSelect, choices.Options)
	}

	fmt.Printf("Motivations:\n")
	for motivationType, choices := range b.Motivations {
		fmt.Printf("  %s:\n", motivationType)
		rollIds := make([]int, 0, len(choices))
		for rollId := range choices {
			rollIds = append(rollIds, rollId)
		}
		sort.Ints(rollIds)

		// Print in sorted order
		for _, rollId := range rollIds {
			fmt.Printf("  %d: %s\n", rollId, choices[rollId])
		}

	}

}
