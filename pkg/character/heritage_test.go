package character

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetHeritageByName(t *testing.T) {
	tests := []struct {
		name           string
		heritageName   string
		expectedExists bool
		expectedName   string
	}{
		{
			name:           "Existing heritage with exact case",
			heritageName:   "Anointed",
			expectedExists: true,
			expectedName:   "Anointed",
		},
		{
			name:           "Existing heritage with lowercase",
			heritageName:   "anointed",
			expectedExists: true,
			expectedName:   "Anointed",
		},
		{
			name:           "Existing heritage with mixed case",
			heritageName:   "AnOiNtEd",
			expectedExists: true,
			expectedName:   "Anointed",
		},
		{
			name:           "Non-existing heritage",
			heritageName:   "NonExistentHeritage",
			expectedExists: false,
			expectedName:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heritage, err := GetHeritageByName(tt.heritageName)

			if tt.expectedExists {
				if err != nil {
					t.Errorf("Expected heritage '%s' to exist, but got error: %v", tt.heritageName, err)
				}
				if heritage.Name != tt.expectedName {
					t.Errorf("Expected heritage name '%s', but got '%s'", tt.expectedName, heritage.Name)
				}
			} else {
				if err == nil {
					t.Errorf("Expected heritage '%s' to not exist, but it does", tt.heritageName)
				}
			}
		})
	}
}

func TestHeritageProperties(t *testing.T) {
	// Test a few select heritages to verify their properties are correctly set
	tests := []struct {
		heritageName         string
		expectedSource       string
		expectedDefaultLang  []string
		expectedTraitCount   int
		expectedTraitOptions int
	}{
		{
			heritageName:         "anointed",
			expectedSource:       "Players Guide, pg 112",
			expectedDefaultLang:  []string{"Common"},
			expectedTraitCount:   2,
			expectedTraitOptions: 2,
		},
		{
			heritageName:         "cloud",
			expectedSource:       "Players Guide, pg 113",
			expectedDefaultLang:  []string{"Common"},
			expectedTraitCount:   1,
			expectedTraitOptions: 3,
		},
		{
			heritageName:         "wildlands",
			expectedSource:       "Players Guide, pg 116",
			expectedDefaultLang:  []string{"Common"},
			expectedTraitCount:   3,
			expectedTraitOptions: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.heritageName, func(t *testing.T) {
			heritage, err := GetHeritageByName(tt.heritageName)
			if err != nil {
				t.Fatalf("Failed to get heritage: %v", err)
			}

			if heritage.HeritageSource != tt.expectedSource {
				t.Errorf("Expected source '%s', got '%s'", tt.expectedSource, heritage.HeritageSource)
			}

			if !reflect.DeepEqual(heritage.LanguageDefaults, tt.expectedDefaultLang) {
				t.Errorf("Expected default languages %v, got %v", tt.expectedDefaultLang, heritage.LanguageDefaults)
			}

			if len(heritage.Traits) != tt.expectedTraitCount {
				t.Errorf("Expected %d traits, got %d", tt.expectedTraitCount, len(heritage.Traits))
			}

			if len(heritage.TraitOptions) != tt.expectedTraitOptions {
				t.Errorf("Expected %d trait options, got %d", tt.expectedTraitOptions, len(heritage.TraitOptions))
			}
		})
	}
}

