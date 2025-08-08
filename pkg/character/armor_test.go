package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

// Character stub for testing prerequisites
type CharacterStub struct {
	StrValue int
}

// Implementation of Character for testing
func (c *CharacterStub) GetAbilityValue(ability string) int {
	if ability == "str" {
		return c.StrValue
	}
	return 10 // default value for other abilities
}

func TestArmor(t *testing.T) {
	// Test that the Armor map is not empty
	assert.Greater(t, len(Armor), 0, "Armor map should not be empty")

	// Test specific armor pieces exist
	armorNames := []string{
		"padded", "leather", "studded_leather", "brigandine",
		"hide", "chain_shirt", "scale_mail", "breastplate", "half_plate",
		"ring_mail", "chain_mail", "splint", "plate", "shield",
	}

	for _, name := range armorNames {
		_, exists := Armor[name]
		assert.True(t, exists, "Armor '%s' should exist in the Armor map", name)
	}

	// Test armor categories
	lightArmor := []string{"padded", "leather", "studded_leather", "brigandine"}
	mediumArmor := []string{"hide", "chain_shirt", "scale_mail", "breastplate", "half_plate"}
	heavyArmor := []string{"ring_mail", "chain_mail", "splint", "plate"}
	shields := []string{"shield"}

	for _, name := range lightArmor {
		assert.Equal(t, "Light", Armor[name].Category, "Armor '%s' should be in 'Light' category", name)
	}

	for _, name := range mediumArmor {
		assert.Equal(t, "Medium", Armor[name].Category, "Armor '%s' should be in 'Medium' category", name)
	}

	for _, name := range heavyArmor {
		assert.Equal(t, "Heavy", Armor[name].Category, "Armor '%s' should be in 'Heavy' category", name)
	}

	for _, name := range shields {
		assert.Equal(t, "Shield", Armor[name].Category, "Armor '%s' should be in 'Shield' category", name)
	}

	// Test armor class calculations
	t.Run("TestArmorClassCalculations", func(t *testing.T) {
		// Light armor: Add full DEX modifier
		for _, name := range lightArmor {
			armorClass := Armor[name].ArmorClass
			assert.True(t, armorClass.AddDexterityModifier, "Light armor '%s' should add DEX modifier", name)
			assert.Equal(t, 0, armorClass.DexterityModifierMax, "Light armor '%s' should not cap DEX modifier", name)
		}

		// Medium armor: Add DEX modifier up to 2
		for _, name := range mediumArmor {
			armorClass := Armor[name].ArmorClass
			assert.True(t, armorClass.AddDexterityModifier, "Medium armor '%s' should add DEX modifier", name)
			assert.Equal(t, 2, armorClass.DexterityModifierMax, "Medium armor '%s' should cap DEX modifier at 2", name)
		}

		// Heavy armor: No DEX modifier
		for _, name := range heavyArmor {
			armorClass := Armor[name].ArmorClass
			assert.False(t, armorClass.AddDexterityModifier, "Heavy armor '%s' should not add DEX modifier", name)
		}

		// Shield: +2 AC, no DEX modifier
		shield := Armor["shield"].ArmorClass
		assert.Equal(t, 2, shield.BaseAC, "Shield should provide +2 AC")
		assert.False(t, shield.AddDexterityModifier, "Shield should not add DEX modifier")
	})

	// Test base AC values
	baseACValues := map[string]int{
		"padded":          11,
		"leather":         11,
		"studded_leather": 12,
		"brigandine":      13,
		"hide":            12,
		"chain_shirt":     13,
		"scale_mail":      14,
		"breastplate":     14,
		"half_plate":      15,
		"ring_mail":       15,
		"chain_mail":      16,
		"splint":          17,
		"plate":           18,
		"shield":          2,
	}

	for name, expectedAC := range baseACValues {
		assert.Equal(t, expectedAC, Armor[name].ArmorClass.BaseAC,
			"Armor '%s' should have base AC of %d", name, expectedAC)
	}

	// Test properties
	t.Run("TestArmorProperties", func(t *testing.T) {
		// Test Noisy property
		noisyArmor := []string{"padded", "brigandine", "scale_mail", "half_plate",
			"ring_mail", "chain_mail", "splint", "plate"}

		for _, name := range noisyArmor {
			armor := Armor[name]
			found := false
			for _, prop := range armor.Properties {
				if prop == "Noisy" || prop == "may be Noisy" {
					found = true
					break
				}
			}
			assert.True(t, found, "Armor '%s' should have Noisy property", name)
		}

		// Test Natural Materials property
		naturalMaterialsArmor := []string{"leather", "hide"}
		for _, name := range naturalMaterialsArmor {
			assert.Contains(t, Armor[name].Properties, "Natural Materials",
				"Armor '%s' should have Natural Materials property", name)
		}

		// Test Cumbersome properties
		cumbersomeArmor := map[string]string{
			"chain_mail": "Cumbersome (STR 13)",
			"splint":     "Cumbersome (STR 15)",
			"plate":      "Cumbersome (STR 16)",
		}

		for name, property := range cumbersomeArmor {
			assert.Contains(t, Armor[name].Properties, property,
				"Armor '%s' should have '%s' property", name, property)
		}
	})

	// Test prerequisites
	// Given
	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "common"
	ctxRef := "Test Armor Prerequisites"
	// Create a test character
	testCharacter, err := NewCharacter(
		"Test Wizard", 5, "Wizard",
		"battle mage", "human",
		"nomadic", "Medium", rollingOption, map[string]string{}, []string{}, []string{},
		"Standard", ClassBuildType{}, ctxRef, observedLoggerSugared)
	testCharacter.SetAbilitySkills()

	assert.NoError(t, err, "Unexpected error when creating test character")

	t.Run("TestArmorPrerequisites", func(t *testing.T) {
		testCharacter.Abilities.Values["str"] = 12
		assert.False(t, Armor["chain_mail"].Prerequisite(testCharacter),
			"Chain mail should require STR 13")
		testCharacter.Abilities.Values["str"] = 13
		assert.True(t, Armor["chain_mail"].Prerequisite(testCharacter),
			"Chain mail should be usable with STR 13")

		testCharacter.Abilities.Values["str"] = 14
		assert.False(t, Armor["splint"].Prerequisite(testCharacter),
			"Splint should require STR 15")
		testCharacter.Abilities.Values["str"] = 15
		assert.True(t, Armor["splint"].Prerequisite(testCharacter),
			"Splint should be usable with STR 15")

		testCharacter.Abilities.Values["str"] = 15
		assert.False(t, Armor["plate"].Prerequisite(testCharacter),
			"Plate should require STR 16")
		testCharacter.Abilities.Values["str"] = 16
		assert.True(t, Armor["plate"].Prerequisite(testCharacter),
			"Plate should be usable with STR 16")

		// Light and medium armor should have no STR prerequisite

		testCharacter.Abilities.Values["str"] = 8
		for _, name := range append(lightArmor, mediumArmor...) {
			assert.True(t, Armor[name].Prerequisite(testCharacter),
				"Armor '%s' should have no STR prerequisite", name)
		}
	})

	// Test specific armor details
	t.Run("TestSpecificArmorDetails", func(t *testing.T) {
		// Test plate armor details
		plate := Armor["plate"]
		assert.Equal(t, "Plate", plate.Name)
		assert.Contains(t, plate.Description, "shaped, interlocking metal plates")
		assert.Equal(t, "Heavy", plate.Category)
		assert.Equal(t, 1500, plate.CostAmount)
		assert.Equal(t, "gp", plate.CostCoin)
		assert.Equal(t, 65.0, plate.Weight)
		assert.Contains(t, plate.Properties, "Cumbersome (STR 16)")
		assert.Contains(t, plate.Properties, "Noisy")
		assert.Equal(t, 18, plate.ArmorClass.BaseAC)
		assert.False(t, plate.ArmorClass.AddDexterityModifier)

		// Test leather armor details
		leather := Armor["leather"]
		assert.Equal(t, "Leather", leather.Name)
		assert.Contains(t, leather.Description, "stiffened by being boiled in oil")
		assert.Equal(t, "Light", leather.Category)
		assert.Equal(t, 10, leather.CostAmount)
		assert.Equal(t, "gp", leather.CostCoin)
		assert.Equal(t, 10.0, leather.Weight)
		assert.Contains(t, leather.Properties, "Natural Materials")
		assert.Equal(t, 11, leather.ArmorClass.BaseAC)
		assert.True(t, leather.ArmorClass.AddDexterityModifier)
		assert.Equal(t, 0, leather.ArmorClass.DexterityModifierMax)

		// Test breastplate details
		breastplate := Armor["breastplate"]
		assert.Equal(t, "Breastplate", breastplate.Name)
		assert.Contains(t, breastplate.Description, "fitted metal chest")
		assert.Equal(t, "Medium", breastplate.Category)
		assert.Equal(t, 400, breastplate.CostAmount)
		assert.Equal(t, "gp", breastplate.CostCoin)
		assert.Equal(t, 20.0, breastplate.Weight)
		assert.Empty(t, breastplate.Properties)
		assert.Equal(t, 14, breastplate.ArmorClass.BaseAC)
		assert.True(t, breastplate.ArmorClass.AddDexterityModifier)
		assert.Equal(t, 2, breastplate.ArmorClass.DexterityModifierMax)
	})
}
