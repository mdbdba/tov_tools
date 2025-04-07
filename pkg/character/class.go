package character

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"tov_tools/pkg/helpers"
)

type ClassBuildType struct {
	KeyAbilities                []string
	AbilityScoreOrderPreference []string
}

type Subclass struct {
	Name                string
	Description         string
	SpellcastingAbility SpellcastingAbilityType // Optional: Exists only if the subclass grants it
}

// SpellcastingAbilityType defines a custom type for allowed spellcasting Abilities
type SpellcastingAbilityType string

// Allowed spellcasting ability values
const (
	Str SpellcastingAbilityType = "str"
	Dex SpellcastingAbilityType = "dex"
	Con SpellcastingAbilityType = "con"
	Int SpellcastingAbilityType = "int"
	Wis SpellcastingAbilityType = "wis"
	Cha SpellcastingAbilityType = "cha"
)

// IsValid checks if a value is a valid SpellcastingAbilityType
func (s SpellcastingAbilityType) IsValid() bool {
	switch s {
	case Str, Dex, Con, Int, Wis, Cha:
		return true
	}
	return false
}

type Class struct {
	Name                   string
	ClassBuildTypes        map[string]ClassBuildType
	Description            string
	HitDie                 string
	SaveProficiencies      []string
	EquipmentProficiencies []string
	SpellcastingAbility    SpellcastingAbilityType
	Subclasses             map[string]Subclass
}

// SetSpellcastingAbility sets the SpellcastingAbility for the Class, with validation
func (c *Class) SetSpellcastingAbility(value string) error {
	ability := SpellcastingAbilityType(value)
	if !ability.IsValid() {
		return fmt.Errorf("invalid SpellcastingAbility: %s", value)
	}
	c.SpellcastingAbility = ability
	return nil
}

func (c *Class) GetSubclass(name string) (Subclass, error) {
	subclass, exists := c.Subclasses[name]
	if !exists {
		return Subclass{},
			fmt.Errorf("subclass '%s' does not exist for class '%s'", name, c.Name)
	}
	return subclass, nil
}

// GetClass retrieves the Class for a given Name
func GetClass(name string) (Class, error) {
	class, exists := Classes[name]
	if !exists {
		return Class{}, fmt.Errorf("class %s does not exist", name)
	}
	return class, nil
}

// GetClassByName returns a Class by its Name or an error if it doesn't exist
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
	buildTypeString := "}"
	for buildName, buildType := range c.ClassBuildTypes {
		buildTypeString += fmt.Sprintf(
			"{ Build Type: %s Key Abilities: %s Ability Scoare Order Preference: %s } ",
			buildName,
			helpers.StringSliceToString(buildType.KeyAbilities),
			helpers.StringSliceToString(buildType.AbilityScoreOrderPreference))
	}
	buildTypeString += "}"
	return fmt.Sprintf(
		"Name: %s\nDescription: %s\nHit Die: %s\nBuild Types: %s\nSave Proficiencies: %s\nEquipment Proficiencies: %s\n",
		c.Name,
		c.Description,
		c.HitDie,
		buildTypeString,
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
		abilityScoreOrderPreferenceString := ""
		keyAbilitiesString := ""
		for buildName, buildType := range class.ClassBuildTypes {
			abilityScoreOrderPreferenceString +=
				fmt.Sprintf("{%s: %s} ", buildName,
					helpers.StringSliceToString(buildType.AbilityScoreOrderPreference))
			keyAbilitiesString +=
				fmt.Sprintf("{%s: %s} ", buildName,
					helpers.StringSliceToString(buildType.KeyAbilities))
		}

		_, err = fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			class.Name,
			class.Description,
			class.HitDie,
			abilityScoreOrderPreferenceString,
			keyAbilitiesString,
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
