package character

import (
	"tov_tools/pkg/helpers"
)

var PredefinedTraitsData = map[string]LineagePreDefinedTraits{

	"dwarf": {
		Lineage: "Dwarf",
		Traits: map[string]string{
			"Darkvision":         "You have a superior vision in dark and dim conditions for 60 feet.",
			"Dwarven Resilience": "You have advantage on saving throws against poison, and you have resistance against poison damage.",
			"Dwarven Toughness":  "Your hit point maximum increases by 1, and it increases by 1 every time you gain a level.",
		},
	},
	"elf": {
		Lineage: "Elf",
		Traits: map[string]string{
			"Heightened Senses": "You have advantage on Perception checks that rely on sight or hearing.  You can see through lightly obscured areas normally and areas of dim light as if it were bright light.",
			"Magic Ancestry":    "You have advantage on saving throws against being charmed, and magic can't put you to sleep.",
			"Trance":            "Elves don't need to sleep. Instead, they enter a meditative trance state, remaining semiconscious for 4 hours a day. You choose whether or not you can dream while meditating. After resting in this way, you gain the same benefit that other creatures do from 8 hours of sleep.",
		},
	},
	"human": {
		Lineage: "Human",
		Traits: map[string]string{
			"Ambitious": "You gain proficiency in one skill of your choice, and you gain one talent of your choice. This talent can be from any of the talent lists, but you must meet the talent's prerequisites if any are required.",
		},
	},
	"kobold": {
		Lineage: "Kobold",
		Traits: map[string]string{
			"Darkvision":           "You have a superior vision in dark and dim conditions for 60 feet.",
			"Tinker's Fascination": "Your innate fascination with how things work allows you to use tools with ease. When you make an ability check, you can roll a d8 and add the result to the check.",
		},
	},
	"orc": {
		Lineage: "Orc",
		Traits: map[string]string{
			"Heightened Senses":   "You have advanage on Perception checks that rely on sight or hearing.  You can see through lightly obscured areas normally and areas of dim light as if it were bright light.",
			"Orcish Perseverance": "When you would die due to suffocation or gaining levels of exhaustion, you instead enter a death-like stasis. While in stasis you are incapacitated, can't move, can't speak, and are unaware of your surroundings. You also cease to age, and your body is protected from decay. You can remain in this state until you are restored by mundane or magical healing, or your body is completely destroyed.",
			"Stalwart":            "When you are subjected to an effect that requires you to make a save at the end of your turn, you can instead choose to make the save at the start of your turn.",
		},
	},
	"syderean": {
		Lineage: "Syderean",
		Traits: map[string]string{
			"Far Sight":         "You have darkvision to a range of 60 feet and can see in magical darkness to a range of 30 feet.",
			"Otherworldly Form": "You have resistance to necrotic damage and the amount of time you survive without air, food, water, or sleep is double that of a typical character.",
		},
	},
	// Add more predefined Traits for other lineages as needed
}