func TestHeritageSuggestionMap(t *testing.T) {
	// Test that heritage suggestions for lineages are valid
	suggestions := HeritageSuggestion()

	// Check a few known lineages and their suggested heritages
	tests := []struct {
		lineage           string
		expectedHeritages []string
	}{
		{
			lineage:           "Dwarf",
			expectedHeritages: []string{"Fireforge", "Stone"},
		},
		{
			lineage:           "Elf",
			expectedHeritages: []string{"Cloud", "Grove"},
		},
		{
			lineage:           "Human",
			expectedHeritages: []string{"Cosmopolitan", "Nomadic"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.lineage, func(t *testing.T) {
			suggestedHeritages, exists := suggestions[tt.lineage]
			if !exists {
				t.Errorf("Expected lineage '%s' to have heritage suggestions", tt.lineage)
				return
			}

			if !reflect.DeepEqual(suggestedHeritages, tt.expectedHeritages) {
				t.Errorf("Expected heritages %v for lineage '%s', got %v",
					tt.expectedHeritages, tt.lineage, suggestedHeritages)
			}

			// Verify all suggested heritages actually exist
			for _, heritageName := range suggestedHeritages {
				_, err := GetHeritageByName(heritageName)
				if err != nil {
					t.Errorf("Suggested heritage '%s' for lineage '%s' does not exist",
						heritageName, tt.lineage)
				}
			}
		})
	}
}

func TestHeritageLanguageOptions(t *testing.T) {
	// Test that language options exclude default languages
	heritageNames := []string{"anointed", "cloud", "cosmopolitan", "cottage", "wildlands"}

	for _, name := range heritageNames {
		t.Run(name, func(t *testing.T) {
			heritage, err := GetHeritageByName(name)
			if err != nil {
				t.Fatalf("Failed to get heritage: %v", err)
			}

			// Check if "Languages" is in the trait options
			langOptions, exists := heritage.TraitOptions["Languages"]
			if !exists {
				t.Fatal("Expected 'Languages' in trait options")
			}

			// Ensure Common is not in the options (since it's a default)
			for _, option := range langOptions.Options {
				if option == "Common" {
					t.Error("Common should be excluded from language options since it's a default")
				}
			}

			// Verify NumberToSelect is positive
			if langOptions.NumberToSelect <= 0 {
				t.Errorf("Expected NumberToSelect to be positive, got %d", langOptions.NumberToSelect)
			}
		})
	}
}

func TestAllHeritagesHaveValidFields(t *testing.T) {
	// Test that all heritages have valid and non-empty required fields
	for heritageKey, heritage := range Heritages {
		t.Run(heritageKey, func(t *testing.T) {
			// Name should match the capitalized version of the key
			expectedName := strings.ToUpper(heritageKey[:1]) + heritageKey[1:]
			if heritage.Name != expectedName {
				t.Errorf("Expected name '%s', got '%s'", expectedName, heritage.Name)
			}

			// Source should not be empty
			if heritage.HeritageSource == "" {
				t.Error("Heritage source should not be empty")
			}

			// Should have at least Common in default languages
			if len(heritage.LanguageDefaults) == 0 || heritage.LanguageDefaults[0] != "Common" {
				t.Error("Heritage should have Common as a default language")
			}

			// Should have at least one trait
			if len(heritage.Traits) == 0 {
				t.Error("Heritage should have at least one trait")
			}

			// Language options should always be present
			if _, exists := heritage.TraitOptions["Languages"]; !exists {
				t.Error("Heritage should have Language options in TraitOptions")
			}
		})
	}
}

func TestHeritageTraitContent(t *testing.T) {
	// Test specific traits for selected heritages
	tests := []struct {
		heritageName string
		traitName    string
		expectedText string
	}{
		{
			heritageName: "anointed",
			traitName:    "Favored Disciple",
			expectedText: "You know the thaumaturgy cantrip and you have advantage on death saves.",
		},
		{
			heritageName: "fireforge",
			traitName:    "Heat Resilience",
			expectedText: "Lifelong exposure has made you resilient to the effects of severe heat. You are resistant to fire damage.",
		},
		{
			heritageName: "grove",
			traitName:    "Canopy Walker",
			expectedText: "You have a climbing speed equal to your walking speed.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.heritageName+"_"+tt.traitName, func(t *testing.T) {
			heritage, err := GetHeritageByName(tt.heritageName)
			if err != nil {
				t.Fatalf("Failed to get heritage: %v", err)
			}

			traitText, exists := heritage.Traits[tt.traitName]
			if !exists {
				t.Fatalf("Expected trait '%s' to exist", tt.traitName)
			}

			if traitText != tt.expectedText {
				t.Errorf("Expected trait text '%s', got '%s'", tt.expectedText, traitText)
			}
		})
	}
}

func TestHeritageCaseInsensitivity(t *testing.T) {
	// Test that heritage retrieval is case-insensitive for all heritages
	for heritageKey := range Heritages {
		// Create mixed-case version
		mixedCase := ""
		for i, char := range heritageKey {
			if i%2 == 0 {
				mixedCase += strings.ToUpper(string(char))
			} else {
				mixedCase += strings.ToLower(string(char))
			}
		}

		t.Run(mixedCase, func(t *testing.T) {
			heritage, err := GetHeritageByName(mixedCase)
			if err != nil {
				t.Errorf("Failed to get heritage with mixed case '%s': %v", mixedCase, err)
			}

			// Check that the name is properly capitalized in the result
			expectedName := strings.ToUpper(heritageKey[:1]) + heritageKey[1:]
			if heritage.Name != expectedName {
				t.Errorf("Expected name '%s', got '%s'", expectedName, heritage.Name)
			}
		})
	}
}
