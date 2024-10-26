package character

var Lineages = map[string]Lineage{
	"beastkin": Lineage{
		Name:        "Beastkin",
		AgeDiceRoll: "d12",
		SizeOptions: []string{"Medium", "Small"},
		Speed:       30,
		TraitOptions: map[string]TraitChoices{
			"Animal Instinct": TraitChoices{
				NumberToSelect: 1,
				Options:        []string{"Perception", "Survival"},
			},
			"Natural Weapons": TraitChoices{
				NumberToSelect: 1,
				Options:        []string{"Claws", "Horns", "Hooves", "Fangs", "Spines"},
			},
			"Natural Adaptation": TraitChoices{
				NumberToSelect: 1,
				Options:        []string{"Avian", "Agile", "Aquatic", "Sturdy"},
			},
		},
		LineageSource: "Players Guide, pg 105",
	},
	// Add more lineages here as needed
}
