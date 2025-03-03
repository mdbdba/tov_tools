package character

var Talents = map[string]Talent{
	"arcanist": {
		Name:     "Arcanist",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Spellcasting Class Feature
			return true
		},
	},
	"combat casting": {
		Name:     "Combat Casting",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"elemental savant": {
		Name:     "Elemental Savant",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Ability to cast at least one spell that deals damage
			return true
		},
	},
	"focus (death)": {
		Name:     "Focus (Death)",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Access to 2nd circle spell slots
			return true
		},
	},
	"focus (creation)": {
		Name:     "Focus (Creation)",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Access to 2nd circle spell slots
			return true
		},
	},
	"focus (fey)": {
		Name:     "Focus (Fey)",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Access to 2nd circle spell slots
			return true
		},
	},
	"focus (war)": {
		Name:     "Focus (War)",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Access to 2nd circle spell slots
			return true
		},
	},
	"mental fortitude": {
		Name:     "Mental Fortitude",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for mental fortitude
			return true
		},
	},
	"psycanist": {
		Name:     "Psycanist",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// int 13 or higher
			if c.Abilities.Values["int"] >= 13 {
				return true
			}
			return false
		},
	},
	"ritualist": {
		Name:     "Ritualist",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// spellcasting class feature
			return true
		},
	},
	"school specialization": {
		Name:     "School Specialization",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for school specialization
			return true
		},
	},
	"spell duelist": {
		Name:     "Spell Duelist",
		Category: "magic",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// Ability to cast one or more cantrips
			return true
		},
	},
	"athletic": {
		Name:     "Athletic",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for athletic
			return true
		},
	},
	"armor expert": {
		Name:     "Armor Expert",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// str 13 or higher
			if c.Abilities.Values["str"] >= 13 {
				return true
			}
			return false
		},
	},
	"armor training": {
		Name:     "Armor Training",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// proficiency with light or medium armor
			return true
		},
	},
	"artillerist": {
		Name:     "Artillerist",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// str 13 or higher
			if c.Abilities.Values["str"] >= 13 {
				return true
			}
			return false
		},
	},
	"combat conditioning": {
		Name:     "Combat Conditioning",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat conditioning
			return true
		},
	},
	"critical training": {
		Name:     "Critical Training",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for critical training
			return true
		},
	},
	"furious charge": {
		Name:     "Furious Charge",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for furious charge
			return true
		},
	},
	"hand to hand": {
		Name:     "Hand to Hand",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for hand to hand
			return true
		},
	},
	"heavy weapon mastery": {
		Name:     "Heavy Weapon Mastery",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// character 4th level or higher
			if c.Level >= 4 {
				return true
			}
			return false
		},
	},
	"opportunist": {
		Name:     "Opportunist",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for opportunist
			return true
		},
	},
	"physical fortitude": {
		Name:     "Physical Fortitude",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for physical fortitude
			return true
		},
	},
	"ranged weapon mastery": {
		Name:     "Ranged Weapon Mastery",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// character 4th level or higher
			if c.Level >= 4 {
				return true
			}
			return false
		},
	},
	"return fire": {
		Name:     "Return Fire",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for return fire
			return true
		},
	},
	"shield mastery": {
		Name:     "Shield Mastery",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// character 4th level or higher
			if c.Level >= 4 {
				return true
			}
			return false
		},
	},
	"spell hunter": {
		Name:     "Spell Hunter",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for spell hunter
			return true
		},
	},
	"two weapon mastery": {
		Name:     "Two Weapon Mastery",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// character 4th level or higher
			if c.Level >= 4 {
				return true
			}
			return false
		},
	},
	"vanguard": {
		Name:     "Vanguard",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for vanguard
			return true
		},
	},
	"weapon discipline": {
		Name:     "Weapon Discipline",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// proficiency with at least on marital weapon
			return true
		},
	},
	"wrestling mastery": {
		Name:     "Wrestling Mastery",
		Category: "martial",
		Prerequisite: func(c *Character) bool {
			// str 15 or higher, level 4 or higher
			if c.Abilities.Values["str"] >= 15 && c.Level >= 4 {
				return true
			}
			return false
		},
	},
	"aware": {
		Name:     "Aware",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for aware
			return true
		},
	},
	"bottomless luck": {
		Name:     "Bottomless Luck",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for bottomless luck
			return true
		},
	},
	"comrade": {
		Name:     "Comrade",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for comrade
			return true
		},
	},
	"covert": {
		Name:     "Covert",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			_, exists := c.SkillProficiencies["stealth"]
			if exists && c.Abilities.Values["dex"] >= 13 {
				return true
			}
			return false
		},
	},
	"dungeoneer": {
		Name:     "Dungeoneer",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for dungeoneer
			return true
		},
	},
	"far traveler": {
		Name:     "Far Traveler",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for far traveler
			return true
		},
	},
	"field medic": {
		Name:     "Field Medic",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for field medic
			return true
		},
	},
	"hard target": {
		Name:     "Hard Target",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for hard target
			return true
		},
	},
	"noxious apothecary": {
		Name:     "Noxious Apothecary",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// TODO: implement Prerequisite
			// int 13 or higher or proficiency with herbalism tools
			return true
		},
	},
	"polyglot": {
		Name:     "Polyglot",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for polyglot
			return true
		},
	},
	"quick": {
		Name:     "Quick",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for quick
			return true
		},
	},
	"scrutinous": {
		Name:     "Scrutinous",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for scrutinous
			return true
		},
	},
	"trade skills": {
		Name:     "Trade Skills",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for trade skills
			return true
		},
	},
	"touch of luck": {
		Name:     "Touch of Luck",
		Category: "technical",
		Prerequisite: func(c *Character) bool {
			// no prerequisite for touch of luck
			return true
		},
	},
}
