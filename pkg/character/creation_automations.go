package character

import (
	"math/rand"
	"time"
)

// RandomAge generates a random age for a character based on its lineage
func RandomAge(lineage Lineage) int {
	rand.Seed(time.Now().UnixNano())
	age := lineage.MaturityAge
	for i := 0; i < lineage.AgeDiceRolls; i++ {
		age += rand.Intn(lineage.AgeDiceSides) + 1
	}
	return age
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
