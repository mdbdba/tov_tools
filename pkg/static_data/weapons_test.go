package static_data_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tov_tools/pkg/static_data"
)

func TestWeapons(t *testing.T) {
	// Test that the Weapons map is not empty
	assert.Greater(t, len(static_data.Weapons), 0, "Weapons map should not be empty")

	// Test specific weapons exist
	weaponNames := []string{
		"club", "dagger", "greatclub", "handaxe", "javelin",
		"light_hammer", "mace", "quarterstaff", "sickle", "spear",
		"light_crossbow", "dart", "shortbow", "sling",
		"battleaxe", "flail", "glaive", "greataxe", "greatsword",
		"halberd", "lance", "longsword", "maul", "morningstar",
		"pike", "rapier", "scimitar", "scythe", "shortsword",
		"trident", "war_pick", "warhammer", "whip",
		"blowgun", "hand_crossbow", "heavy_crossbow", "longbow",
	}

	for _, name := range weaponNames {
		_, exists := static_data.Weapons[name]
		assert.True(t, exists, "Weapon '%s' should exist in the Weapons map", name)
	}

	// Test weapon categories
	simpleMelee := []string{"club", "dagger", "greatclub", "handaxe", "javelin", "light_hammer", "mace", "quarterstaff", "sickle", "spear"}
	simpleRanged := []string{"light_crossbow", "dart", "shortbow", "sling"}
	martialMelee := []string{"battleaxe", "flail", "glaive", "greataxe", "greatsword", "halberd", "lance", "longsword", "maul", "morningstar", "pike", "rapier", "scimitar", "scythe", "shortsword", "trident", "war_pick", "warhammer", "whip"}
	martialRanged := []string{"blowgun", "hand_crossbow", "heavy_crossbow", "longbow"}

	for _, name := range simpleMelee {
		assert.Equal(t, "Simple Melee", static_data.Weapons[name].Category, "Weapon '%s' should be in 'Simple Melee' category", name)
	}

	for _, name := range simpleRanged {
		assert.Equal(t, "Simple Ranged", static_data.Weapons[name].Category, "Weapon '%s' should be in 'Simple Ranged' category", name)
	}

	for _, name := range martialMelee {
		assert.Equal(t, "Martial Melee", static_data.Weapons[name].Category, "Weapon '%s' should be in 'Martial Melee' category", name)
	}

	for _, name := range martialRanged {
		assert.Equal(t, "Martial Ranged", static_data.Weapons[name].Category, "Weapon '%s' should be in 'Martial Ranged' category", name)
	}

	// Test versatile weapons have both one-handed and two-handed damage
	versatileWeapons := []string{"quarterstaff", "spear", "battleaxe", "longsword", "trident", "warhammer"}
	for _, name := range versatileWeapons {
		weapon := static_data.Weapons[name]
		_, oneHanded := weapon.Damage["one-handed"]
		_, twoHanded := weapon.Damage["two-handed"]

		assert.True(t, oneHanded, "Versatile weapon '%s' should have one-handed damage", name)
		assert.True(t, twoHanded, "Versatile weapon '%s' should have two-handed damage", name)
		assert.Contains(t, weapon.Properties, "Versatile", "Versatile weapon '%s' should have 'Versatile' property", name)
	}

	// Test ranged weapons have proper range values
	rangedWeapons := map[string]static_data.WeaponRange{
		"dagger":         {Min: 20, Max: 60},
		"handaxe":        {Min: 20, Max: 60},
		"javelin":        {Min: 30, Max: 120},
		"light_hammer":   {Min: 20, Max: 60},
		"spear":          {Min: 20, Max: 60},
		"dart":           {Min: 20, Max: 60},
		"light_crossbow": {Min: 80, Max: 320},
		"shortbow":       {Min: 80, Max: 320},
		"sling":          {Min: 30, Max: 120},
		"blowgun":        {Min: 25, Max: 100},
		"hand_crossbow":  {Min: 30, Max: 120},
		"heavy_crossbow": {Min: 100, Max: 400},
		"longbow":        {Min: 150, Max: 600},
	}

	for name, expectedRange := range rangedWeapons {
		weapon := static_data.Weapons[name]
		var actualRange static_data.WeaponRange

		// Get the range from the appropriate damage type
		if _, ok := weapon.Damage["one-handed"]; ok {
			actualRange = weapon.Damage["one-handed"].Range
		} else if _, ok := weapon.Damage["two-handed"]; ok {
			actualRange = weapon.Damage["two-handed"].Range
		}

		assert.Equal(t, expectedRange.Min, actualRange.Min, "Weapon '%s' should have minimum range of %d", name, expectedRange.Min)
		assert.Equal(t, expectedRange.Max, actualRange.Max, "Weapon '%s' should have maximum range of %d", name, expectedRange.Max)
	}

	// Test specific weapon properties
	t.Run("TestSpecificWeapons", func(t *testing.T) {
		// Test Longsword
		longsword := static_data.Weapons["longsword"]
		assert.Equal(t, "Longsword", longsword.Name)
		assert.Equal(t, "Martial Melee", longsword.Category)
		assert.Equal(t, 15, longsword.CostAmount)
		assert.Equal(t, "gp", longsword.CostCoin)
		assert.Equal(t, 3.0, longsword.Weight)
		assert.ElementsMatch(t, []string{"Disarm", "Hamstring"}, longsword.Options)
		assert.ElementsMatch(t, []string{"Versatile"}, longsword.Properties)
		assert.Equal(t, 1, longsword.Damage["one-handed"].TimesToRoll)
		assert.Equal(t, 8, longsword.Damage["one-handed"].Sides)
		assert.Equal(t, "slashing", longsword.Damage["one-handed"].DamageType)
		assert.Equal(t, 1, longsword.Damage["two-handed"].TimesToRoll)
		assert.Equal(t, 10, longsword.Damage["two-handed"].Sides)
		assert.Equal(t, "slashing", longsword.Damage["two-handed"].DamageType)

		// Test Greatsword (2d6 damage)
		greatsword := static_data.Weapons["greatsword"]
		assert.Equal(t, 2, greatsword.Damage["two-handed"].TimesToRoll)
		assert.Equal(t, 6, greatsword.Damage["two-handed"].Sides)
		assert.Equal(t, "slashing", greatsword.Damage["two-handed"].DamageType)

		// Test Blowgun (1d1 damage - special case)
		blowgun := static_data.Weapons["blowgun"]
		assert.Equal(t, 1, blowgun.Damage["one-handed"].TimesToRoll)
		assert.Equal(t, 1, blowgun.Damage["one-handed"].Sides)
		assert.Equal(t, "piercing", blowgun.Damage["one-handed"].DamageType)
	})

	// Test WeaponOptions map
	t.Run("TestWeaponOptions", func(t *testing.T) {
		assert.Greater(t, len(static_data.WeaponOptions), 0, "WeaponOptions map should not be empty")

		optionNames := []string{"Bash", "Disarm", "Hamstring", "Pinning Shot", "Pull", "Ricochet Shot", "Trip"}
		for _, name := range optionNames {
			_, exists := static_data.WeaponOptions[name]
			assert.True(t, exists, "Option '%s' should exist in the WeaponOptions map", name)
		}

		// Verify all option references in weapons are valid
		for weaponName, weapon := range static_data.Weapons {
			for _, option := range weapon.Options {

				_, exists := static_data.WeaponOptions[option]
				assert.True(t, exists, "Weapon '%s' references option '%s' which should exist in WeaponOptions",
					weaponName, option)
			}
		}
	})
}

