package character

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
	"tov_tools/pkg/helpers"
)

func TestCharacterCreation(t *testing.T) {

	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

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
			selectedTraits: map[string]string{"Natural Adaptation": "Agile",
				"Animal Instinct": "Perception",
				"Natural Weapons": "Claws"},
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
			expectedSize:       "Small", // Or "Small" based on specific test cases
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
		//character, err := NewCharacter(testLineage.name, testLineage.lineageKey, testLineage.expectedSize, testLineage.selectedTraits)

		ctxRef := fmt.Sprintf("Character lineage test: %s", testLineage.lineageKey)
		// Create a character to test against
		character, err := NewCharacter(
			testLineage.name, 5, "Wizard",
			"battle mage",
			testLineage.lineageKey, Heritage{},
			Lineages[testLineage.lineageKey].SizeOptions[0], rollingOption,
			testLineage.selectedTraits, []string{},
			"Standard", ctxRef, observedLoggerSugared)

		if err != nil {
			t.Fatalf("Error creating character: %v", err)
		}

		// Check character's name
		if character.Name != testLineage.name {
			t.Errorf("Expected name to be '%s', but got '%s'", testLineage.name, character.Name)
		}

		// Check lineage
		if character.Lineage.Name != helpers.ToTitleCase(testLineage.lineageKey) {
			t.Errorf("Expected lineage name to be '%s', but got '%s' for %s",
				helpers.ToTitleCase(testLineage.lineageKey), character.Lineage.Name,
				testLineage.name)
		}

		// Check lineage source
		if character.Lineage.LineageSource != testLineage.expectedLineageSrc {
			t.Errorf("Expected lineage source to be '%s', but got '%s' for %s",
				testLineage.expectedLineageSrc, character.Lineage.LineageSource,
				testLineage.name)
		}

		// Check chosen size
		if character.ChosenSize != testLineage.expectedSize {
			t.Errorf("Expected chosen size to be '%s', but got '%s' for %s",
				testLineage.expectedSize, character.ChosenSize,
				testLineage.name)
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
			fmt.Printf("Trait: %s, Expected: %s, Actual: %s\n", traitKey, expectedValue, actualValue)
			if !exists {
				t.Errorf("Expected chosen trait '%s' not found in character's traits for %s",
					traitKey, testLineage.name)
			} else if actualValue != expectedValue {
				t.Errorf("Expected chosen trait '%s' to be '%s', but got '%s' for %s",
					traitKey, expectedValue, actualValue, testLineage.name)
			}
		}
	}
}

func TestSetAbilitySkills(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "common"
	ctxRef := "Test SetAbilitySkills"
	// Create a test character
	testCharacter, err := NewCharacter(
		"Test Wizard", 5, "Wizard",
		"battle mage", "human",
		Heritage{}, "Medium", rollingOption, map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)
	assert.NoError(t, err, "Unexpected error when creating test character")
	testCharacter.SkillProficiencies = []AbilitySkillProficiency{
		{Skill: "athletics", Source: "Training"},
	}
	testCharacter.SkillBonus = map[string]map[string]AbilitySkillBonus{
		"athletics": {
			"Training": {
				Bonus:  1,
				Source: "Training",
			},
			"Magic Belt": {
				Bonus:  1,
				Source: "Magic Belt",
			},
		},
		"acrobatics": {
			"Training": {
				Bonus:  2,
				Source: "Training",
			},
		},
	}

	// Run the function
	testCharacter.SetAbilitySkills()

	testCharacter.PrintDetails()
	// Define test cases
	tests := []struct {
		Skill              string
		ExpectedValue      int
		ExpectedProficient bool
	}{
		{
			Skill:              "athletics",
			ExpectedValue:      testCharacter.Abilities.Modifiers["str"] + 5, //(strength mod) + 3 (proficiency) + 1 + 1 (bonuses)
			ExpectedProficient: true,
		},
		{
			Skill:              "acrobatics",
			ExpectedValue:      testCharacter.Abilities.Modifiers["dex"] + 2, //  (dex mod) + 2 (bonuses from training [non-proficient])
			ExpectedProficient: false,
		},
		{
			Skill:              "history",
			ExpectedValue:      testCharacter.Abilities.Modifiers["int"], // (intelligence mod, no proficiency, no bonuses)
			ExpectedProficient: false,
		},
	}

	// Validate results
	for _, test := range tests {
		abilitySkill, exists := testCharacter.AbilitySkills[test.Skill]
		if !exists {
			t.Errorf("Skill %s not found in AbilitySkills", test.Skill)
			continue
		}
		bonusStr := ""
		separator := ""
		for i := range testCharacter.SkillBonus[test.Skill] {
			bonusStr += spew.Sprintf("%s%d", separator, testCharacter.SkillBonus[test.Skill][i].Bonus)
			if separator == "" {
				separator = " + "
			}
		}

		pbStr := ""
		for _, p := range testCharacter.SkillProficiencies {
			if p.Skill == test.Skill {
				pbStr = fmt.Sprintf(" + pb: %d", testCharacter.GetProficiencyBonus())
			}
			break
		}

		if abilitySkill.Value != test.ExpectedValue {
			t.Errorf("Skill %s [%s]: expected value %d, got %d = skill mod: %d%s + bonuses: %s",
				test.Skill, abilitySkill.Ability, test.ExpectedValue, abilitySkill.Value,
				testCharacter.Abilities.Modifiers[abilitySkill.Ability],
				pbStr,
				bonusStr)
		}

		if abilitySkill.Proficient != test.ExpectedProficient {
			t.Errorf("Skill %s: expected proficient %v, got %v", test.Skill, test.ExpectedProficient, abilitySkill.Proficient)
		}
	}
}