var LineageNaturalAdaptations = map[string]LineageNaturalAdaptationTraitDescriptions{
	"beastkin": LineageNaturalAdaptationTraitDescriptions{
		Lineage: "Beastkin",
		Traits: map[string]map[string]string{
			"Natural Adaptation": {
				"Avian":   "Leathery, feathery, or gossamer wings sprout from our back and/or connect to your outstretched arms. You have a flying speed equal to your walking speed. You can't fly while wearing medium or heavy armor. Carrying a heavy load or a creature of your size or larger while flying reduces your speed by half",
				"Agile":   "Sharp claws, cloven hooves, robust limbs, reversed joints, or suction pads on your hands and feet allow you to scrabble up trees, rock faces, and uneven surfaces with ease.  You have a climbing speed equal to your walking speed. In addition, you have advantage on saves made to avoid being knocked prone.",
				"Aquatic": "Oily fur, scales, fins, webbed hands and feet, or long, gangly limbs are common for your people. You have a swimming speed equal to your walking speed. You an hold your breath for up to 20 minutes.",
				"Sturdy":  "Powerful limbs, fat reserves, or a thick hide allow you to endure harm and accomplish feat of remarkable athleticism. When you arent wearing armor, you have a natural AC of 13+ your DEX modifier. In addition, you count as one size larger when determining your carrying capacity and the weight you can push or drag.",
			},
			"Animal Instinct": {
				"Perception": "You have proficiency in the Perception skill.",
				"Survival":   "You have proficiency in the Survival skill.",
			},
			"Natural Weapons": {
				"Other":  "You have some other feature that serves as a natural weapon.",
				"Claws":  "You have claws that serve as a natural weapon.",
				"Horns":  "You have horns that serve as a natural weapon.",
				"Hooves": "You have hooves that serve as a natural weapon.",
				"Fangs":  "You have fangs that serve as a natural weapon.",
				"Spines": "You have spines that serve as a natural weapon.",
			},
		},
	},
	"kobold": LineageNaturalAdaptationTraitDescriptions{
		Lineage: "Kobold",
		Traits: map[string]map[string]string{
			"Natural Adaptation": {
				"Fierce (Small)":     "When a Large or larger creature you can see within 5 feet of you attacks you, you can use your reaction to attack that creature immediately after its attack.",
				"Truescale (Medium)": "You naturally thick scales provide significant protection. You have a natural AC of 13+ your DEX modifier. In addition, you have resistance to one of the following types of damage: acid, cold, fire, lightning, or poison (choose during character creation",
			},
		},
	},
	"syderean": LineageNaturalAdaptationTraitDescriptions{
		Lineage: "Syderean",
		Traits: map[string]map[string]string{
			"Natural Adaptation": {
				"Celestial": "You possess notable physical characteristics that mark your connection to realms of good and order. You might have luminous eyes, metallic-hued skin, or possess the ability to stay perfectly still for hours. You also gain Blessed Guise.  Once per long rest, you can use a bonus action to assume an otherworldly guise for 1 minute. When you do so, you sprout spectral wings and gain a flying speed equal to you walking speed for the duration of your transformation.  While transformed, once on each of your turns when you deal damage with an attack or spell, you can choose to convert the damage type to radiant damage.",
				"Fiendish":  "You possess notable physical characteristics that mark your connection to realms of evil or chaos. You might bear bony horns that jut from your skull, emit a perpetual odor of smoke, or have a barbed tail. You also gain Dreadful Guise. Once per long rest, you can use a bonus action to assume an otherworldy guise for 1 minute. While the transformation lasts, creatures of you choice that come within 10 feet of you for the first time on a turn or start their turn there must succeed on a CHA save (DC equals 10 + your PB) or become frightened of you until the end of your next turn.  Once a creature succeeds on this save, they can't be affected by this feature again for 24 hours. While transformed, once on each of your turns when you deal damage with an attack or spell, you choose to convert the damage type to necrotic or fire damage.",
			},
		},
	},
	"smallfolk": LineageNaturalAdaptationTraitDescriptions{
		Lineage: "Smallfolk",
		Traits: map[string]map[string]string{
			"Natural Adaptation": {
				"Gnomish":  "Your ancestors made their homes underground or in dark woodlands, providing you with darkvision to a range of 60 feet. In addition, you learn the minor illusion cantrip.  Choose whether INT, WIS, or CHA is you spellcasting ability for this spell during character creation.",
				"Halfling": "Your ancestors made their homes on the surface, displaying outsized pluck to drive back hostile wildlife and monsters. You have advantage on saves against being charmed or frightened.",
			},
		},
	},
}

