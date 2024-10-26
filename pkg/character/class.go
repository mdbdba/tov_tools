package character

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"text/tabwriter"
	"time"
)

type Class struct {
	Name                   string
	Description            string
	HitDie                 string
	KeyAbilities           [][]string
	SaveProficiencies      []string
	EquipmentProficiencies []string
}

// GetClass retrieves the Class for a given name
func GetClass(name string) (Class, error) {
	class, exists := Classes[name]
	if !exists {
		return Class{}, fmt.Errorf("class %s does not exist", name)
	}
	return class, nil
}

// RandomClass returns a randomly selected Class
func RandomClass() Class {
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	randomGenerator := rand.New(randomSource)

	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	keys := make([]string, 0, len(Classes))
	for key := range Classes {
		keys = append(keys, key)
	}
	randomKey := keys[randomGenerator.Intn(len(keys))]
	return Classes[randomKey]
}

// GetClassByName returns a Class by its name or an error if it doesn't exist
func GetClassByName(name string) (Class, error) {
	lowerName := strings.ToLower(name)
	class, exists := Classes[lowerName]
	if !exists {
		return Class{}, fmt.Errorf("class '%s' does not exist", name)
	}
	return class, nil
}

// ToString formats a single Class as a string
func (c *Class) ToString() string {
	return fmt.Sprintf(
		"Name: %s\nDescription: %s\nHit Die: %s\nKey Abilities: %s\nSave Proficiencies: %s\nEquipment Proficiencies: %s\n",
		c.Name,
		c.Description,
		c.HitDie,
		formatKeyAbilities(c.KeyAbilities),
		strings.Join(c.SaveProficiencies, ", "),
		strings.Join(c.EquipmentProficiencies, ", "),
	)
}

// ToStringTable formats all classes as a string table
func ToStringTable() string {
	var buf bytes.Buffer
	writer := tabwriter.NewWriter(&buf, 0, 0, 3, ' ', 0)
	_, err := fmt.Fprintf(writer, "Name\tDescription\tHit Die\tKey Abilities\tSave Proficiencies\tEquipment Proficiencies\n")
	if err != nil {
		return ""
	}
	for _, class := range Classes {
		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\n",
			class.Name,
			class.Description,
			class.HitDie,
			formatKeyAbilities(class.KeyAbilities),
			strings.Join(class.SaveProficiencies, ", "),
			strings.Join(class.EquipmentProficiencies, ", "))
		if err != nil {
			return ""
		}
	}
	err = writer.Flush()
	if err != nil {
		return ""
	}
	return buf.String()
}

func formatKeyAbilities(abilities [][]string) string {
	var formattedParts []string
	for _, group := range abilities {
		formattedParts = append(formattedParts, "["+strings.Join(group, ", ")+"]")
	}
	return strings.Join(formattedParts, " or ")
}
