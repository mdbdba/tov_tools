package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSliceToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"Normal Slice", []string{"test", "all", "things"}, "[\"test\", \"all\", \"things\"]"},
		{"Empty Slice", []string{}, "[]"},
		{"Single Element", []string{"single"}, "[\"single\"]"},
		{"Multiple Empty Strings", []string{"", ""}, "[\"\", \"\"]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := StringSliceToString(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name        string
		length      int
		expectError bool
	}{
		{"Valid Length 5", 5, false},
		{"Valid Length 10", 10, false},
		{"Zero Length", 0, false},
		{"Negative Length", -5, true}, // This will naturally fail; valid test scenario.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GenerateRandomString(tt.length)
			if tt.expectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.length, len(actual))
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		item     string
		expected bool
	}{
		{"Exists in list", []string{"a", "b", "c"}, "b", true},
		{"Does not exist in list", []string{"a", "b", "c"}, "d", false},
		{"Empty list", []string{}, "a", false},
		{"Nil list", nil, "a", false},
		{"Case Sensitivity Test", []string{"A", "b", "C"}, "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Contains(tt.slice, tt.item)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestToTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Normal Case", "test all things", "Test All Things"},
		{"Empty String", "", ""},
		{"Single Word", "test", "Test"},
		{"Mixed Case", "TeSt aLL tHiNgS", "Test All Things"},
		{"With Punctuation", "hello, world!", "Hello, World!"},
		{"Multiple Spaces", "hello    world", "Hello World"},
		{"Leading and Trailing Spaces", "  hello world  ", "Hello World"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ToTitleCase(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
