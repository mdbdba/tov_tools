package character

import (
	"math/rand"
	"time"
)

func getRandomGen() *rand.Rand {
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	randomGenerator := rand.New(randomSource)
	return randomGenerator
}

// RandomAge generates a random age for a character based on its lineage
func RandomAge(lineage Lineage) int {
	randomGenerator := getRandomGen()
	age := lineage.MaturityAge
	for i := 0; i < lineage.AgeDiceRolls; i++ {
		age += randomGenerator.Intn(lineage.AgeDiceSides) + 1
	}
	return age
}

// RandomClass returns a randomly selected Class
func RandomClass() Class {

	randomGenerator := getRandomGen()

	keys := make([]string, 0, len(Classes))
	for key := range Classes {
		keys = append(keys, key)
	}
	randomKey := keys[randomGenerator.Intn(len(keys))]
	return Classes[randomKey]
}

// RandomLineage returns a randomly selected Lineage
func RandomLineage() Lineage {

	randomGenerator := getRandomGen()

	keys := make([]string, 0, len(Lineages))
	for key := range Lineages {
		keys = append(keys, key)
	}
	randomKey := keys[randomGenerator.Intn(len(keys))]
	return Lineages[randomKey]
}

// RandomSize returns a random size from Lineage options
func RandomSize(lineage Lineage) string {
	if len(lineage.SizeOptions) == 1 {
		return lineage.SizeOptions[0]
	}
	randomGenerator := getRandomGen()
	randomIndex := randomGenerator.Intn(len(lineage.SizeOptions))
	return lineage.SizeOptions[randomIndex]
}

func getClassBuildTypes(classBuildTypes map[string]ClassBuildType) []string {
	keys := make([]string, 0, len(classBuildTypes))
	for key := range classBuildTypes {
		keys = append(keys, key)
	}
	return keys
}

func ValidateClassBuildType(classBuildType string, classBuildTypes map[string]ClassBuildType) bool {
	_, exists := classBuildTypes[classBuildType]
	return exists
}

func RandomClassBuildType(classBuildTypes map[string]ClassBuildType) string {

	randomGenerator := getRandomGen()
	keys := getClassBuildTypes(classBuildTypes)
	randomKey := keys[randomGenerator.Intn(len(keys))]
	return randomKey

}
