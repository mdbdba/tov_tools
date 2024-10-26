package character

import (
	"testing"
)

func TestCharacterCreation(t *testing.T) {
	beastkinLineage, exists := Lineages["beastkin"]
	if !exists {
		t.Fatalf("Lineage 'beastkin' not found")
	}

	urbanHeritage := Heritage{
		Name:               "Urban",
		SkillProficiencies: []string{"Stealth", "Persuasion"},
		Languages:          []string{"Common", "Elvish"},
		CulturalTraits: map[string]string{
			"City Navigation": "Bonus to find your way in big cities",
		},
	}

	chosenSize := "Medium"
	chosenTraits := map[string]string{
		"Animal Instinct":    "Perception",
		"Natural Weapons":    "Claws",
		"Natural Adaptation": "Avian",
	}

	myCharacter := Character{
		Name:         "Fang",
		Lineage:      beastkinLineage,
		Heritage:     urbanHeritage,
		ChosenSize:   chosenSize,
		ChosenTraits: chosenTraits,
	}

	// Check character's name
	if myCharacter.Name != "Fang" {
		t.Errorf("Expected name to be 'Fang', but got '%s'", myCharacter.Name)
	}

	// Check lineage
	if myCharacter.Lineage.Name != "Beastkin" {
		t.Errorf("Expected lineage name to be 'Beastkin', but got '%s'", myCharacter.Lineage.Name)
	}

	// Check lineage source
	tSearchString := "Players Guide, pg 105"
	if myCharacter.Lineage.LineageSource != tSearchString {
		t.Errorf("Expected lineage source to be '%s', but got '%s'", tSearchString, myCharacter.Lineage.LineageSource)
	}

	// Check chosen size
	if myCharacter.ChosenSize != "Medium" {
		t.Errorf("Expected chosen size to be 'Medium', but got '%s'", myCharacter.ChosenSize)
	}

	// Check chosen traits
	if myCharacter.ChosenTraits["Animal Instinct"] != "Perception" {
		t.Errorf("Expected 'Animal Instinct' trait to be 'Perception', but got '%s'", myCharacter.ChosenTraits["Animal Instinct"])
	}
	if myCharacter.ChosenTraits["Natural Weapons"] != "Claws" {
		t.Errorf("Expected 'Natural Weapons' trait to be 'Claws', but got '%s'", myCharacter.ChosenTraits["Natural Weapons"])
	}
	if myCharacter.ChosenTraits["Natural Adaptation"] != "Avian" {
		t.Errorf("Expected 'Natural Adaptation' trait to be 'Avian', but got '%s'", myCharacter.ChosenTraits["Natural Weapons"])
	}
}

func TestPrintDetails(t *testing.T) {
	beastkinLineage, exists := Lineages["beastkin"]
	if !exists {
		t.Fatalf("Lineage 'beastkin' not found")
	}

	urbanHeritage := Heritage{
		Name:               "Urban",
		SkillProficiencies: []string{"Stealth", "Persuasion"},
		Languages:          []string{"Common", "Elvish"},
		CulturalTraits: map[string]string{
			"City Navigation": "Bonus to find your way in big cities",
		},
	}

	chosenSize := "Medium"
	chosenTraits := map[string]string{
		"Animal Instinct": "Perception",
		"Natural Weapons": "Claws",
	}

	myCharacter := Character{
		Name:         "Fang",
		Lineage:      beastkinLineage,
		Heritage:     urbanHeritage,
		ChosenSize:   chosenSize,
		ChosenTraits: chosenTraits,
	}

	myCharacter.PrintDetails() // This is primarily a visual check
}
