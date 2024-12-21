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