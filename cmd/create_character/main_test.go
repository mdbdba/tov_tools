package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"tov_tools/pkg/character"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestJSONTraitsParsing tests the JSON parsing logic for traits
func TestJSONTraitsParsing(t *testing.T) {
	tests := []struct {
		name        string
		traitsJSON  string
		expected    map[string]string
		expectError bool
	}{
		{
			name:       "Valid JSON with traits",
			traitsJSON: `{"Natural Adaptation": "Agile", "Animal Instinct": "Perception"}`,
			expected: map[string]string{
				"Natural Adaptation": "Agile",
				"Animal Instinct":    "Perception",
			},
			expectError: false,
		},
		{
			name:        "Empty JSON object",
			traitsJSON:  `{}`,
			expected:    map[string]string{},
			expectError: false,
		},
		{
			name:        "Empty string",
			traitsJSON:  "",
			expected:    map[string]string{},
			expectError: false,
		},
		{
			name:        "Invalid JSON",
			traitsJSON:  `{"invalid": json}`,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "JSON with wrong value type",
			traitsJSON:  `{"trait": 123}`,
			expected:    nil,
			expectError: true,
		},
		{
			name:       "Complex trait values",
			traitsJSON: `{"Natural Adaptation": "Fierce (Small)", "Natural Weapons": "Claws"}`,
			expected: map[string]string{
				"Natural Adaptation": "Fierce (Small)",
				"Natural Weapons":    "Claws",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var selectedTraits map[string]string
			var err error

			if tt.traitsJSON != "" {
				err = json.Unmarshal([]byte(tt.traitsJSON), &selectedTraits)
			} else {
				selectedTraits = make(map[string]string)
			}

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, selectedTraits)
		})
	}
}

