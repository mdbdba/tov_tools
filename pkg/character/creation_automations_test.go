package character

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// TestRandomAge tests the RandomAge function to ensure it generates an age within the expected range
func TestRandomAge(t *testing.T) {
	lineage := Lineage{
		Name:         "Testkin",
		MaturityAge:  15,
		AgeDiceSides: 12,
		AgeDiceRolls: 2,
	}
	for i := 0; i < 100; i++ { // Run the test multiple times to account for randomness
		age := RandomAge(lineage)
		if age < lineage.MaturityAge || age > lineage.MaturityAge+lineage.AgeDiceSides*lineage.AgeDiceRolls {
			t.Errorf("Generated age %d is out of expected range [%d, %d]", age, lineage.MaturityAge, lineage.MaturityAge+lineage.AgeDiceSides*lineage.AgeDiceRolls)
		}
	}
}

func TestRandomClass(t *testing.T) {
	randomClass := RandomClass()
	lowerName := strings.ToLower(randomClass.Name)
	_, exists := Classes[lowerName]
	assert.True(t, exists, "Randomly selected class should exist in Classes")
}