func TestWeaponDamageTypes(t *testing.T) {
	damageTypes := map[string]bool{}

	// Extract all damage types
	for _, weapon := range static_data.Weapons {
		for _, damage := range weapon.Damage {
			damageTypes[damage.DamageType] = true
		}
	}

	// Verify we have the expected damage types
	expectedDamageTypes := []string{"slashing", "piercing", "bludgeoning"}
	for _, dt := range expectedDamageTypes {
		assert.True(t, damageTypes[dt], "Damage type '%s' should be used by at least one weapon", dt)
	}

	// Check that we don't have unexpected damage types
	for dt := range damageTypes {
		found := false
		for _, expected := range expectedDamageTypes {
			if dt == expected {
				found = true
				break
			}
		}
		assert.True(t, found, "Unexpected damage type '%s' found", dt)
	}
}

func TestWeaponProperties(t *testing.T) {
	properties := map[string]bool{}

	// Extract all properties
	for _, weapon := range static_data.Weapons {
		for _, prop := range weapon.Properties {
			properties[prop] = true
		}
	}

	// Check for expected properties
	expectedProperties := []string{
		"Light", "Finesse", "Thrown", "Thrown (range 20/60 ft.)",
		"Thrown (range 30/120 ft.)", "Two-handed", "Versatile",
		"Ammunition (range 80/320 ft.)", "Ammunition (range 30/120 ft.)",
		"Ammunition (range 25/100 ft.)", "Ammunition (range 100/400 ft.)",
		"Ammunition (range 150/600 ft.)", "Loading", "Heavy", "Reach", "Special",
	}

	for _, prop := range expectedProperties {
		found := false
		for p := range properties {
			if p == prop || (len(p) >= len(prop) && p[:len(prop)] == prop) {
				found = true
				break
			}
		}
		assert.True(t, found, "Property '%s' should be used by at least one weapon", prop)
	}
}
