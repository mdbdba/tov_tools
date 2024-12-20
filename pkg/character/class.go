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

type Class struct {
	Name                   string
	ClassBuildTypes        map[string]ClassBuildType
	Description            string
	HitDie                 string
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
