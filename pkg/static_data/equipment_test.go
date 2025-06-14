package static_data

import (
	"strings"
	"testing"
)

func TestAdventuringGearStructure(t *testing.T) {
	// Verify the top-level map has all expected categories
	expectedCategories := []string{
		"general",
		"alchemical concoctions",
		"ammunition",
		"containers",
		"herbal concoctions",
		"spellcasting foci",
	}

	for _, category := range expectedCategories {
		if _, exists := AdventuringGear[category]; !exists {
			t.Errorf("Expected category '%s' not found in AdventuringGear", category)
		}
	}

	// Check that there are no unexpected categories
	for category := range AdventuringGear {
		found := false
		for _, expected := range expectedCategories {
			if category == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected category '%s' found in AdventuringGear", category)
		}
	}
}

func TestAdventuringGearItemCount(t *testing.T) {
	// Check that each category has a reasonable number of items
	expectedMinItems := map[string]int{
		"general":                30, // At least 30 general items
		"alchemical concoctions": 3,  // At least 3 alchemical concoctions
		"ammunition":             4,  // At least 4 types of ammunition
		"containers":             7,  // At least 7 containers
		"herbal concoctions":     4,  // At least 4 herbal concoctions
		"spellcasting foci":      4,  // At least 4 spellcasting foci
	}

	for category, minCount := range expectedMinItems {
		items, exists := AdventuringGear[category]
		if !exists {
			t.Errorf("Category '%s' not found in AdventuringGear", category)
			continue
		}

		if len(items) < minCount {
			t.Errorf("Category '%s' has %d items, expected at least %d",
				category, len(items), minCount)
		}
	}
}

func TestSpecificItems(t *testing.T) {
	// Test specific items from each category
	testCases := []struct {
		category   string
		key        string
		name       string
		costAmount int
		costCoin   string
		weightEach float64
		quantity   int
	}{
		{"general", "abacus", "Abacus", 2, "gp", 2, 1},
		{"general", "rope", "Rope (50 feet)", 1, "gp", 10, 1},
		{"general", "spyglass", "Spyglass", 1000, "gp", 1, 1},
		{"alchemical concoctions", "acid (vial)", "Acid (vial)", 25, "gp", 1, 1},
		{"ammunition", "arrows", "Arrows (20)", 1, "gp", 1, 20},
		{"containers", "backpack", "Backpack", 2, "gp", 5, 1},
		{"herbal concoctions", "antitoxin", "Antitoxin (vial)", 50, "gp", 0, 1},
		{"spellcasting foci", "arcane focus", "Arcane Focus", 5, "gp", 2, 1},
	}

	for _, tc := range testCases {
		t.Run(tc.category+"/"+tc.key, func(t *testing.T) {
			items, exists := AdventuringGear[tc.category]
			if !exists {
				t.Fatalf("Category '%s' not found", tc.category)
			}

			item, exists := items[tc.key]
			if !exists {
				t.Fatalf("Item '%s' not found in category '%s'", tc.key, tc.category)
			}

			if item.Name != tc.name {
				t.Errorf("Expected name '%s', got '%s'", tc.name, item.Name)
			}

			if item.CostAmount != tc.costAmount {
				t.Errorf("Expected cost amount %d, got %d", tc.costAmount, item.CostAmount)
			}

			if item.CostCoin != tc.costCoin {
				t.Errorf("Expected cost coin '%s', got '%s'", tc.costCoin, item.CostCoin)
			}

			if item.WeightEach != tc.weightEach {
				t.Errorf("Expected weight %f, got %f", tc.weightEach, item.WeightEach)
			}

			if item.Quantity != tc.quantity {
				t.Errorf("Expected quantity %d, got %d", tc.quantity, item.Quantity)
			}
		})
	}
}

func TestItemConsistency(t *testing.T) {
	// Check that all items have consistent values
	for category, items := range AdventuringGear {
		expectedCategoryCapitalized := strings.Title(category)

		for key, item := range items {
			t.Run(category+"/"+key, func(t *testing.T) {
				// Check that all items have a name
				if item.Name == "" {
					t.Errorf("Item '%s' in category '%s' has empty name", key, category)
				}

				// Check that item category matches its parent category
				if !strings.EqualFold(item.Category, expectedCategoryCapitalized) &&
					!strings.EqualFold(item.Category, category) {
					t.Errorf("Item '%s' has category '%s' but is in map category '%s'",
						key, item.Category, category)
				}

				// Check that all cost values are valid
				if item.CostAmount < 0 {
					t.Errorf("Item '%s' has negative cost amount: %d", key, item.CostAmount)
				}

				// Check for valid coin types
				validCoins := []string{"gp", "sp", "cp", ""}
				validCoin := false
				for _, coin := range validCoins {
					if item.CostCoin == coin {
						validCoin = true
						break
					}
				}
				if !validCoin {
					t.Errorf("Item '%s' has invalid coin type: '%s'", key, item.CostCoin)
				}

				// Check that weight is non-negative
				if item.WeightEach < 0 {
					t.Errorf("Item '%s' has negative weight: %f", key, item.WeightEach)
				}

				// Check that quantity is positive
				if item.Quantity <= 0 {
					t.Errorf("Item '%s' has invalid quantity: %d", key, item.Quantity)
				}
			})
		}
	}
}

func TestFindItem(t *testing.T) {
	// Test helper function to find items (could be added to equipment.go later)
	testCases := []struct {
		searchTerm string
		shouldFind bool
	}{
		{"rope", true},
		{"potion of healing", true},
		{"backpack", true},
		{"TORCH", true}, // Case insensitive
		{"nonexistent item", false},
		{"invalid thing", false},
	}

	for _, tc := range testCases {
		t.Run(tc.searchTerm, func(t *testing.T) {
			found := false
			searchTermLower := strings.ToLower(tc.searchTerm)

			for _, items := range AdventuringGear {
				for key := range items {
					if strings.Contains(key, searchTermLower) {
						found = true
						break
					}
				}
				if found {
					break
				}
			}

			if found != tc.shouldFind {
				if tc.shouldFind {
					t.Errorf("Expected to find item '%s' but didn't", tc.searchTerm)
				} else {
					t.Errorf("Expected not to find item '%s' but did", tc.searchTerm)
				}
			}
		})
	}
}
