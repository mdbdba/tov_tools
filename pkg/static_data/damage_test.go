package static_data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDamageModifiers(t *testing.T) {
	actual := DamageModifiers()
	assert.Equal(t, 3, len(actual))
	assert.Equal(t, float32(2), actual["vulnerable"])
}

func TestDamageType(t *testing.T) {
	actual := DamageType()
	compareValue := "Radiant damage, dealt by a cleric's flame strike spell or an angel's smiting weapon, sears the flesh like fire and overloads the spirit with power."
	assert.Equal(t, 13, len(actual))
	assert.Equal(t, compareValue, actual["radiant"])
}