var Lineages = map[string]Lineage{
	"beastkin": Lineage{
		Name:         "Beastkin",
		MaturityAge:  5,
		AgeDiceSides: 8,
		AgeDiceRolls: 1,
		SizeOptions:  []string{"Medium", "Small"},
		Speed:        30,
		TraitOptions: map[string]ChoiceOptions{
			"Animal Instinct": ChoiceOptions{
				NumberToSelect: 1,
				Options:        []string{"Perception", "Survival"},
			},
			"Natural Weapons": ChoiceOptions{
				NumberToSelect: 1,
				Options:        []string{"Claws", "Horns", "Hooves", "Fangs", "Spines"},
			},
			"Natural Adaptation": ChoiceOptions{
				NumberToSelect: 1,
				Options:        helpers.GetMapKeys(LineageNaturalAdaptations["beastkin"].Traits),
			},
		},
		LineageSource: "Players Guide, pg 105",
	},
	"dwarf": Lineage{
		Name:          "Dwarf",
		MaturityAge:   50,
		AgeDiceSides:  20,
		AgeDiceRolls:  5,
		SizeOptions:   []string{"Medium"},
		Speed:         30,
		Traits:        helpers.GetMapKeys(PredefinedTraitsData["dwarf"].Traits),
		LineageSource: "Players Guide, pg 106",
	},
	"elf": Lineage{
		Name:          "Elf",
		MaturityAge:   100,
		AgeDiceSides:  20,
		AgeDiceRolls:  8,
		SizeOptions:   []string{"Medium"},
		Speed:         30,
		Traits:        helpers.GetMapKeys(PredefinedTraitsData["elf"].Traits),
		LineageSource: "Players Guide, pg 106",
	},
	"human": Lineage{
		Name:          "Human",
		MaturityAge:   18,
		AgeDiceSides:  10,
		AgeDiceRolls:  2,
		SizeOptions:   []string{"Small", "Medium"},
		Speed:         30,
		Traits:        helpers.GetMapKeys(PredefinedTraitsData["human"].Traits),
		LineageSource: "Players Guide, pg 107",
	},
	"kobold": Lineage{
		Name:         "Kobold",
		MaturityAge:  14,
		AgeDiceSides: 20,
		AgeDiceRolls: 2,
		SizeOptions:  []string{"Small"},
		Speed:        30,
		Traits:       helpers.GetMapKeys(PredefinedTraitsData["kobold"].Traits),
		TraitOptions: map[string]ChoiceOptions{
			"Natural Adaptation": ChoiceOptions{
				NumberToSelect: 1,
				Options:        helpers.GetMapKeys(LineageNaturalAdaptations["kobold"].Traits),
			},
		},
		LineageSource: "Players Guide, pg 108",
	},
	"orc": Lineage{
		Name:          "Orc",
		MaturityAge:   20,
		AgeDiceSides:  10,
		AgeDiceRolls:  2,
		SizeOptions:   []string{"Medium"},
		Speed:         30,
		Traits:        helpers.GetMapKeys(PredefinedTraitsData["orc"].Traits),
		LineageSource: "Players Guide, pg 108",
	},
	"syderean": Lineage{
		Name:         "Syderean",
		MaturityAge:  20,
		AgeDiceSides: 20,
		AgeDiceRolls: 3,
		SizeOptions:  []string{"Medium"},
		Speed:        30,
		TraitOptions: map[string]ChoiceOptions{
			"Natural Adaptation": ChoiceOptions{
				NumberToSelect: 1,
				Options:        helpers.GetMapKeys(LineageNaturalAdaptations["syderean"].Traits),
			},
		},
		Traits:        helpers.GetMapKeys(PredefinedTraitsData["syderean"].Traits),
		LineageSource: "Players Guide, pg 109",
	},
	"smallfolk": Lineage{
		Name:         "Smallfolk",
		MaturityAge:  20,
		AgeDiceSides: 20,
		AgeDiceRolls: 4,
		SizeOptions:  []string{"Small"},
		Speed:        30,
		TraitOptions: map[string]ChoiceOptions{
			"Natural Adaptation": ChoiceOptions{
				NumberToSelect: 1,
				Options:        helpers.GetMapKeys(LineageNaturalAdaptations["smallfolk"].Traits),
			},
		},
		Traits:        helpers.GetMapKeys(PredefinedTraitsData["smallfolk"].Traits),
		LineageSource: "Players Guide, pg 109",
	},
}
