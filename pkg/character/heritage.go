package character

import (
	"fmt"
	"strings"
)

type LanguageSuggestion struct {
	LanguageType      string
	Speakers          []string
	SuggestedLanguage string
	Script            string
}

// LanguageSuggestions returns a recommended languages by archetype
var LanguageSuggestions = func() []LanguageSuggestion {
	return []LanguageSuggestion{
		LanguageSuggestion{
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
}

// HeritageSuggestion returns a recommended heritages by lineage
var HeritageSuggestion = func() map[string][]string {
	return map[string][]string{
		"Beastkin":  {"Slayer", "Wildlands"},
		"Dwarf":     {"Fireforge", "Stone"},
		"Elf":       {"Cloud", "Grove"},
		"Human":     {"Cosmopolitan", "Nomadic"},
		"Kobold":    {"Supplicant", "Salvager"},
		"Orc":       {"Diaspora", "Slayer"},
		"Syderean":  {"Anointed", "Vexed"},
		"Smallfolk": {"Cottage", "Salvager"},
	}
}

// Heritage represents upbringing and cultural traits
type Heritage struct {
	Name                string
	SkillProficiencies  []string // e.g., ["Stealth", "Arcana"]
	LanguageDefaults    []string
	LanguageSuggestions []string
	LanguageChoices     TraitChoices
	Languages           []string
	CulturalTraits      map[string]string // e.g., "City Navigation": "Bonus to find your way in big cities"
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
