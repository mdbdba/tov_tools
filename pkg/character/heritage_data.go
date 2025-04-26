package character

var Heritages = map[string]Heritage{
	"anointed": {
		Name:                "Anointed",
		LanguageDefaults:    []string{"Common"},
		LanguageSuggestions: []string{"Abyssal", "Celestial", "Infernal"},
		LanguageSuggestionNote: "Typical anointed heritage characters choose an esoteric language aligned with " +
			"their guiding power.",
		HeritageSource: "Players Guide, pg 112",
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
		Name:                   "Cloud",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Elvish", "Draconic"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 113",
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
		Name:                   "Cosmopolitan",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Elvish", "Dwarvish"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 113",
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
		Name:                   "Cottage",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Halfling", "Gnomish"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 113",
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
		LanguageSuggestionNote: "Typically, the languages common to soldiers, mercenaries, and traders near where they " +
			"reside.",
		HeritageSource: "Players Guide, pg 113",
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
	"fireforge": {
		Name:                   "Fireforge",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Dwarvish"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 114",
		Traits: map[string]string{
			"Forgecraft (tools proficiency)": "You gain proficiency with smithing tools.",
			"Forgecraft (cantrip)":           "You know the mending cantrip.",
			"Heat Resilience": "Lifelong exposure has made you resilient to the effects of severe heat. " +
				"You are resistant to fire damage.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
	"grove": {
		Name:                   "Grove",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Elvish"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 114",
		Traits: map[string]string{
			"Canopy Walker": "You have a climbing speed equal to your walking speed.",
			"Nature's Camouflage": "You have advantage on dex (Stealth) checks made while you are lightly obscured " +
				"by foliage, heavy rain, falling snow, mist, and other natural phenomena. While in such conditions, " +
				"you can always attempt to take the Hide action, even if circumstances would not normally allow you to do so.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
	"nomadic": {
		Name:                   "Nomadic",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Dwarvish", "Elvish"},
		LanguageSuggestionNote: "Typically, the languages of the communities with witch your people trade.",
		HeritageSource:         "Players Guide, pg 115",
		Traits: map[string]string{
			"Resilient (weather effects)": "You have advantage on checks or saves made to resist debilitating weather " +
				"effects, such as those caused by extreme heat or cold.",
			"Resilient (exhaustion)": "Once per long rest, when you complete a short rest, you can reduce your exhaustion " +
				"level by one.",
			"Traveller": "You have proficiency in the Survival skill.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
	"salvager": {
		Name:                   "Salvager",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Draconic", "Gnomish"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 115",
		Traits: map[string]string{
			"Repurpose": "You can create Tiny nonmagical items using materials from your surroundings. An item takes " +
				"1 minute to create and can be anything of 25 gp value or less from the Adventuring Gear table. " +
				"When done, it must sit or float on a surface within 5 feet of you. The item is obviously kitbashed, " +
				"and resale value is minimal. After one use, the item becomes nonfunctional.",
			"Tinkerer": "You have proficiency with tinker's tools or one other kind of tool of your choice. When you " +
				"make a check to create, identify, or disarm a magical or nonmagical object, trap, or device, where " +
				"you have a relevant proficiency, double your PB for the roll",
			"Traveller": "You have proficiency in the Survival skill.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
	"slayer": {
		Name:                   "Slayer",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Primordial", "Sylvan"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 115",
		Traits: map[string]string{
			"Natural Predator": "You have proficiency in the Intimidation skill. You have advantage on Intimidation " +
				"checks to influence Beasts and creatures with the Animal tag.",
			"Tracker": "When you make a check to locate, spot, or track a creature, you can add your PB to the roll. " +
				"If you have proficiency in the skill or tool being used, double your PB for the roll.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
	"stone": {
		Name:                   "Stone",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Dwarvish"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 115",
		Traits: map[string]string{
			"Ancestral Arts (proficiency)": "You have proficiency with Construction tools. Double your PB for any ability check " +
				"you make that uses them",
			"Eye for Quality": "When you make an ability check related to the origin or purpose of an object or structure " +
				"made of metal or stone, you can add your PB to the roll. If you have proficiency in a relevant skill " +
				"or tool, double your PB for the roll.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Ancestral Arts (Weapon Proficiency)": {
				NumberToSelect: 1,
				// todo: Implement a list of martial weapons to choose from
				Options: []string{"Sword", "WarHammer"},
			},
		},
	},
	"supplicant": {
		Name:                   "Supplicant",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Draconic", "Giant", "Undercommon"},
		LanguageSuggestionNote: "Typically, the language favored by your current or previous overlord.",
		HeritageSource:         "Players Guide, pg 116",
		Traits: map[string]string{
			"Scurry": "As a bonus action, you can move up to 10 feet without provoking opportunity attacks. " +
				"This movement doesn't trigger traps or hazards that you are aware of, even if they are armed.",
			"Supplicant (Doom)": "When a creature within 30 feet of you spends Doom, you have an advantage on ability " +
				"checks and saves until the beginning of your next turn.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Supplicant (proficiency)": {
				NumberToSelect: 1,
				Options:        []string{"Insight", "Persuasion"},
			},
		},
	},
	"vexed": {
		Name:                   "Vexed",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Abyssal", "Celestial", "Infernal"},
		LanguageSuggestionNote: "Typically, esoteric languages most closely aligned with your pursing power or force.",
		HeritageSource:         "Players Guide, pg 116",
		Traits: map[string]string{
			"Prodigal Disciple": "When you make a save to resist becoming charmed or possessed, you can treat any " +
				"d20 die roll of 9 or lower as a 10.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
			"Quarry's Cunning": {
				NumberToSelect: 1,
				Options:        []string{"Deception", "Insight"},
			},
		},
	},
	"wildlands": {
		Name:                   "Wildlands",
		LanguageDefaults:       []string{"Common"},
		LanguageSuggestions:    []string{"Sylvan"},
		LanguageSuggestionNote: "",
		HeritageSource:         "Players Guide, pg 116",
		Traits: map[string]string{
			"Beast Affinity": "Using gestures and sounds, you can communicate simple ideas to Beasts and creatures " +
				"with the Animal tag, and you have advantage on checks made to interact with such creatures.",
			"Shepherd's Gift (proficiency)": "You have proficiency in the Animal Handling skill.",
			"Shepherd's Gift (melee)": "Any Beast or Creature with the Animal tag whose CR is equal to or less than " +
				"your PB that targets you with an attack must first make a wis (Animal Handling) check. If you " +
				"succeed, the creature must choose a new target or lose the attack.",
		},
		TraitOptions: map[string]ChoiceOptions{
			"Languages": {
				NumberToSelect: 1,
				Options:        LanguageNames([]string{"Common"}),
			},
		},
	},
}
