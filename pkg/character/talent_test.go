package character

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func TestSkillBonusTalent(t *testing.T) {
	// Define a talent that doubles the proficiency bonus for the Arcana skill
	talentArcaneMind := Talent{
		Name:        "Arcane Mind",
		Category:    "magic",
		Description: "Double your proficiency bonus for any ability check that uses the Arcana skill.",
		Prerequisite: func(c *Character) bool {
			return c.Level >= 3 // Requires Level 3+
		},
		Benefits: []Benefit{
			&SkillBonusMultiplierBenefit{
				SkillName:       "Arcana",
				BonusMultiplier: 2.0, // Double proficiency bonus
			},
		},
	}

	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	// sortOrder := []string{"dex", "con", "str", "cha", "wis", "int"}
	rollingOption := "common"
	// Create a test character
	testCharacter := NewCharacter(
		"Test Mage", 3, "Wizard", "battle mage",
		Lineage{}, Heritage{}, "medium", rollingOption, map[string]string{}, []string{},
		"Standard", "Character talent test", observedLoggerSugared)

	testCharacter.BaseSkillBonus = map[string]int{"Arcana": 0}

	// Add the talent to the character
	err := testCharacter.AddTalent(talentArcaneMind)
	if err != nil {
		t.Fatalf("failed to add talent: %v", err)
	}

	// Expected Arcana bonus: base proficiency bonus * 2
	expectedBonus := int(testCharacter.GetBaseProficiencyBonus() * 2)
	actualBonus := testCharacter.BaseSkillBonus["Arcana"]

	// Verify the result
	if actualBonus != expectedBonus {
		t.Errorf("Expected Arcana bonus to be %d, but got %d", expectedBonus, actualBonus)
	}
}

func TestFlatBonusTalent(t *testing.T) {
	talentStrongArm := Talent{
		Name:        "Strong Arm",
		Category:    "martial",
		Description: "Gain a +2 bonus to Strength.",
		Prerequisite: func(c *Character) bool {
			return c.Level >= 1 // No significant restrictions
		},
		Benefits: []Benefit{
			&FlatBonusBenefit{
				Attribute: "str",
				Bonus:     2,
			},
		},
	}
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "common"
	// Create a test character
	testCharacter := NewCharacter(
		"Test Fighter", 1, "Fighter", "weapon master",
		Lineage{}, Heritage{}, "medium", rollingOption, map[string]string{}, []string{},
		"Standard", "Character talent test", observedLoggerSugared)

	// Add the talent to the character
	err := testCharacter.AddTalent(talentStrongArm)
	if err != nil {
		t.Fatalf("failed to add talent: %v", err)
	}

	// Verify that the flat bonus was applied
	strengthScore := testCharacter.Abilities.Values["str"]
	// fmt.Println(testCharacter.Abilities.ToPrettyString())
	expectedScore := testCharacter.Abilities.Base["str"] + 2 // 13 + 2 bonus
	if strengthScore != expectedScore {
		t.Errorf("Expected Strength score to be %d, but got %d", expectedScore, strengthScore)
	}
}

func TestSpellSwapTalent(t *testing.T) {
	// Create a test talent that swaps Firebolt for Ray of Frost
	talentVersatileSpellcaster := Talent{
		Name:        "Versatile Spellcaster",
		Category:    "magic",
		Description: "Swap the spell 'Firebolt' with 'Ray of Frost'.",
		Prerequisite: func(c *Character) bool {
			// Prerequisite: Wizard class, level >= 5
			return c.CharacterClass == "Wizard" && c.Level >= 5
		},
		Benefits: []Benefit{
			&SpellSwapBenefit{
				OldSpell: "Firebolt",
				NewSpell: "Ray of Frost",
			},
		},
	}

	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "common"
	// Create a test character
	testCharacter := NewCharacter(
		"Test Wizard", 5, "Wizard", "battle mage",
		Lineage{}, Heritage{}, "medium", rollingOption, map[string]string{}, []string{},
		"Standard", "Character talent test", observedLoggerSugared)

	testCharacter.SpellBook = []string{"Firebolt", "Mage Armor"}

	// Add the talent to the character
	err := testCharacter.AddTalent(talentVersatileSpellcaster)
	if err != nil {
		t.Fatalf("failed to add talent: %v", err)
	}

	// Verify that Firebolt is replaced with Ray of Frost
	foundNewSpell := false
	foundOldSpell := false
	for _, spell := range testCharacter.SpellBook {
		if spell == "Ray of Frost" {
			foundNewSpell = true
		}
		if spell == "Firebolt" {
			foundOldSpell = true
		}
	}

	if !foundNewSpell {
		t.Errorf("Expected 'Ray of Frost' to be in the character's spellbook, but it was not found")
	}

	if foundOldSpell {
		t.Errorf("Expected 'Firebolt' to be removed from the character's spellbook, but it is still present")
	}
}
