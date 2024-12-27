package character

var Classes = map[string]Class{
	"barbarian": {
		Name:        "Barbarian",
		Description: "Ferocious warriors who harness primal power.",
		HitDie:      "d12",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"str", "con", "dex", "wis", "cha", "int"},
				KeyAbilities:                []string{"str"},
			},
		},
		SaveProficiencies: []string{"str", "con"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"weapons",
		},
		Subclasses: map[string]Subclass{
			"berserker": {
				Name:        "Berserker",
				Description: "",
			},
			"wild fury": {
				Name:        "Wild Fury",
				Description: "",
			},
		},
	},
	"bard": {
		Name:        "Bard",
		Description: "Skilled performers who inspire allies and wield Arcane magic.",
		HitDie:      "d8",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"cha", "dex", "con", "str", "wis", "int"},
				KeyAbilities:                []string{"cha"},
			},
		},
		SaveProficiencies: []string{"dex", "cha"},
		EquipmentProficiencies: []string{
			"light armor",
			"simple weapons",
			"finesse weapons",
		},
		SpellcastingAbility: Cha,
		Subclasses: map[string]Subclass{
			"lore": {
				Name:        "Lore",
				Description: "",
			},
			"victory": {
				Name:        "Victory",
				Description: "",
			},
		},
	},
	"cleric": {
		Name:        "Cleric",
		Description: "Faithful casters who wield Divine magic.",
		HitDie:      "d8",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"wis", "con", "cha", "str", "dex", "int"},
				KeyAbilities:                []string{"wis"},
			},
		},
		SaveProficiencies: []string{"wis", "cha"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"simple weapons"},
		SpellcastingAbility: Wis,
		Subclasses: map[string]Subclass{
			"life domain": {
				Name:        "Life Domain",
				Description: "",
			},
			"light domain": {
				Name:        "Light Domain",
				Description: "",
			},
			"war domain": {
				Name:        "War Domain",
				Description: "",
			},
		},
	},
	"druid": {
		Name:        "Druid",
		Description: "Guardians of nature who wield Primordial magic.",
		HitDie:      "d8",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"wis", "con", "int", "dex", "cha", "str"},
				KeyAbilities:                []string{"wis"},
			},
		},
		SaveProficiencies: []string{"int", "wis"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"simple weapons"},
		SpellcastingAbility: Wis,
		Subclasses: map[string]Subclass{
			"leaf": {
				Name:        "Leaf",
				Description: "",
			},
			"shifter": {
				Name:        "Shifter",
				Description: "",
			},
		},
	},
	"fighter": {
		Name:        "Fighter",
		Description: "Hardy adventurers who excel in combat.",
		HitDie:      "d10",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"str", "con", "dex", "cha", "wis", "int"},
				KeyAbilities:                []string{"str"},
			},
			"Dexterity": {
				AbilityScoreOrderPreference: []string{"dex", "con", "str", "cha", "wis", "int"},
				KeyAbilities:                []string{"dex"},
			},
		},
		SaveProficiencies: []string{"str", "con"},
		EquipmentProficiencies: []string{
			"armor",
			"shields",
			"weapons"},
		// Implement SpellcastingAbility = Int if subclass = Spell Blade
		Subclasses: map[string]Subclass{
			"spell blade": {
				Name:                "Spell Blade",
				Description:         "",
				SpellcastingAbility: Int,
			},
			"weapon master": {
				Name:        "Weapon Master",
				Description: "",
			},
		},
	},
	"mechanist": {
		Name:        "Mechanist",
		Description: "Crafty engineers who sculpt mystic forces into items.",
		HitDie:      "d10",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"int", "con", "dex", "wis", "cha", "str"},
				KeyAbilities:                []string{"int"},
			},
		},
		SaveProficiencies: []string{"con", "int"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"weapons"},
		// Implement SpellcastingAbility = Int if subclass = Spellwright
		Subclasses: map[string]Subclass{
			"metallurgist": {
				Name:        "Metallurgist",
				Description: "",
			},
			"spellwright": {
				Name:                "Spellwright",
				Description:         "",
				SpellcastingAbility: Int,
			},
		},
	},
	"monk": {
		Name:        "Monk",
		Description: "Martial artists who harness mystical energy.",
		HitDie:      "d8",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"str", "dex", "wis", "con", "int", "cha"},
				KeyAbilities:                []string{"dex", "wis"},
			},
		},
		SaveProficiencies:      []string{"str", "dex"},
		EquipmentProficiencies: []string{"simple weapons", "shortswords"},
		Subclasses: map[string]Subclass{
			"flickering dark": {
				Name:        "Flickering Dark",
				Description: "",
			},
			"open hand": {
				Name:        "Open Hand",
				Description: "",
			},
		},
	},
	"paladin": {
		Name:        "Paladin",
		Description: "Holy warriors who smite foes with Divine power.",
		HitDie:      "d10",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"str", "cha", "con", "wis", "dex", "int"},
				KeyAbilities:                []string{"str", "cha"},
			},
			"Dexterity": {
				AbilityScoreOrderPreference: []string{"dex", "cha", "con", "wis", "str", "int"},
				KeyAbilities:                []string{"dex", "cha"},
			},
		},
		SaveProficiencies:      []string{"wis", "cha"},
		EquipmentProficiencies: []string{"armor", "shields", "weapons"},
		SpellcastingAbility:    Cha,
		Subclasses: map[string]Subclass{
			"devotion": {
				Name:        "Devotion",
				Description: "",
			},
			"justice": {
				Name:        "Justice",
				Description: "",
			},
		},
	},
	"ranger": {
		Name:        "Ranger",
		Description: "Resourceful survivalists with a mystic connection to nature.",
		HitDie:      "d10",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"dex", "wis", "str", "con", "int", "cha"},
				KeyAbilities:                []string{"dex", "wis"},
			},
		},
		SaveProficiencies: []string{"str", "dex"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields", "weapons"},
		SpellcastingAbility: Wis,
		Subclasses: map[string]Subclass{
			"hunter": {
				Name:        "Hunter",
				Description: "",
			},
			"pack master": {
				Name:        "Pack Master",
				Description: "",
			},
		},
	},
	"rogue": {
		Name:        "Rogue",
		Description: "Cunning adventurers who rely on agility and trickery.",
		HitDie:      "d8",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"dex", "cha", "con", "int", "str", "wis"},
				KeyAbilities:                []string{"dex"},
			},
		},
		SaveProficiencies: []string{"dex", "int"},
		EquipmentProficiencies: []string{
			"light armor",
			"simple weapons",
			"finesse weapons"},
		Subclasses: map[string]Subclass{
			"enforcer": {
				Name:        "Enforcer",
				Description: "",
			},
			"thief": {
				Name:        "Thief",
				Description: "",
			},
		},
	},
	"sorcerer": {
		Name:        "Sorcerer",
		Description: "Powerful casters who channel raw Arcane power from within.",
		HitDie:      "d6",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"cha", "con", "dex", "wis", "con", "str"},
				KeyAbilities:                []string{"cha"},
			},
		},
		SaveProficiencies:      []string{"con", "cha"},
		EquipmentProficiencies: []string{"simple weapons"},
		SpellcastingAbility:    Cha,
		Subclasses: map[string]Subclass{
			"chaos": {
				Name:        "Chaos",
				Description: "",
			},
			"draconic": {
				Name:        "Draconic",
				Description: "",
			},
		},
	},
	"warlock": {
		Name:        "Warlock",
		Description: "Supernatural casters who draw magic from Wyrd forces.",
		HitDie:      "d8",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"cha", "wis", "con", "dex", "int", "str"},
				KeyAbilities:                []string{"cha"},
			},
		},
		SaveProficiencies: []string{"wis", "cha"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"simple weapons"},
		SpellcastingAbility: Cha,
		Subclasses: map[string]Subclass{
			"fiend": {
				Name:        "Fiend",
				Description: "",
			},
			"reaper": {
				Name:        "Reaper",
				Description: "",
			},
		},
	},
	"wizard": {
		Name:        "Wizard",
		Description: "Cerebral casters who wield Arcane magic.",
		HitDie:      "d6",
		ClassBuildTypes: map[string]ClassBuildType{
			"Standard": {
				AbilityScoreOrderPreference: []string{"int", "wis", "dex", "con", "str", "cha"},
				KeyAbilities:                []string{"int"},
			},
		},
		SaveProficiencies:      []string{"int", "wis"},
		EquipmentProficiencies: []string{"simple weapons"},
		SpellcastingAbility:    Int,
		Subclasses: map[string]Subclass{
			"battle mage": {
				Name:        "Battle Mage",
				Description: "",
			},
			"cantrip adept": {
				Name:        "Cantrip Adept",
				Description: "",
			},
		},
	},
}
