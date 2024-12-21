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
		"Test Mage", 3, "Wizard", Lineage{}, Heritage{}, "medium", rollingOption,
		map[string]string{}, "Standard", "Character talent test", observedLoggerSugared)

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
		"Test Fighter", 1, "Fighter", Lineage{}, Heritage{}, "medium", rollingOption,
		map[string]string{}, "Standard", "Character talent test", observedLoggerSugared)

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
