package character

import "fmt"

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
