package character

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tov_tools/pkg/helpers"
)

// Helper function to create a character
func createCharacter(name, lineageKey, size string, optionalTraits map[string]string) (Character, error) {
	lineage, exists := Lineages[lineageKey]
	if !exists {
		return Character{}, fmt.Errorf("Lineage '%s' not found", lineageKey)
	}

	urbanHeritage := Heritage{
		Name:               "Urban",
		SkillProficiencies: []string{"Stealth", "Persuasion"},
		Languages:          []string{"Common", "Elvish"},
		CulturalTraits: map[string]string{
			"City Navigation": "Bonus to find your way in big cities",
		},
	}

	return Character{
		Name:         name,
		Lineage:      lineage,
		Heritage:     urbanHeritage,
		ChosenSize:   size,
		ChosenTraits: optionalTraits,
	}, nil
}

func TestCharacterCreation(t *testing.T) {
	lineageTests := []struct {
		name               string
		lineageKey         string
		predefinedTraits   []string
		selectedTraits     map[string]string
		expectedLineageSrc string
		expectedSize       string
	}{
		{
			name:               "Fang",
			lineageKey:         "beastkin",
			expectedLineageSrc: "Players Guide, pg 105",
			expectedSize:       "Medium",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["beastkin"].Traits),
			selectedTraits:     map[string]string{"Animal Instinct": "Perception", "Natural Weapons": "Claws"},
		},
		{
			name:               "Gimli",
			lineageKey:         "dwarf",
			expectedLineageSrc: "Players Guide, pg 106",
			expectedSize:       "Medium",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["dwarf"].Traits),
			selectedTraits:     map[string]string{},
		},
		{
			name:               "Legolas",
			lineageKey:         "elf",
			expectedLineageSrc: "Players Guide, pg 106",
			expectedSize:       "Medium",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["elf"].Traits),
			selectedTraits:     map[string]string{},
		},
		{
			name:               "Aragorn",
			lineageKey:         "human",
			expectedLineageSrc: "Players Guide, pg 107",
			expectedSize:       "Medium", // Or "Small" based on specific test cases
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["human"].Traits),
			selectedTraits:     map[string]string{},
		},
		{
			name:               "Tik",
			lineageKey:         "kobold",
			expectedLineageSrc: "Players Guide, pg 108",
			expectedSize:       "Small",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["kobold"].Traits),
			selectedTraits:     map[string]string{"Natural Adaptation": "Fierce (Small)"},
		},
		{
			name:               "Rog",
			lineageKey:         "orc",
			expectedLineageSrc: "Players Guide, pg 108",
			expectedSize:       "Medium",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["orc"].Traits),
			selectedTraits:     map[string]string{},
		},
		{
			name:               "Alien",
			lineageKey:         "syderean",
			expectedLineageSrc: "Players Guide, pg 109",
			expectedSize:       "Medium",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["syderean"].Traits),
			selectedTraits:     map[string]string{"Natural Adaptation": "Celestial"},
		},
		{
			name:               "Frodo",
			lineageKey:         "smallfolk",
			expectedLineageSrc: "Players Guide, pg 109",
			expectedSize:       "Small",
			predefinedTraits:   helpers.GetMapKeys(PredefinedTraitsData["smallfolk"].Traits),
			selectedTraits:     map[string]string{"Natural Adaptation": "Halfling"},
		},
	}

	for _, testLineage := range lineageTests {
		character, err := createCharacter(testLineage.name, testLineage.lineageKey, testLineage.expectedSize, testLineage.selectedTraits)
		if err != nil {
			t.Fatalf("Error creating character: %v", err)
		}

		// Check character's name
		if character.Name != testLineage.name {
			t.Errorf("Expected name to be '%s', but got '%s'", testLineage.name, character.Name)
		}

		// Check lineage
		if character.Lineage.Name != helpers.ToTitleCase(testLineage.lineageKey) {
			t.Errorf("Expected lineage name to be '%s', but got '%s'", helpers.ToTitleCase(testLineage.lineageKey), character.Lineage.Name)
		}

		// Check lineage source
		if character.Lineage.LineageSource != testLineage.expectedLineageSrc {
			t.Errorf("Expected lineage source to be '%s', but got '%s'", testLineage.expectedLineageSrc, character.Lineage.LineageSource)
		}

		// Check chosen size
		if character.ChosenSize != testLineage.expectedSize {
			t.Errorf("Expected chosen size to be '%s', but got '%s'", testLineage.expectedSize, character.ChosenSize)
		}

		// Check predefined traits
		if character.Lineage.Traits != nil {
			for _, expectedValues := range testLineage.predefinedTraits {
				assert.Equal(t, true, helpers.Contains(character.Lineage.Traits, expectedValues))
			}
		}

		// Check chosen traits
		for traitKey, expectedValue := range testLineage.selectedTraits {
			actualValue, exists := character.ChosenTraits[traitKey]
			if !exists {
				t.Errorf("Expected chosen trait '%s' not found in character's traits", traitKey)
			} else if actualValue != expectedValue {
				t.Errorf("Expected chosen trait '%s' to be '%s', but got '%s'", traitKey, expectedValue, actualValue)
			}
		}
	}
}
