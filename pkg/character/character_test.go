package character

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
	"tov_tools/pkg/dice"
	"tov_tools/pkg/helpers"
)

func TestCharacterCreation(t *testing.T) {

	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

	testTraits := []struct {
		name                    string
		lineageKey              string
		heritageKey             string
		lineagePredefinedTraits []string
		lineageSelectedTraits   map[string]string
		lineageExpectedSrc      string
		lineageExpectedSize     string
	}{
		{
			name:                    "Fang",
			lineageKey:              "beastkin",
			heritageKey:             "slayer",
			lineageExpectedSrc:      "Players Guide, pg 105",
			lineageExpectedSize:     "Medium",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["beastkin"].Traits),
			lineageSelectedTraits: map[string]string{"Natural Adaptation": "Agile",
				"Animal Instinct": "Perception",
				"Natural Weapons": "Claws"},
		},
		{
			name:                    "Gimli",
			lineageKey:              "dwarf",
			heritageKey:             "fireforge",
			lineageExpectedSrc:      "Players Guide, pg 106",
			lineageExpectedSize:     "Medium",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["dwarf"].Traits),
			lineageSelectedTraits:   map[string]string{},
		},
		{
			name:                    "Legolas",
			lineageKey:              "elf",
			heritageKey:             "cloud",
			lineageExpectedSrc:      "Players Guide, pg 106",
			lineageExpectedSize:     "Medium",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["elf"].Traits),
			lineageSelectedTraits:   map[string]string{},
		},
		{
			name:                    "Aragorn",
			lineageKey:              "human",
			heritageKey:             "cosmopolitan",
			lineageExpectedSrc:      "Players Guide, pg 107",
			lineageExpectedSize:     "Small", // Or "Small" based on specific test cases
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["human"].Traits),
			lineageSelectedTraits:   map[string]string{},
		},
		{
			name:                    "Tik",
			lineageKey:              "kobold",
			heritageKey:             "salvager",
			lineageExpectedSrc:      "Players Guide, pg 108",
			lineageExpectedSize:     "Small",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["kobold"].Traits),
			lineageSelectedTraits:   map[string]string{"Natural Adaptation": "Fierce (Small)"},
		},
		{
			name:                    "Rog",
			lineageKey:              "orc",
			heritageKey:             "diaspora",
			lineageExpectedSrc:      "Players Guide, pg 108",
			lineageExpectedSize:     "Medium",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["orc"].Traits),
			lineageSelectedTraits:   map[string]string{},
		},
		{
			name:                    "Alien",
			lineageKey:              "syderean",
			heritageKey:             "anointed",
			lineageExpectedSrc:      "Players Guide, pg 109",
			lineageExpectedSize:     "Medium",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["syderean"].Traits),
			lineageSelectedTraits:   map[string]string{"Natural Adaptation": "Celestial"},
		},
		{
			name:                    "Frodo",
			lineageKey:              "smallfolk",
			heritageKey:             "cottage",
			lineageExpectedSrc:      "Players Guide, pg 109",
			lineageExpectedSize:     "Small",
			lineagePredefinedTraits: helpers.GetMapKeys(PredefinedTraitsData["smallfolk"].Traits),
			lineageSelectedTraits:   map[string]string{"Natural Adaptation": "Halfling"},
		},
	}

	for _, testCase := range testTraits {
		//character, err := NewCharacter(testCase.Name, testCase.lineageKey, testCase.lineageExpectedSize, testCase.lineageSelectedTraits)

		ctxRef := fmt.Sprintf("Character Lineage test: %s", testCase.lineageKey)
		// Create a character to test against

		character, err := NewCharacter(
			testCase.name, 5, "Wizard",
			"battle mage",
			testCase.lineageKey, testCase.heritageKey,
			Lineages[testCase.lineageKey].SizeOptions[0], rollingOption,
			testCase.lineageSelectedTraits, []string{},
			"Standard", ctxRef, observedLoggerSugared)

		if err != nil {
			t.Fatalf("Error creating character: %v", err)
		}

		// Check character's Name
		if character.Name != testCase.name {
			t.Errorf("Expected Name to be '%s', but got '%s'", testCase.name, character.Name)
		}

		// Check Lineage
		if character.Lineage.Name != helpers.ToTitleCase(testCase.lineageKey) {
			t.Errorf("Expected Lineage Name to be '%s', but got '%s' for %s",
				helpers.ToTitleCase(testCase.lineageKey), character.Lineage.Name,
				testCase.name)
		}

		// Check Lineage source
		if character.Lineage.LineageSource != testCase.lineageExpectedSrc {
			t.Errorf("Expected Lineage source to be '%s', but got '%s' for %s",
				testCase.lineageExpectedSrc, character.Lineage.LineageSource,
				testCase.name)
		}

		// Check chosen size
		if character.CharacterSize != testCase.lineageExpectedSize {
			t.Errorf("Expected chosen size to be '%s', but got '%s' for %s",
				testCase.lineageExpectedSize, character.CharacterSize,
				testCase.name)
		}

		// Check predefined Traits
		if character.Lineage.Traits != nil {
			for _, expectedValues := range testCase.lineagePredefinedTraits {
				assert.Equal(t, true, helpers.Contains(character.Lineage.Traits, expectedValues))
			}
		}

		// Check chosen Traits
		for traitKey, expectedValue := range testCase.lineageSelectedTraits {
			actualValue, exists := character.Traits[traitKey]
			fmt.Printf("Trait: %s, Expected: %s, Actual: %s\n", traitKey, expectedValue, actualValue)
			if !exists {
				t.Errorf("Expected chosen trait '%s' not found in character's Traits for %s",
					traitKey, testCase.name)
			} else if actualValue != expectedValue {
				t.Errorf("Expected chosen trait '%s' to be '%s', but got '%s' for %s",
					traitKey, expectedValue, actualValue, testCase.name)
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
		"nomadic", "Medium", rollingOption, map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)
	assert.NoError(t, err, "Unexpected error when creating test character")
	testCharacter.SkillProficiencies = map[string]AbilitySkillProficiency{
		"athletics": {Skill: "athletics", Source: "Training"},
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
		"human", "nomadic", "Medium", rollingOption, map[string]string{}, []string{},
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
		ctxRef := fmt.Sprintf("Character invalid Lineage test: %s", tc.lineageKey)

		_, err := NewCharacter(
			"Test Wizard", 5, "Wizard",
			"battle mage",
			tc.lineageKey, "nomadic",
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
	ctxRef := "Character no Traits test"

	//character, err := createCharacter("Lineless", "elf", "Medium", nil)
	// Create a test character
	character, err := NewCharacter(
		"Mr NoTraits", 1, "ranger", "pack master",
		"human", "nomadic",
		"Medium", rollingOption,
		map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)

	assert.NoError(t, err, "Unexpected error when creating character with no Traits")
	assert.Equal(t, 0, len(character.Traits), "Expected no chosen Traits")
}

func TestCharacterWithEdgeCaseNames(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

	edgeCaseNames := []struct {
		name        string
		lineageKey  string
		heritageKey string
	}{
		{"", "human", "nomadic"},
		{"@InvalidName!", "dwarf", "fireforge"},
		{"1234", "elf", "grove"},
	}

	for _, tc := range edgeCaseNames {
		ctxRef := fmt.Sprintf("Character Name edge case test: %s", tc.name)

		_, err := NewCharacter(
			tc.name, 1, "Wizard", "battle mage",
			tc.lineageKey, tc.heritageKey,
			Lineages[tc.lineageKey].SizeOptions[0], rollingOption,
			map[string]string{}, []string{},
			"Standard", ctxRef, observedLoggerSugared)

		assert.Error(t, err, fmt.Sprintf("Expected error when creating %s character with invalid Name", tc.lineageKey))
	}
}

func TestCharacterWithEdgeCaseSizes(t *testing.T) {
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"

	edgeCaseSizes := []struct {
		name        string
		lineageKey  string
		heritageKey string
		size        string
	}{
		{"bob", "human", "nomadic", "Tiny"},
		{"sally", "human", "nomadic", "Huge"},
	}
	for _, tc := range edgeCaseSizes {
		_, err := NewCharacter(
			tc.name, 1, "Wizard",
			"battle mage",
			tc.lineageKey, tc.heritageKey,
			tc.size, rollingOption,
			map[string]string{}, []string{},
			"Standard", "Character Size Edge Case", observedLoggerSugared)
		if err == nil {
			t.Errorf("Character creation should have failed for Name '%s' tested size: %s", tc.name, tc.size)
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
		{"Negative OverallLevel", "TestInvalid", -1, "Wizard", "Medium", "standard", true},
		{"Excessive OverallLevel", "TestInvalid", 101, "Wizard", "Medium", "standard", true},
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
				"human", "nomadic", tc.size, tc.rollingOption, map[string]string{}, []string{},
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
		{"OverallLevel One Wizard", "wizard", "battle mage", 1},
		{"OverallLevel Five Fighter", "fighter", "weapon master", 5},
		{"OverallLevel Two Cleric", "cleric", "life domain", 2},
		{"OverallLevel Three Rogue", "rogue", "enforcer", 3},
		{"OverallLevel Ten Barbarian", "barbarian", "berserker", 10},
		{"OverallLevel Twenty Ranger", "ranger", "hunter", 20},
	}

	for _, tc := range hitPointTests {
		t.Run(tc.name, func(t *testing.T) {
			ctxRef := fmt.Sprintf("Hit Point Generation Test: %s", tc.name)

			character, err := NewCharacter(
				tc.name, tc.level, tc.class, tc.subClass,
				"human", "nomadic", "Medium", "standard", map[string]string{}, []string{},
				"Standard", ctxRef, observedLoggerSugared)

			if err != nil {
				t.Fatalf("Error creating character %s: %v", tc.name, err)
			}

			assert.NotNil(t, character.CurrentHitPointsAudit, "CurrentHitPointsAudit should not be nil")
			tmpTotal := 0
			tmpRolls := 0
			for _, v := range character.CurrentHitPointsAudit {
				if roll, ok := v.NewValue.(dice.Roll); ok {
					// fmt.Printf("Rolls: %v, Result: %d Ctx: %s\n", len(v.RollsUsed), v.Result, v.CtxRef)
					tmpTotal += roll.Result
					tmpRolls += len(roll.RollsUsed)
				} else {
					fmt.Printf("Warning: NewValue is not a Roll struct: %T\n", v.NewValue)
				}

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
		"human", "cosmopolitan", "Small", "standard", map[string]string{}, []string{},
		"Standard", ctxRef, observedLoggerSugared)
	assert.NoError(t, err, "Unexpected error when creating character")

	// Add temporary HP
	tempHP := 10
	character.ModifyTemporaryHitPoints(tempHP)

	// Verify temporary HP
	assert.Equal(t, tempHP, character.TemporaryHitPoints, "Temporary HP not added correctly")
	assert.Equal(t, tempHP+character.MaxHitPoints, character.GetTotalHitPoints(), "Total HP not correct")

	damage := 10
	character.Damage(damage, "bludgeoning")
	// fmt.Printf("Starting HP: %d\n", character.DamageAudits[0].HitPointsBefore)
	// fmt.Printf("Temp HP used: %d\n", character.DamageAudits[0].Adjustments["temporary hit points"])
	// fmt.Printf("Base Amount: %d\n", character.DamageAudits[0].BaseAmount)
	// fmt.Printf("Ending HP: %d\n", character.DamageAudits[0].HitPointsAfter)
	// fmt.Printf("character HP: %d\n", character.CurrentHitPoints)
	// Verify remaining HP and temporary HP
	assert.Equal(t, 0, character.TemporaryHitPoints, "Temporary HP not reduced correctly")
	assert.Equal(t, character.DamageAudits[0].HitPointsBefore-10, character.GetTotalHitPoints(), "Regular HP should remain unchanged after temporary HP absorbs damage")
}
