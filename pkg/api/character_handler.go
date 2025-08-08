package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"tov_tools/pkg/character"
	"tov_tools/pkg/types"
)

// In-memory storage for characters (in a real app, this would be a database)
var (
	characters       = make(map[int]*character.Character)
	charactersByName = make(map[string]*character.Character)
	nextID           = 1
	charMutex        sync.RWMutex
)

// CreateCharacter handles POST /api/v1/character/create
func CreateCharacter(c *gin.Context) {
	var req types.CharacterCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set defaults
	level := 1
	if req.Level != nil {
		level = *req.Level
	}

	if req.AbilityGenMethod == "" {
		req.AbilityGenMethod = "standard"
	}

	if req.Traits == nil {
		req.Traits = make(map[string]string)
	}

	if req.Talents == nil {
		req.Talents = []string{}
	}

	if req.Languages == nil {
		req.Languages = []string{}
	}

	// Validate lineage
	if !character.ValidateLineage(req.Lineage) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid lineage: %s", req.Lineage)})
		return
	}

	// Check if character name already exists
	charMutex.RLock()
	if _, exists := charactersByName[strings.ToLower(req.Name)]; exists {
		charMutex.RUnlock()
		c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("character with name '%s' already exists", req.Name)})
		return
	}
	charMutex.RUnlock()

	// Create logger for character creation
	observedZapCore, _ := observer.New(zap.InfoLevel)
	logger := zap.New(observedZapCore).Sugar()

	ctxRef := fmt.Sprintf("api character creation: %s", req.Name)

	// Determine size if not provided
	size := req.Size
	if size == "" {
		if lineageData, exists := character.Lineages[req.Lineage]; exists && len(lineageData.SizeOptions) > 0 {
			size = lineageData.SizeOptions[0]
		} else {
			size = "Medium" // default fallback
		}
	}

	// Create the character
	char, err := character.NewCharacter(
		req.Name,
		level,
		req.Class,
		req.Subclass,
		req.Lineage,
		req.Heritage,
		size,
		req.AbilityGenMethod,
		req.Traits,
		req.Talents,
		req.Languages,
		"Standard", // buildType
		character.ClassBuildType{},
		ctxRef,
		logger,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to create character: %v", err)})
		return
	}

	// Store character with generated ID
	charMutex.Lock()
	id := nextID
	nextID++
	characters[id] = char
	charactersByName[strings.ToLower(req.Name)] = char
	charMutex.Unlock()

	// Convert to response format
	response := convertToCharacterResponse(char, id)

	c.JSON(http.StatusCreated, response)
}

// GetCharacterByName handles GET /api/v1/character/name/{name}
func GetCharacterByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "character name is required"})
		return
	}

	charMutex.RLock()
	char, exists := charactersByName[strings.ToLower(name)]
	charMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("character with name '%s' not found", name)})
		return
	}

	// Find the ID for this character
	var id int
	charMutex.RLock()
	for charID, storedChar := range characters {
		if storedChar == char {
			id = charID
			break
		}
	}
	charMutex.RUnlock()

	response := convertToCharacterResponse(char, id)
	c.JSON(http.StatusOK, response)
}

// GetCharacterByID handles GET /api/v1/character/id/{id}
func GetCharacterByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid character ID"})
		return
	}

	charMutex.RLock()
	char, exists := characters[id]
	charMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("character with ID %d not found", id)})
		return
	}

	response := convertToCharacterResponse(char, id)
	c.JSON(http.StatusOK, response)
}

// GetAllCharacters handles GET /api/v1/characters
func GetAllCharacters(c *gin.Context) {
	charMutex.RLock()
	defer charMutex.RUnlock()

	var responses []types.CharacterResponse
	for id, char := range characters {
		response := convertToCharacterResponse(char, id)
		responses = append(responses, response)
	}

	c.JSON(http.StatusOK, gin.H{"characters": responses})
}

// UpdateCharacter handles PUT /api/v1/character/id/{id}
func UpdateCharacter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid character ID"})
		return
	}

	charMutex.RLock()
	char, exists := characters[id]
	charMutex.RUnlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("character with ID %d not found", id)})
		return
	}

	var req types.CharacterCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// For now, we'll just update simple fields
	// In a full implementation, you might want to recreate the character
	// or have specific update methods

	// Remove old name mapping if name changed
	if strings.ToLower(char.Name) != strings.ToLower(req.Name) {
		charMutex.Lock()
		delete(charactersByName, strings.ToLower(char.Name))
		charactersByName[strings.ToLower(req.Name)] = char
		charMutex.Unlock()

		char.Name = req.Name
	}

	response := convertToCharacterResponse(char, id)
	c.JSON(http.StatusOK, response)
}

// DeleteCharacter handles DELETE /api/v1/character/id/{id}
func DeleteCharacter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid character ID"})
		return
	}

	charMutex.Lock()
	char, exists := characters[id]
	if exists {
		delete(characters, id)
		delete(charactersByName, strings.ToLower(char.Name))
	}
	charMutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("character with ID %d not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("character with ID %d deleted successfully", id)})
}

// convertToCharacterResponse converts a character.Character to CharacterResponse
func convertToCharacterResponse(char *character.Character, id int) types.CharacterResponse {
	abilityScores := make(map[string]int)
	abilityModifiers := make(map[string]int)

	abilityScores = char.Abilities.Values
	abilityModifiers = char.Abilities.Modifiers

	// Convert Talents map to slice of talent names
	talentNames := make([]string, 0, len(char.Talents))
	for talentName := range char.Talents {
		talentNames = append(talentNames, talentName)
	}

	return types.CharacterResponse{
		ID:               id,
		Name:             char.Name,
		Level:            char.OverallLevel,
		Class:            char.CharacterClassStr,
		Subclass:         char.CharacterSubClassToImplement.Name,
		Lineage:          char.Lineage.Name,
		Heritage:         char.Heritage.Name,
		Size:             char.CharacterSize,
		AbilityScores:    abilityScores,
		AbilityModifiers: abilityModifiers,
		Traits:           char.Traits,
		Talents:          talentNames,
		Languages:        char.KnownLanguages,
		CreatedAt:        time.Now(), // In a real app, this would be stored
		UpdatedAt:        time.Now(),
	}
}
