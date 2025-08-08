package types

import "time"

// CharacterCreateRequest represents the request body for creating a character
type CharacterCreateRequest struct {
	Name             string            `json:"name" binding:"required"`
	Level            *int              `json:"level,omitempty"`
	Class            string            `json:"class" binding:"required"`
	Subclass         string            `json:"subclass,omitempty"`
	Lineage          string            `json:"lineage" binding:"required"`
	Heritage         string            `json:"heritage" binding:"required"`
	Size             string            `json:"size,omitempty"`
	AbilityGenMethod string            `json:"ability_generation_method,omitempty"`
	Traits           map[string]string `json:"traits,omitempty"`
	Talents          []string          `json:"talents,omitempty"`
	Languages        []string          `json:"languages,omitempty"`
}

// CharacterResponse represents the response structure for character operations
type CharacterResponse struct {
	ID               int               `json:"id"`
	Name             string            `json:"name"`
	Level            int               `json:"level"`
	Class            string            `json:"class"`
	Subclass         string            `json:"subclass,omitempty"`
	Lineage          string            `json:"lineage"`
	Heritage         string            `json:"heritage"`
	Size             string            `json:"size"`
	AbilityScores    map[string]int    `json:"ability_scores"`
	AbilityModifiers map[string]int    `json:"ability_modifiers"`
	Traits           map[string]string `json:"traits"`
	Talents          []string          `json:"talents"`
	Languages        []string          `json:"languages"`
	CreatedAt        time.Time         `json:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at"`
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"`
}
