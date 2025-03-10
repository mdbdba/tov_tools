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

// Languages returns a map[string]string holding all language choices
var Languages = func() map[string]string {
	return map[string]string{
		"Common":         "Spoken by most humanoids, commonly used for trade and communication.",
		"Elvish":         "The language of the elves, characterized by its melodic tones.",
		"Dwarvish":       "Spoken by dwarves, with a harsh and ancient tone.",
		"Orcish":         "A guttural language spoken by orcs and related tribes.",
		"Draconic":       "The ancient language of dragons, also used in magical writings.",
		"Abyssal":        "The twisted tongue of demons and other chaotic fiends.",
		"Infernal":       "A structured and precise language spoken by devils.",
		"Celestial":      "The radiant language of celestials, beings of the upper planes.",
		"Gnomish":        "Spoken by gnomes, filled with inventive and curious terms.",
		"Halfling":       "The casual and friendly language spoken by halflings.",
		"Sylvan":         "The ancient tongue of fey and forest creatures.",
		"Giant":          "The booming and straightforward language of giants.",
		"Machine Speech": "The language of machines, with a mysterious and unintelligible tone.",
		"Primordial":     "The language of the elemental beings.",
		"Undercommon":    "A secretive language used in the Underdark.",
	}
}

// LanguageNames returns the names of available languages
var LanguageNames = func(exceptions ...[]string) []string {
	var names []string
	exceptionMap := make(map[string]bool)

	// Process exceptions if they are provided
	if len(exceptions) > 0 && len(exceptions[0]) > 0 {
		for _, ex := range exceptions[0] {
			exceptionMap[ex] = true
		}
	}

	// Iterate over the languages and exclude the ones in exceptionMap
	for name := range Languages() {
		if !exceptionMap[name] { // Only add if it's not in the exception list
			names = append(names, name)
		}
	}
	return names
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
	LanguageDefaults    []string
	LanguageSuggestions []string
	Traits              map[string]string // predefined
	TraitOptions        map[string]TraitChoices
	HeritageSource      string
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
