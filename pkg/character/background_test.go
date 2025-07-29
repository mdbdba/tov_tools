package character

import (
	"strings"
	"testing"
)

func TestGetBackgroundByName(t *testing.T) {
	tests := []struct {
		name        string
		bgName      string
		expectError bool
	}{
		{"Valid background - adherent", "adherent", false},
		{"Valid background - artist", "artist", false},
		{"Valid background - courtier", "courtier", false},
		{"Valid background - criminal", "criminal", false},
		{"Valid background - homesteader", "homesteader", false},
		{"Valid background - maker", "maker", false},
		{"Valid background - outcast", "outcast", false},
		{"Valid background - rustic", "rustic", false},
		{"Valid background - scholar", "scholar", false},
		{"Valid background - soldier", "soldier", false},
		{"Valid background - case insensitive", "ADHERENT", false},
		{"Valid background - mixed case", "ArTiSt", false},
		{"Invalid background", "nonexistent", true},
		{"Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bg, err := GetBackgroundByName(tt.bgName)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none for background name: %s", tt.bgName)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for background name %s: %v", tt.bgName, err)
				}
				if bg.Name == "" {
					t.Errorf("Expected background name to be set for: %s", tt.bgName)
				}
			}
		})
	}
}

func TestBackgroundsMapCompleteness(t *testing.T) {
	expectedBackgrounds := []string{
		"adherent", "artist", "courtier", "criminal", "homesteader",
		"maker", "outcast", "rustic", "scholar", "soldier",
	}

	for _, expectedBg := range expectedBackgrounds {
		t.Run("Background exists: "+expectedBg, func(t *testing.T) {
			bg, exists := Backgrounds[expectedBg]
			if !exists {
				t.Errorf("Expected background %s to exist in Backgrounds map", expectedBg)
			}
			if bg.Name == "" {
				t.Errorf("Background %s has empty Name field", expectedBg)
			}
		})
	}

	// Check that we have exactly the expected number of backgrounds
	if len(Backgrounds) != len(expectedBackgrounds) {
		t.Errorf("Expected %d backgrounds, but found %d", len(expectedBackgrounds), len(Backgrounds))
	}
}

func TestBackgroundFieldValidation(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Validate "+bgKey, func(t *testing.T) {
			// Test required fields
			if bg.Name == "" {
				t.Errorf("Background %s has empty Name", bgKey)
			}
			if bg.Description == "" {
				t.Errorf("Background %s has empty Description", bgKey)
			}
			if bg.BackgroundSource == "" {
				t.Errorf("Background %s has empty BackgroundSource", bgKey)
			}

			// Validate BackgroundSource format
			expectedSourcePrefix := "Players Guide, pg "
			if !strings.HasPrefix(bg.BackgroundSource, expectedSourcePrefix) {
				t.Errorf("Background %s has invalid BackgroundSource format: %s", bgKey, bg.BackgroundSource)
			}

			// Test that key matches lowercase name
			expectedKey := strings.ToLower(bg.Name)
			if bgKey != expectedKey {
				t.Errorf("Background key %s doesn't match lowercase name %s", bgKey, expectedKey)
			}
		})
	}
}

func TestBackgroundSkillProficiencies(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Skills "+bgKey, func(t *testing.T) {
			// Test that either fixed proficiencies or options exist (but data shows most use options)
			hasFixedSkills := len(bg.SkillProficiencies) > 0
			hasSkillOptions := len(bg.SkillProficiencyOptions) > 0

			if !hasFixedSkills && !hasSkillOptions {
				t.Errorf("Background %s has no skill proficiencies or options", bgKey)
			}

			// If skill options exist, validate structure
			for optionName, choices := range bg.SkillProficiencyOptions {
				if choices.NumberToSelect <= 0 {
					t.Errorf("Background %s skill option %s has invalid NumberToSelect: %d",
						bgKey, optionName, choices.NumberToSelect)
				}
				if len(choices.Options) == 0 {
					t.Errorf("Background %s skill option %s has no options", bgKey, optionName)
				}
				if choices.NumberToSelect > len(choices.Options) {
					t.Errorf("Background %s skill option %s wants to select %d from %d options",
						bgKey, optionName, choices.NumberToSelect, len(choices.Options))
				}
			}
		})
	}
}

func TestBackgroundTalentOptions(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Talents "+bgKey, func(t *testing.T) {
			if len(bg.TalentOptions) == 0 {
				t.Errorf("Background %s has no talent options", bgKey)
			}

			for optionName, choices := range bg.TalentOptions {
				if choices.NumberToSelect <= 0 {
					t.Errorf("Background %s talent option %s has invalid NumberToSelect: %d",
						bgKey, optionName, choices.NumberToSelect)
				}
				if len(choices.Options) == 0 {
					t.Errorf("Background %s talent option %s has no options", bgKey, optionName)
				}
				if choices.NumberToSelect > len(choices.Options) {
					t.Errorf("Background %s talent option %s wants to select %d from %d options",
						bgKey, optionName, choices.NumberToSelect, len(choices.Options))
				}

				// Validate that talent names are lowercase
				for _, talent := range choices.Options {
					if talent != strings.ToLower(talent) {
						t.Errorf("Background %s has talent option with incorrect case: %s", bgKey, talent)
					}
				}
			}
		})
	}
}