// TestCLIArgumentParsing tests the CLI argument parsing using a helper function
func TestCLIArgumentParsing(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected CLIArgs
	}{
		{
			name: "All arguments provided",
			args: []string{
				"-name=Fang",
				"-level=1",
				"-class=barbarian",
				"-subclass=berserker",
				"-lineage=beastkin",
				"-heritage=slayer",
				`-traits={"Natural Adaptation": "Agile"}`,
			},
			expected: CLIArgs{
				name:         "Fang",
				level:        1,
				class:        "barbarian",
				subclass:     "berserker",
				lineage:      "beastkin",
				heritage:     "slayer",
				traitsJSON:   `{"Natural Adaptation": "Agile"}`,
				parsedTraits: map[string]string{"Natural Adaptation": "Agile"},
			},
		},
		{
			name: "Minimal required arguments",
			args: []string{
				"-name=TestChar",
				"-class=fighter",
				"-lineage=human",
				"-heritage=nomadic",
			},
			expected: CLIArgs{
				name:         "TestChar",
				level:        1, // default
				class:        "fighter",
				subclass:     "", // default
				lineage:      "human",
				heritage:     "nomadic",
				traitsJSON:   "", // default
				parsedTraits: map[string]string{},
			},
		},
		{
			name: "With empty traits",
			args: []string{
				"-name=EmptyTraits",
				"-class=wizard",
				"-lineage=elf",
				"-heritage=cloud",
				"-traits={}",
			},
			expected: CLIArgs{
				name:         "EmptyTraits",
				level:        1,
				class:        "wizard",
				subclass:     "",
				lineage:      "elf",
				heritage:     "cloud",
				traitsJSON:   "{}",
				parsedTraits: map[string]string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseArgs(tt.args)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestCLIValidation tests the validation logic
func TestCLIValidation(t *testing.T) {
	tests := []struct {
		name        string
		args        CLIArgs
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid arguments",
			args: CLIArgs{
				name:     "ValidChar",
				lineage:  "human",
				heritage: "nomadic",
			},
			expectError: false,
		},
		{
			name: "Missing lineage",
			args: CLIArgs{
				name:     "NoLineage",
				lineage:  "",
				heritage: "nomadic",
			},
			expectError: true,
			errorMsg:    "lineage name is required",
		},
		{
			name: "Invalid lineage",
			args: CLIArgs{
				name:     "BadLineage",
				lineage:  "invalid_lineage",
				heritage: "nomadic",
			},
			expectError: true,
			errorMsg:    "invalid lineage name: invalid_lineage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateArgs(tt.args)
			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

// TestCLICharacterCreation tests the character creation from CLI args
func TestCLICharacterCreation(t *testing.T) {
	tests := []struct {
		name         string
		args         CLIArgs
		expectError  bool
		validateChar func(*testing.T, *character.Character)
	}{
		{
			name: "Successful character creation",
			args: CLIArgs{
				name:         "TestHero",
				level:        1,
				class:        "fighter",
				subclass:     "weapon master",
				lineage:      "human",
				heritage:     "nomadic",
				parsedTraits: map[string]string{},
			},
			expectError: false,
			validateChar: func(t *testing.T, c *character.Character) {
				assert.Equal(t, "TestHero", c.Name)
				assert.Equal(t, 1, c.OverallLevel)
				assert.Equal(t, "fighter", c.CharacterClassStr)
				assert.Equal(t, "Human", c.Lineage.Name)
			},
		},

		{
			name: "Character with traits",
			args: CLIArgs{
				name:     "TraitChar",
				level:    1,
				class:    "barbarian",
				subclass: "berserker",
				lineage:  "beastkin",
				heritage: "slayer",
				parsedTraits: map[string]string{
					"Natural Adaptation": "Agile",
					"Animal Instinct":    "Perception",
				},
			},
			expectError: false,
			validateChar: func(t *testing.T, c *character.Character) {
				assert.Equal(t, "TraitChar", c.Name)
				assert.Equal(t, "Beastkin", c.Lineage.Name)
				assert.Equal(t, "Agile", c.Traits["Natural Adaptation"])
				assert.Equal(t, "Perception", c.Traits["Animal Instinct"])
			},
		},
		{
			name: "Invalid character parameters",
			args: CLIArgs{
				name:         "", // Invalid empty name
				level:        1,
				class:        "fighter",
				lineage:      "human",
				heritage:     "nomadic",
				parsedTraits: map[string]string{},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			char, err := createCharacterFromArgs(tt.args)

			if tt.expectError {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, char)

			if tt.validateChar != nil {
				tt.validateChar(t, char)
			}
		})
	}
}

// TestCLIIntegration tests the CLI by running it as a subprocess
func TestCLIIntegration(t *testing.T) {
	// Skip if running in CI or if we can't build the binary
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Build the CLI binary
	cmd := exec.Command("go", "build", "-o", "test_cli", ".")
	err := cmd.Run()
	require.NoError(t, err, "Failed to build CLI binary")
	defer os.Remove("test_cli") // Clean up

	tests := []struct {
		user_id        string
		name           string
		args           []string
		expectExitCode int
		expectOutput   []string
		expectError    []string
	}{
		{
			name: "Successful character creation",
			args: []string{
				"-user_id=Skelly",
				"-name=CLITest",
				"-class=fighter",
				"-lineage=human",
				"-heritage=nomadic",
			},
			expectExitCode: 0,
			expectOutput:   []string{"Created character with class:", "fighter", "Character: CLITest"},
		},
		{
			name: "Character with traits",
			args: []string{
				"-user_id=Skelly",
				"-name=TraitTest",
				"-class=barbarian",
				"-lineage=beastkin",
				"-heritage=slayer",
				`-traits={"Natural Adaptation": "Agile"}`,
			},
			expectExitCode: 0,
			expectOutput:   []string{"Created character with class:", "barbarian", "Natural Adaptation: Agile"},
		},
		{
			name:           "Missing required lineage",
			args:           []string{"-name=NoLineage", "-class=fighter"},
			expectExitCode: 2,
			expectOutput:   []string{"lineage name is required"},
		},
		{
			name: "Invalid lineage",
			args: []string{
				"-user_id=Skelly",
				"-name=BadLineage",
				"-class=fighter",
				"-lineage=invalid_lineage",
				"-heritage=nomadic",
			},
			expectExitCode: 2,
			expectOutput:   []string{"invalid lineage name: invalid_lineage"},
		},
		{
			name: "Invalid JSON traits",
			args: []string{
				"-user_id=Skelly",
				"-name=BadJSON",
				"-class=fighter",
				"-lineage=human",
				"-heritage=nomadic",
				"-traits={invalid json}",
			},
			expectExitCode: 2,
			expectOutput:   []string{"error parsing traits JSON"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("./test_cli", tt.args...)
			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			// Check exit code
			if exitError, ok := err.(*exec.ExitError); ok {
				assert.Equal(t, tt.expectExitCode, exitError.ExitCode())
			} else if tt.expectExitCode == 0 {
				assert.NoError(t, err)
			}

			// Check expected output
			output := stdout.String()
			for _, expected := range tt.expectOutput {
				assert.Contains(t, output, expected, "Expected output not found")
			}

			// Check expected errors
			errorOutput := stderr.String()
			for _, expected := range tt.expectError {
				assert.Contains(t, errorOutput, expected, "Expected error not found")
			}
		})
	}
}

// Helper types and functions for testing

type CLIArgs struct {
	user_id      string
	name         string
	level        int
	class        string
	subclass     string
	lineage      string
	heritage     string
	traitsJSON   string
	parsedTraits map[string]string
}

func parseArgs(args []string) (CLIArgs, error) {
	// Create a new flag set for testing
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.SetOutput(bytes.NewBuffer(nil)) // Suppress output during tests
	user_id := fs.String("user_id", "Skelly", "The name of the creator of the character")
	name := fs.String("name", "", "The name of the character to create")
	level := fs.Int("level", 1, "The level of the character to create")
	class := fs.String("class", "", "The class of the character to create")
	subclass := fs.String("subclass", "", "The subclass of the character to create")
	lineage := fs.String("lineage", "", "The lineage of the character to create")
	heritage := fs.String("heritage", "", "The heritage of the character to create")
	traitsJSON := fs.String("traits", "", "The traits of the character to create")

	err := fs.Parse(args)
	if err != nil {
		return CLIArgs{}, err
	}

	// Parse traits JSON
	var parsedTraits map[string]string
	if *traitsJSON != "" {
		err = json.Unmarshal([]byte(*traitsJSON), &parsedTraits)
		if err != nil {
			return CLIArgs{}, fmt.Errorf("error parsing traits JSON: %v", err)
		}
	} else {
		parsedTraits = make(map[string]string)
	}

	return CLIArgs{
		user_id:      *user_id,
		name:         *name,
		level:        *level,
		class:        *class,
		subclass:     *subclass,
		lineage:      *lineage,
		heritage:     *heritage,
		traitsJSON:   *traitsJSON,
		parsedTraits: parsedTraits,
	}, nil
}

func validateArgs(args CLIArgs) error {
	if args.lineage == "" {
		return fmt.Errorf("lineage name is required")
	}
	if !character.ValidateLineage(args.lineage) {
		return fmt.Errorf("invalid lineage name: %s", args.lineage)
	}
	return nil
}

func createCharacterFromArgs(args CLIArgs) (*character.Character, error) {

	observedZapCore, _ := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()

	ctxRef := fmt.Sprintf("cli character generator test: %s", args.name)

	// This would call your actual NewCharacter function
	return character.NewCharacter(args.user_id,
		args.name, args.level, args.class,
		args.subclass, args.lineage, args.heritage,
		character.Lineages[args.lineage].SizeOptions[0], "common",
		args.parsedTraits, []string{}, []string{},
		"Standard", character.ClassBuildType{}, ctxRef, observedLoggerSugared) // Using nil for logger in tests
}

// Benchmark tests for performance
func BenchmarkJSONTraitsParsing(b *testing.B) {
	traitsJSON := `{"Natural Adaptation": "Agile", "Animal Instinct": "Perception", "Natural Weapons": "Claws"}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var selectedTraits map[string]string
		json.Unmarshal([]byte(traitsJSON), &selectedTraits)
	}
}

func BenchmarkCharacterCreation(b *testing.B) {
	args := CLIArgs{
		user_id:  "Skelly",
		name:     "BenchmarkChar",
		level:    1,
		class:    "fighter",
		subclass: "weapon master",
		lineage:  "human",
		heritage: "nomadic",
		parsedTraits: map[string]string{
			"Natural Adaptation": "Agile",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		char, err := createCharacterFromArgs(args)
		if err != nil {
			b.Fatal(err)
		}
		_ = char // Use the character to avoid optimization
	}
}

// Example test showing how to test CLI help output
func TestCLIHelp(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping CLI help test in short mode")
	}

	// Build the CLI binary
	cmd := exec.Command("go", "build", "-o", "test_cli_help", ".")
	err := cmd.Run()
	require.NoError(t, err)
	defer os.Remove("test_cli_help")

	// Test help flag
	cmd = exec.Command("./test_cli_help", "-help")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err = cmd.Run()
	// Help typically exits with code 2 in Go's flag package
	if exitError, ok := err.(*exec.ExitError); ok {
		assert.Equal(t, 2, exitError.ExitCode())
	}

	helpOutput := stdout.String()

	// Verify help output contains expected flag descriptions
	expectedFlags := []string{"-name", "-level", "-class", "-lineage", "-heritage", "-traits"}
	for _, flag := range expectedFlags {
		assert.Contains(t, helpOutput, flag, fmt.Sprintf("Help should contain %s flag", flag))
	}
}
