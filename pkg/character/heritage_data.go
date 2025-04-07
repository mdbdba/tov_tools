package character

var Heritages = map[string]Heritage{
	"anointed": {
		Name:                "Anointed",
		LanguageDefaults:    []string{"Common"},
		LanguageSuggestions: []string{"Abyssal", "Celestial", "Infernal"},
		Traits: map[string]string{
			"Favored Disciple": "You know the thaumaturgy cantrip and you have advantage on death saves.",
			"Occult Studies": "When you make a check to recall " +
				"or interpret information about Celestials, Fiends, or creatures with the " +
				"Outsider tag, you can make a skill check with advantage.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 2,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Occult Studies Skills": {
				NumberToSelect: 1,
				Options:        []string{"History", "Religion"},
			},
		},
	},
	"cloud": {
		Name:                "Cloud",
		LanguageDefaults:    []string{"Common"},
		LanguageSuggestions: []string{"Elvish", "Draconic"},
		Traits: map[string]string{
			"World of Wonders": "You have proficiency in the Arcana skill.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 2,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Touch of Magic (school)": {
				NumberToSelect: 1,
				Options:        []string{"Arcane", "Divine", "Primordial", "Wyrd"},
			},
			"Touch of Magic (spell casting ability)": {
				NumberToSelect: 1,
				Options:        []string{"cha", "int", "wis"},
			},
		},
	},
	"cosmopolitan": {
		Name:                "Cosmopolitan",
		LanguageDefaults:    []string{"Common"},
		LanguageSuggestions: []string{"Elvish", "Dwarvish"},
		Traits: map[string]string{
			"Street Smarts": "While in a city or other urban environment, you have advantage on ability checks made " +
				"to avoid getting lost and checks made to find a particular kind of business or other destination " +
				"open to the public. In addition, while you are in such environments, you can’t be surprised unless " +
				"you are asleep or otherwise incapacitated.",
			"Worldly Wisdom": "You have proficiency in the History skill. When you make a check related to " +
				"understanding the purpose or significance of a building, rite, or object from a culture you aren’t " +
				"familiar with, you can add your PB to the roll. If you have proficiency in a relevant skill or tool, " +
				"double your PB for the roll.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 3,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
	"cottage": {
		Name:                "Cottage",
		LanguageDefaults:    []string{"Common"},
		LanguageSuggestions: []string{"Halfling", "Gnomish"},
		Traits: map[string]string{
			"Comforts of Home": "As part of a long rest, you can cook a meal, tell stories, " +
				"or perform some other activity that comforts your allies. Choose a number of " +
				"creatures who participated in the long rest equal to your PB (this can include " +
				"you). Those creatures gain temporary HP equal to twice your PB. These temporary " +
				"HP last until expended or until you complete your next long rest.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Homesteader": {
				NumberToSelect: 1,
				Options:        []string{"Animal Handling", "Nature"},
			},
		},
	},
	"diaspora": {
		Name:                "Diaspora",
		LanguageDefaults:    []string{"Common"},
		LanguageSuggestions: []string{"Orcish", "Dwarvish"},
		Traits: map[string]string{
			"Preserved Traditions (skill)": "You gain proficiency in the history skill.",
			"Timeless Resolve": "When you or an allied creature within 5 feet of you makes " +
				"a save against becoming frightened, you and the ally have " +
				"advantage on the save.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Preserved Traditions (weapon proficiency)": {
				NumberToSelect: 1,
				// todo: Implement a list of martial weapons to choose from
				Options: []string{"Sword", "WarHammer"},
			},
		},
	},
}