func TestBackgroundEquipment(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Equipment "+bgKey, func(t *testing.T) {
			if len(bg.Equipment) == 0 {
				t.Errorf("Background %s has no equipment", bgKey)
			}

			for _, item := range bg.Equipment {
				if item.Name == "" {
					t.Errorf("Background %s has equipment item with empty name", bgKey)
				}
				if item.Quantity <= 0 {
					t.Errorf("Background %s has equipment item %s with invalid quantity: %d",
						bgKey, item.Name, item.Quantity)
				}
			}
		})
	}
}

func TestBackgroundMoney(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Money "+bgKey, func(t *testing.T) {
			totalValue := bg.Money.GoldPieces + (bg.Money.SilverPieces / 10) + (bg.Money.CopperPieces / 100)
			if totalValue <= 0 {
				t.Errorf("Background %s has no starting money", bgKey)
			}

			// Validate money values are non-negative
			if bg.Money.GoldPieces < 0 {
				t.Errorf("Background %s has negative gold pieces: %d", bgKey, bg.Money.GoldPieces)
			}
			if bg.Money.SilverPieces < 0 {
				t.Errorf("Background %s has negative silver pieces: %d", bgKey, bg.Money.SilverPieces)
			}
			if bg.Money.CopperPieces < 0 {
				t.Errorf("Background %s has negative copper pieces: %d", bgKey, bg.Money.CopperPieces)
			}
		})
	}
}

func TestBackgroundMotivations(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Motivations "+bgKey, func(t *testing.T) {
			if len(bg.Motivations) == 0 {
				t.Errorf("Background %s has no motivations", bgKey)
			}

			// All backgrounds should have adventuring motivations
			adventuringMotivations, hasAdventuring := bg.Motivations["adventuring"]
			if !hasAdventuring {
				t.Errorf("Background %s missing adventuring motivations", bgKey)
			} else {
				if len(adventuringMotivations) == 0 {
					t.Errorf("Background %s has empty adventuring motivations", bgKey)
				}

				// Validate motivation numbering (should start at 1)
				for rollId, motivation := range adventuringMotivations {
					if rollId <= 0 {
						t.Errorf("Background %s has invalid motivation roll ID: %d", bgKey, rollId)
					}
					if motivation == "" {
						t.Errorf("Background %s has empty motivation text for roll %d", bgKey, rollId)
					}
				}
			}

			// Validate other motivation types
			for motivationType, motivations := range bg.Motivations {
				if len(motivations) == 0 {
					t.Errorf("Background %s has empty %s motivations", bgKey, motivationType)
				}
			}
		})
	}
}

func TestBackgroundAdditionalProficiencyOptions(t *testing.T) {
	for bgKey, bg := range Backgrounds {
		t.Run("Additional Proficiencies "+bgKey, func(t *testing.T) {
			for optionName, choices := range bg.AdditionalProficiencyOptions {
				if choices.NumberToSelect <= 0 {
					t.Errorf("Background %s additional proficiency option %s has invalid NumberToSelect: %d",
						bgKey, optionName, choices.NumberToSelect)
				}
				if len(choices.Options) == 0 {
					t.Errorf("Background %s additional proficiency option %s has no options", bgKey, optionName)
				}
				if choices.NumberToSelect > len(choices.Options) {
					t.Errorf("Background %s additional proficiency option %s wants to select %d from %d options",
						bgKey, optionName, choices.NumberToSelect, len(choices.Options))
				}
			}
		})
	}
}

func TestSpecificBackgroundFeatures(t *testing.T) {
	// Test specific background features that should exist

	t.Run("Adherent specific features", func(t *testing.T) {
		adherent := Backgrounds["adherent"]

		// Should have artist tools as fixed proficiency
		found := false
		for _, prof := range adherent.AdditionalProficiencies {
			if prof == "artist tools" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Adherent should have 'artist tools' as additional proficiency")
		}
	})

	t.Run("Criminal specific features", func(t *testing.T) {
		criminal := Backgrounds["criminal"]

		// Should have thieves cant as fixed proficiency
		found := false
		for _, prof := range criminal.AdditionalProficiencies {
			if prof == "thieves cant" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Criminal should have 'thieves cant' as additional proficiency")
		}

		// Should have criminal secrets
		if _, hasSecrets := criminal.Motivations["secret"]; !hasSecrets {
			t.Error("Criminal should have 'secret' motivations")
		}
	})

	t.Run("Homesteader specific features", func(t *testing.T) {
		homesteader := Backgrounds["homesteader"]

		// Should have Survival as fixed skill proficiency
		found := false
		for _, skill := range homesteader.SkillProficiencies {
			if skill == "Survival" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Homesteader should have 'Survival' as fixed skill proficiency")
		}
	})

	t.Run("Artist specific features", func(t *testing.T) {
		artist := Backgrounds["artist"]

		// Should have artistic expression motivations
		if _, hasArtistic := artist.Motivations["artistic"]; !hasArtistic {
			t.Error("Artist should have 'artistic' motivations")
		}
	})
}

func TestBackgroundPrintDetails(t *testing.T) {
	// Test that PrintDetails doesn't panic
	for bgKey, bg := range Backgrounds {
		t.Run("PrintDetails "+bgKey, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("PrintDetails panicked for background %s: %v", bgKey, r)
				}
			}()

			// Create a copy to avoid modifying the original
			bgCopy := bg
			bgCopy.PrintDetails()
		})
	}
}
