package character

var Classes = map[string]Class{
	"barbarian": {
		Name:                        "Barbarian",
		Description:                 "Ferocious warriors who harness primal power.",
		HitDie:                      "d12",
		AbilityScoreOrderPreference: []string{"str", "con", "dex", "wis", "cha", "int"},
		KeyAbilities:                [][]string{{"str"}},
		SaveProficiencies:           []string{"str", "con"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"weapons",
		},
	},
	"bard": {
		Name:                        "Bard",
		Description:                 "Skilled performers who inspire allies and wield Arcane magic.",
		HitDie:                      "d8",
		AbilityScoreOrderPreference: []string{"cha", "dex", "con", "str", "wis", "int"},
		KeyAbilities:                [][]string{{"cha"}},
		SaveProficiencies:           []string{"dex", "cha"},
		EquipmentProficiencies: []string{
			"light armor",
			"simple weapons",
			"finesse weapons",
		},
	},
	"cleric": {
		Name:                        "Cleric",
		Description:                 "Faithful casters who wield Divine magic.",
		HitDie:                      "d8",
		AbilityScoreOrderPreference: []string{"wis", "con", "cha", "str", "dex", "int"},
		KeyAbilities:                [][]string{{"wis"}},
		SaveProficiencies:           []string{"wis", "cha"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"simple weapons"},
	},
	"druid": {
		Name:                        "Druid",
		Description:                 "Guardians of nature who wield Primordial magic.",
		HitDie:                      "d8",
		AbilityScoreOrderPreference: []string{"wis", "con", "int", "dex", "cha", "str"},
		KeyAbilities:                [][]string{{"wis"}},
		SaveProficiencies:           []string{"int", "wis"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"simple weapons"},
	},
	"fighter": {
		Name:                        "Fighter",
		Description:                 "Hardy adventurers who excel in combat.",
		HitDie:                      "d10",
		AbilityScoreOrderPreference: []string{"str", "con", "dex", "cha", "wis", "int"},
		KeyAbilities:                [][]string{{"str"}, {"dex"}},
		SaveProficiencies:           []string{"str", "con"},
		EquipmentProficiencies: []string{
			"armor",
			"shields",
			"weapons"},
	},
	"mechanist": {
		Name:                        "Mechanist",
		Description:                 "Crafty engineers who sculpt mystic forces into items.",
		HitDie:                      "d10",
		AbilityScoreOrderPreference: []string{"int", "con", "dex", "wis", "cha", "str"},
		KeyAbilities:                [][]string{{"int"}},
		SaveProficiencies:           []string{"con", "int"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"weapons"},
	},
	"monk": {
		Name:                        "Monk",
		Description:                 "Martial artists who harness mystical energy.",
		HitDie:                      "d8",
		AbilityScoreOrderPreference: []string{"dex", "wis", "con", "str", "int", "cha"},
		KeyAbilities:                [][]string{{"dex", "wis"}},
		SaveProficiencies:           []string{"str", "dex"},
		EquipmentProficiencies:      []string{"simple weapons", "shortswords"},
	},
	"paladin": {
		Name:                        "Paladin",
		Description:                 "Holy warriors who smite foes with Divine power.",
		HitDie:                      "d10",
		AbilityScoreOrderPreference: []string{"str", "cha", "con", "wis", "dex", "int"},
		KeyAbilities:                [][]string{{"str", "cha"}, {"dex", "cha"}},
		SaveProficiencies:           []string{"wis", "cha"},
		EquipmentProficiencies:      []string{"armor", "shields", "weapons"},
	},
	"ranger": {
		Name:                        "Ranger",
		Description:                 "Resourceful survivalists with a mystic connection to nature.",
		HitDie:                      "d10",
		AbilityScoreOrderPreference: []string{"dex", "wis", "con", "str", "int", "cha"},
		KeyAbilities:                [][]string{{"dex", "wis"}},
		SaveProficiencies:           []string{"str", "dex"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields", "weapons"},
	},
	"rogue": {
		Name:                        "Rogue",
		Description:                 "Cunning adventurers who rely on agility and trickery.",
		HitDie:                      "d8",
		AbilityScoreOrderPreference: []string{"dex", "cha", "con", "int", "str", "wis"},
		KeyAbilities:                [][]string{{"dex"}},
		SaveProficiencies:           []string{"dex", "int"},
		EquipmentProficiencies: []string{
			"light armor",
			"simple weapons",
			"finesse weapons"},
	},
	"sorcerer": {
		Name:                        "Sorcerer",
		Description:                 "Powerful casters who channel raw Arcane power from within.",
		HitDie:                      "d6",
		AbilityScoreOrderPreference: []string{"cha", "con", "dex", "wis", "con", "str"},
		KeyAbilities:                [][]string{{"cha"}},
		SaveProficiencies:           []string{"con", "cha"},
		EquipmentProficiencies:      []string{"simple weapons"},
	},
	"warlock": {
		Name:        "Warlock",
		Description: "Supernatural casters who draw magic from Wyrd forces.",
		HitDie:      "d8",
		// 0 = "str",
		// 1 = "dex",
		// 2 = "con",
		// 3 = "int",
		// 4 = "wis",
		// 5 = "cha",
		AbilityScoreOrderPreference: []string{"cha", "wis", "con", "dex", "int", "str"},
		KeyAbilities:                [][]string{{"cha"}},
		SaveProficiencies:           []string{"wis", "cha"},
		EquipmentProficiencies: []string{
			"light armor",
			"medium armor",
			"shields",
			"simple weapons"},
	},
	"wizard": {
		Name:                        "Wizard",
		Description:                 "Cerebral casters who wield Arcane magic.",
		HitDie:                      "d6",
		AbilityScoreOrderPreference: []string{"int", "wis", "dex", "con", "str", "cha"},
		KeyAbilities:                [][]string{{"int"}},
		SaveProficiencies:           []string{"int", "wis"},
		EquipmentProficiencies:      []string{"simple weapons"},
	},
}
