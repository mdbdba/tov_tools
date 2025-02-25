package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLanguageSuggestions(t *testing.T) {
	expected := []LanguageSuggestion{
		{
			LanguageType:      "standard",
			Speakers:          []string{"human"},
			SuggestedLanguage: "Common",
			Script:            "Common",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"dwarf"},
			SuggestedLanguage: "Dwarvish",
			Script:            "Dwarvish",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"elf"},
			SuggestedLanguage: "Elvish",
			Script:            "Elvish",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"ogre", "giant"},
			SuggestedLanguage: "Giant",
			Script:            "Dwarvish",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"gnome"},
			SuggestedLanguage: "Gnomish",
			Script:            "Common",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"goblin"},
			SuggestedLanguage: "Goblin",
			Script:            "Dwarvish",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"halfling"},
			SuggestedLanguage: "Halfling",
			Script:            "Common",
		},
		{
			LanguageType:      "standard",
			Speakers:          []string{"orc"},
			SuggestedLanguage: "Orcish",
			Script:            "Dwarvish",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"demon"},
			SuggestedLanguage: "Abyssal",
			Script:            "Infernal",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"celestial"},
			SuggestedLanguage: "Celestial",
			Script:            "Celestial",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"dragon", "dragonborn", "kobold"},
			SuggestedLanguage: "Draconic",
			Script:            "Draconic",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"devil"},
			SuggestedLanguage: "Infernal",
			Script:            "Infernal",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"mechadrons"},
			SuggestedLanguage: "Machine Speech",
			Script:            "Unreadable by non-constructs",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"elemental"},
			SuggestedLanguage: "Primordial",
			Script:            "Dwarvish",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"fey"},
			SuggestedLanguage: "Sylvan",
			Script:            "Elvish",
		},
		{
			LanguageType:      "esoteric",
			Speakers:          []string{"underworld trader"},
			SuggestedLanguage: "Undercommon",
			Script:            "Elvish",
		},
	}

	// Execute the LanguageSuggestions function
	actual := LanguageSuggestions()

	// Compare the entire actual output with the expected output
	assert.Equal(t, expected, actual)

}

func TestHeritageSuggestion(t *testing.T) {
	expected := map[string][]string{
		"Beastkin":  {"Slayer", "Wildlands"},
		"Dwarf":     {"Fireforge", "Stone"},
		"Elf":       {"Cloud", "Grove"},
		"Human":     {"Cosmopolitan", "Nomadic"},
		"Kobold":    {"Supplicant", "Salvager"},
		"Orc":       {"Diaspora", "Slayer"},
		"Syderean":  {"Anointed", "Vexed"},
		"Smallfolk": {"Cottage", "Salvager"},
	}

	// Execute the HeritageSuggestion function
	actual := HeritageSuggestion()

	// Check if the expected and actual maps match
	assert.Equal(t, expected, actual)
}