func TestAbilityUpdateReflectsEverywhere(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"
	ctxRef := "Character Update test"

	c, err := NewCharacter("Test Fighter", 1, "Fighter", "weapon master",
		"human", Heritage{}, "Medium", rollingOption, map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)

	assert.NoError(t, err, "Unexpected error when creating test character")
	fmt.Println("BEFORE")
	c.PrintDetails()
	err = c.IncreaseAbility("dex")
	if err != nil {
		t.Fatalf("failed to increase ability: %v", err)
	}
	err = c.IncreaseAbility("dex")
	if err != nil {
		t.Fatalf("failed to increase ability: %v", err)
	}

	fmt.Println("AFTER")
	c.PrintDetails()

	if c.InitiativeBonus != 2 {
		fmt.Println(c.Abilities.ToPrettyString())
		t.Errorf("Expected InitiativeBonus to be 2, got %d", c.InitiativeBonus)
	}

	acrobatics := c.GetSkillBonus("acrobatics")
	if acrobatics != 2 {
		t.Errorf("Expected acrobatics to be recalculated, got %d", acrobatics)
	}
}

func TestInvalidCharacterCreation(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

	invalidLineageTests := []struct {
		name       string
		lineageKey string
	}{
		{"Invalid1", "unknown_lineage"},
		{"Invalid2", ""},
	}

	for _, tc := range invalidLineageTests {
		ctxRef := fmt.Sprintf("Character invalid lineage test: %s", tc.lineageKey)

		_, err := NewCharacter(
			"Test Wizard", 5, "Wizard",
			"battle mage",
			tc.lineageKey, Heritage{},
			"medium", rollingOption,
			map[string]string{}, []string{},
			"Standard", ctxRef, observedLoggerSugared)

		assert.Error(t, err, "Expected error when creating character with invalid size")
	}
}

func TestCharacterWithNoTraits(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"
	ctxRef := "Character no traits test"

	//character, err := createCharacter("Lineless", "elf", "Medium", nil)
	// Create a test character
	character, err := NewCharacter(
		"Mr NoTraits", 1, "ranger", "pack master",
		"human", Heritage{},
		"Medium", rollingOption,
		map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)

	assert.NoError(t, err, "Unexpected error when creating character with no traits")
	assert.Equal(t, 0, len(character.ChosenTraits), "Expected no chosen traits")
}

func TestCharacterWithEdgeCaseNames(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

	edgeCaseNames := []struct {
		name       string
		lineageKey string
	}{
		{"", "human"},
		{"@InvalidName!", "dwarf"},
		{"1234", "elf"},
	}

	for _, tc := range edgeCaseNames {
		ctxRef := fmt.Sprintf("Character name edge case test: %s", tc.name)

		_, err := NewCharacter(
			tc.name, 1, "Wizard", "battle mage",
			tc.lineageKey, Heritage{},
			Lineages[tc.lineageKey].SizeOptions[0], rollingOption,
			map[string]string{}, []string{},
			"Standard", ctxRef, observedLoggerSugared)

		assert.Error(t, err, fmt.Sprintf("Expected error when creating %s character with invalid name", tc.lineageKey))
	}
}

func TestCharacterWithEdgeCaseSizes(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

	edgeCaseSizes := []struct {
		name       string
		lineageKey string
		size       string
	}{
		{"bob", "human", "Tiny"},
		{"sally", "human", "Huge"},
	}
	for _, tc := range edgeCaseSizes {
		_, err := NewCharacter(
			tc.name, 1, "Wizard",
			"battle mage",
			tc.lineageKey, Heritage{},
			tc.size, rollingOption,
			map[string]string{}, []string{},
			"Standard", "Character Size Edge Case", observedLoggerSugared)
		if err == nil {
			t.Errorf("Character creation should have failed for name '%s' tested size: %s", tc.name, tc.size)
		}
		assert.Error(t, err, fmt.Sprintf("Expected error when creating %s character with invalid size", tc.lineageKey))
	}
}

