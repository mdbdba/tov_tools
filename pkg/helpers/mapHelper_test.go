package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapStringIntToString(t *testing.T) {
	src := map[string]int{
		"all":    1,
		"the":    2,
		"things": 3,
	}
	actual := MapStringIntToString(src)

	expected := "[\"all\": 1, \"the\": 2, \"things\": 3]"
	assert.Equal(t, expected, actual)
}
