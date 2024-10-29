// pkg/character/class_test.go
package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClass(t *testing.T) {
	class, err := GetClass("barbarian")
	assert.NoError(t, err)
	assert.Equal(t, "Barbarian", class.Name)
	assert.Equal(t, "d12", class.HitDie)
	assert.Equal(t, [][]string{{"str"}}, class.KeyAbilities)
	assert.ElementsMatch(t, []string{"str", "con"}, class.SaveProficiencies)
	assert.ElementsMatch(t, []string{"light armor", "medium armor", "shields", "weapons"}, class.EquipmentProficiencies)

	class, err = GetClass("ranger")
	assert.NoError(t, err)
	assert.Equal(t, "Ranger", class.Name)
	assert.Equal(t, "d10", class.HitDie)
	assert.Equal(t, [][]string{{"dex", "wis"}}, class.KeyAbilities) // both str and dex are possible key abilities
	assert.ElementsMatch(t, []string{"str", "dex"}, class.SaveProficiencies)
	assert.ElementsMatch(t, []string{"light armor", "medium armor", "shields", "weapons"}, class.EquipmentProficiencies)

	_, err = GetClass("nonexistentclass")
	assert.Error(t, err)
}

func TestGetClassByName(t *testing.T) {
	_, err := GetClassByName("barbarian")
	assert.NoError(t, err)
	_, err = GetClassByName("nonexistent")
	assert.Error(t, err)
}
