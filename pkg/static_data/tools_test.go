package static_data

import (
	"testing"
)

func TestToolsMapStructure(t *testing.T) {
	// Check if Tools map exists and is not empty
	if len(Tools) == 0 {
		t.Error("Tools map is empty")
	}

	// Check for expected tools
	expectedTools := []string{
		"alchemist tools", "artist tools", "charlatan tools", "clothier tools",
		"construction tools", "gaming set", "herbalist tools", "musical instrument",
		"navigator tools", "provisioner tools", "smithing tools", "thieves' tools",
		"tinker tools", "trapper tools",
	}

	for _, toolName := range expectedTools {
		if _, exists := Tools[toolName]; !exists {
			t.Errorf("Expected tool '%s' not found in Tools map", toolName)
		}
	}
}

func TestToolStructures(t *testing.T) {
	for toolName, tool := range Tools {
		// Check name is not empty
		if tool.Name == "" {
			t.Errorf("Tool '%s' has empty Name field", toolName)
		}

		// Check description is not empty
		if tool.Description == "" {
			t.Errorf("Tool '%s' has empty Description field", toolName)
		}

		// Check associated abilities
		if len(tool.AssociatedAbilities) == 0 {
			t.Errorf("Tool '%s' has no AssociatedAbilities", toolName)
		}

		// Check components
		if len(tool.Components) == 0 {
			t.Errorf("Tool '%s' has no Components", toolName)
		}

		// Check example tasks
		if len(tool.ExampleTasks) == 0 {
			t.Errorf("Tool '%s' has no ExampleTasks", toolName)
		}

		// Validate that example tasks have proper structure
		for i, task := range tool.ExampleTasks {
			if task.Description == "" {
				t.Errorf("Tool '%s' has empty Name in ExampleTask at index %d", toolName, i)
			}
			if task.DC <= 0 {
				t.Errorf("Tool '%s' has invalid DC in ExampleTask at index %d", toolName, i)
			}
		}
	}
}

func TestSpecificTools(t *testing.T) {
	// Test specific tool fields
	t.Run("AlchemistTools", func(t *testing.T) {
		alchemist, exists := Tools["alchemist tools"]
		if !exists {
			t.Fatal("Alchemist tools not found")
		}

		if alchemist.Name != "Alchemist Tools" {
			t.Errorf("Expected 'Alchemist Tools', got '%s'", alchemist.Name)
		}

		if len(alchemist.AssociatedAbilities) != 2 ||
			alchemist.AssociatedAbilities[0] != "dex" ||
			alchemist.AssociatedAbilities[1] != "int" {
			t.Errorf("Incorrect associated abilities for alchemist tools")
		}
	})

	t.Run("ThievesTools", func(t *testing.T) {
		thieves, exists := Tools["thieves' tools"]
		if !exists {
			t.Fatal("Thieves' tools not found")
		}

		// Check special uses
		if thieves.SpecialUses == nil {
			t.Error("Thieves' tools should have special uses")
		} else {
			_, hasPick := thieves.SpecialUses["Pick Locks"]
			_, hasDisarm := thieves.SpecialUses["Disarm Traps"]
			if !hasPick || !hasDisarm {
				t.Error("Thieves' tools missing expected special uses")
			}
		}
	})

	t.Run("ConstructionTools", func(t *testing.T) {
		construction, exists := Tools["construction tools"]
		if !exists {
			t.Fatal("Construction tools not found")
		}

		// Check special uses
		if construction.SpecialUses == nil {
			t.Error("Construction tools should have special uses")
		} else {
			_, hasFortify := construction.SpecialUses["Fortify"]
			_, hasBuildCover := construction.SpecialUses["Build Cover"]
			if !hasFortify || !hasBuildCover {
				t.Error("Construction tools missing expected special uses")
			}
		}
	})
}

func TestExampleTasksDC(t *testing.T) {
	// Verify that each tool has example tasks with valid DCs
	for toolName, tool := range Tools {
		// Most tools follow the pattern of DCs 10, 15, 20
		if len(tool.ExampleTasks) != 3 {
			t.Errorf("Tool '%s' should have exactly 3 example tasks", toolName)
			continue
		}

		expectedDCs := []int{10, 15, 20}
		for i, dc := range expectedDCs {
			if i >= len(tool.ExampleTasks) {
				t.Errorf("Tool '%s' missing example task at index %d", toolName, i)
				continue
			}

			if tool.ExampleTasks[i].DC != dc {
				t.Errorf("Tool '%s' task %d should have DC %d, got %d",
					toolName, i, dc, tool.ExampleTasks[i].DC)
			}
		}
	}
}

func TestCharlatanToolsSpecialUses(t *testing.T) {
	charlatan, exists := Tools["charlatan tools"]
	if !exists {
		t.Fatal("Charlatan tools not found")
	}

	// Check special uses
	if charlatan.SpecialUses == nil {
		t.Error("Charlatan tools should have special uses")
	} else {
		disguise, hasDisguise := charlatan.SpecialUses["Create Disguise"]
		forge, hasForge := charlatan.SpecialUses["Forge Documents"]

		if !hasDisguise {
			t.Error("Charlatan tools missing 'Create Disguise' special use")
		} else if len(disguise) < 50 {
			t.Error("Create Disguise description too short, might be incomplete")
		}

		if !hasForge {
			t.Error("Charlatan tools missing 'Forge Documents' special use")
		} else if len(forge) < 50 {
			t.Error("Forge Documents description too short, might be incomplete")
		}
	}

	// Check example tasks (they were incorrect in the original file)
	expectedTasks := []struct {
		Name string
		DC   int
	}{
		{"Hide a noteworthy physical feature or minor injury", 10},
		{"Forge a signature from memory", 15},
		{"Make yourself look like a well-known celebrity", 20},
	}

	for i, expected := range expectedTasks {
		if i >= len(charlatan.ExampleTasks) {
			t.Errorf("Charlatan tools missing example task %d", i)
			continue
		}

		task := charlatan.ExampleTasks[i]
		if task.Description != expected.Name {
			t.Errorf("Charlatan tools task %d expected name '%s', got '%s'",
				i, expected.Name, task.Description)
		}

		if task.DC != expected.DC {
			t.Errorf("Charlatan tools task %d expected DC %d, got %d",
				i, expected.DC, task.DC)
		}
	}
}