func TestInvalidInputsForNewCharacter(t *testing.T) {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()

	// Define test cases for invalid inputs
	invalidInputs := []struct {
		testCase      string
		name          string
		level         int
		class         string
		size          string
		rollingOption string
		expectError   bool
	}{
		{"Negative Level", "TestInvalid", -1, "Wizard", "Medium", "standard", true},
		{"Excessive Level", "TestInvalid", 101, "Wizard", "Medium", "standard", true},
		{"Invalid Class", "TestInvalid", 5, "InvalidClass", "Medium", "standard", true},
		{"Invalid Size", "TestInvalid", 5, "Wizard", "Giant", "standard", true},
		{"Invalid Rolling Option", "TestInvalid", 5, "Wizard", "Medium", "invalidOption", true},
		{"Empty Name", "", 5, "Wizard", "Medium", "standard", true},
	}

	for _, tc := range invalidInputs {
		t.Run(tc.testCase, func(t *testing.T) {
			ctxRef := "Invalid Inputs Test"
			_, err := NewCharacter(
				tc.name, tc.level, tc.class, "battle mage",
				"human", Heritage{}, tc.size, tc.rollingOption, map[string]string{}, []string{},
				"Standard", ctxRef, observedLoggerSugared)

			if tc.expectError {
				assert.Error(t, err, "Expected error for invalid input")
			} else {
				assert.NoError(t, err, "Unexpected error for valid input")
			}
		})
	}
}

func TestHitPointGenerationAtCreation(t *testing.T) {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()

	hitPointTests := []struct {
		name     string
		class    string
		subClass string
		level    int
	}{
		{"Level One Wizard", "wizard", "battle mage", 1},
		{"Level Five Fighter", "fighter", "weapon master", 5},
		{"Level Two Cleric", "cleric", "life domain", 2},
		{"Level Three Rogue", "rogue", "enforcer", 3},
		{"Level Ten Barbarian", "barbarian", "berserker", 10},
		{"Level Twenty Ranger", "ranger", "hunter", 20},
	}

	for _, tc := range hitPointTests {
		t.Run(tc.name, func(t *testing.T) {
			ctxRef := fmt.Sprintf("Hit Point Generation Test: %s", tc.name)

			character, err := NewCharacter(
				tc.name, tc.level, tc.class, tc.subClass,
				"human", Heritage{}, "Medium", "standard", map[string]string{}, []string{},
				"Standard", ctxRef, observedLoggerSugared)

			if err != nil {
				t.Fatalf("Error creating character %s: %v", tc.name, err)
			}

			assert.NotNil(t, character.HitPointAudit, "HitPointAudit should not be nil")
			tmpTotal := 0
			tmpRolls := 0
			for _, v := range character.HitPointAudit {
				// fmt.Printf("Rolls: %v, Result: %d Ctx: %s\n", len(v.RollsUsed), v.Result, v.CtxRef)
				tmpTotal += v.Result
				tmpRolls += len(v.RollsUsed)
			}

			assert.NoError(t, err, "Unexpected error for valid input")
			assert.Equal(t, tmpTotal, character.MaxHitPoints, "Incorrect total HP calculation")
			assert.Equal(t, tc.level, tmpRolls, "HP rolls should match level")
		})
	}
}
func TestTemporaryHitPoints(t *testing.T) {
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()

	// Create a character
	ctxRef := "Temporary HP Test"
	character, err := NewCharacter(
		"Temp HP Tester", 4, "Warlock", "fiend",
		"human", Heritage{}, "Small", "standard", map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)
	assert.NoError(t, err, "Unexpected error when creating character")

	// Add temporary HP
	tempHP := 10
	character.ModifyTemporaryHitPoints(tempHP)

	// Verify temporary HP
	assert.Equal(t, tempHP, character.TemporaryHitPoints, "Temporary HP not added correctly")
	assert.Equal(t, tempHP+character.MaxHitPoints, character.GetTotalHitPoints(), "Total HP not correct")

	// Apply damage that exceeds temporary HP
	// this is not implemented yet
	// damage := 12
	// character.ApplyDamage(damage)

	// Verify remaining HP and temporary HP
	// assert.Equal(t, tempHP-damage, character.HitPoints.Temp, "Temporary HP not reduced correctly")
	// assert.Equal(t, initialHP, character.HitPoints.Total, "Regular HP should remain unchanged after temporary HP absorbs damage")
}
