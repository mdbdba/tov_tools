package static_data

// Weapon Option Saves. If an option requires a creature
// to make an ability check or save, the DC equals
// 8 + the attacker’s PB + the attacker’s STR or DEX modifier (attacker’s choice).

// Unless specified otherwise, a weapon attack used to perform a weapon option has only the option’s listed effect
// and doesn’t deal normal weapon damage. Weapon options can be used only when a wielder takes the Attack action
// on their turn, unless a feature like the fighter’s Martial Action allows a weapon option attack to be performed
// as a bonus action. Characters with the Multiattack feature can perform a weapon option in place of one of
// the attacks granted by Multiattack.
//

var WeaponOptions = map[string]string{
	"Bash": "Make an attack roll with this weapon. On a hit, the target has disadvantage on its next attack roll.",
	"Disarm": "Make an attack roll with this weapon. On a hit, the target must succeed on a STR or DEX save " +
		"(target’s choice) or drop a weapon, shield, or object it is wielding. The dropped item lands in an unoccupied " +
		"space within 5 feet of the target. If no unoccupied space is within range, the item lands at the target’s feet.",
	"Hamstring": "Make an attack roll with this weapon. On a hit, the target’s base movement speed is reduced by " +
		"10 feet for 1 minute. A creature’s speed can’t be reduced by more than 10 feet with this weapon option. A " +
		"creature within 5 feet of the target can take an action to tend the wound with a successful WIS (Medicine) " +
		"check (against your weapon option DC), ending the effect on the target. The effect also ends if the target " +
		"receives any magical healing.",
	"Pinning Shot": "Make an attack roll with this weapon against a Large or smaller creature. On a hit, the target " +
		"must succeed on a STR or DEX save (target’s choice) or its speed becomes 0 feet until the end of its next " +
		"turn. A creature, including the target, can use its action to attempt to free the target with a STR " +
		"(Athletics) check or a DEX (Acrobatics) check (the creature’s choice) versus the attacker’s weapon option " +
		"DC. On a success, the target is freed and can move as normal. A target must make only one check to free " +
		"itself, using the highest DC of characters performing this weapon option, regardless of the number of " +
		"arrows or bolts holding it in place.",
	"Pull": "Make an attack roll with this weapon against a Large or smaller creature. On a hit, the target is " +
		"pulled up to 5 feet closer to you. If this movement would pull a creature into damaging terrain, such as " +
		"lava or a pit, it can make a STR or DEX (target’s choice) save to avoid the pull on a success.",
	"Ricochet Shot": "Make an attack roll with this weapon against a target you can see that has half or three-quarters " +
		"cover. Your chosen target must be within 10 feet of another object or structure that isn’t the same item " +
		"providing it with cover. When you do so, you can treat the target’s AC as if it wasn’t behind cover. If the " +
		"attack is successful, the target takes damage from the attack as it would with a standard weapon attack. " +
		"This weapon option expends the same ammunition as a normal attack with this weapon.",
	"Trip": "Make an attack roll with this weapon against a Large or smaller creature. On a hit, the target " +
		"must succeed on a STR or DEX save (target’s choice) or fall prone. If the target is mounted, it has " +
		"advantage on the save.",
}

type WeaponRange struct {
	Min int // distance in feet
	Max int // distance in feet
}
type WeaponDamage struct {
	Identifier  string // One-Handed, Two-Handed, Thrown
	Range       WeaponRange
	TimesToRoll int
	Sides       int
	DamageType  string
}

type Weapon struct {
	Name       string // Display name
	Category   string // Simple Melee, Simple Ranged, Martial Melee, Martial Ranged
	CostAmount int    // Cost amount
	CostCoin   string // gp, sp, cp
	Damage     map[string]WeaponDamage
	Weight     float64 // weight in lbs
	Options    []string
	Properties []string
}

var Weapons = map[string]Weapon{
	"club": {
		Name:       "Club",
		Category:   "Simple Melee",
		CostAmount: 1,
		CostCoin:   "sp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     2,
		Options:    []string{"Bash"},
		Properties: []string{"Light"},
	},
	"dagger": {
		Name:       "Dagger",
		Category:   "Simple Melee",
		CostAmount: 2,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 20, Max: 60},
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "piercing",
			},
		},
		Weight:     1,
		Options:    []string{"Pinning Shot"},
		Properties: []string{"Finesse", "Light", "Thrown"},
	},

	"greatclub": {
		Name:       "Greatclub",
		Category:   "Simple Melee",
		CostAmount: 2,
		CostCoin:   "sp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     10,
		Options:    []string{"Bash"},
		Properties: []string{"Two-handed"},
	},

	"handaxe": {
		Name:       "Handaxe",
		Category:   "Simple Melee",
		CostAmount: 5,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 20, Max: 60},
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "slashing",
			},
		},
		Weight:     2,
		Options:    []string{"Hamstring"},
		Properties: []string{"Light", "Thrown (range 20/60 ft.)"},
	},

	"javelin": {
		Name:       "Javelin",
		Category:   "Simple Melee",
		CostAmount: 5,
		CostCoin:   "sp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 30, Max: 120},
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "piercing",
			},
		},
		Weight:     2,
		Options:    []string{"Pinning Shot"},
		Properties: []string{"Thrown (range 30/120 ft.)"},
	},

	"light_hammer": {
		Name:       "Light hammer",
		Category:   "Simple Melee",
		CostAmount: 2,
		CostCoin:   "sp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 20, Max: 60},
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     2,
		Options:    []string{"Bash"},
		Properties: []string{"Light", "Thrown (range 20/60 ft.)"},
	},

	"mace": {
		Name:       "Mace",
		Category:   "Simple Melee",
		CostAmount: 5,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     4,
		Options:    []string{"Bash"},
		Properties: []string{},
	},

	"quarterstaff": {
		Name:       "Quarterstaff",
		Category:   "Simple Melee",
		CostAmount: 2,
		CostCoin:   "sp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "bludgeoning",
			},
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     4,
		Options:    []string{"Bash"},
		Properties: []string{"Versatile"},
	},

	"sickle": {
		Name:       "Sickle",
		Category:   "Simple Melee",
		CostAmount: 1,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "slashing",
			},
		},
		Weight:     2,
		Options:    []string{"Hamstring"},
		Properties: []string{"Light"},
	},
	"spear": {
		Name:       "Spear",
		Category:   "Simple Melee",
		CostAmount: 1,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 20, Max: 60},
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "piercing",
			},
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     3,
		Options:    []string{"Pull", "Trip"},
		Properties: []string{"Thrown (range 20/60 ft.)", "Versatile"},
	},

	"light_crossbow": {
		Name:       "Crossbow, light",
		Category:   "Simple Ranged",
		CostAmount: 25,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				Range:       WeaponRange{Min: 80, Max: 320},
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     5,
		Options:    []string{},
		Properties: []string{"Ammunition (range 80/320 ft.)", "Loading", "Two-handed"},
	},

	"dart": {
		Name:       "Dart",
		Category:   "Simple Ranged",
		CostAmount: 5,
		CostCoin:   "cp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 20, Max: 60},
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "piercing",
			},
		},
		Weight:     0.25,
		Options:    []string{},
		Properties: []string{"Finesse", "Thrown (range 20/60 ft.)"},
	},

	"shortbow": {
		Name:       "Shortbow",
		Category:   "Simple Ranged",
		CostAmount: 25,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				Range:       WeaponRange{Min: 80, Max: 320},
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "piercing",
			},
		},
		Weight:     2,
		Options:    []string{},
		Properties: []string{"Ammunition (range 80/320 ft.)", "Two-handed"},
	},

	"sling": {
		Name:       "Sling",
		Category:   "Simple Ranged",
		CostAmount: 1,
		CostCoin:   "sp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 30, Max: 120},
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     0,
		Options:    []string{"Ricochet Shot"},
		Properties: []string{"Ammunition (range 30/120 ft.)"},
	},

	"battleaxe": {
		Name:       "Battleaxe",
		Category:   "Martial Melee",
		CostAmount: 10,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "slashing",
			},
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "slashing",
			},
		},
		Weight:     4,
		Options:    []string{"Disarm", "Hamstring"},
		Properties: []string{"Versatile"},
	},

	"flail": {
		Name:       "Flail",
		Category:   "Martial Melee",
		CostAmount: 10,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     2,
		Options:    []string{"Bash", "Disarm"},
		Properties: []string{},
	},

	"glaive": {
		Name:       "Glaive",
		Category:   "Martial Melee",
		CostAmount: 20,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "slashing",
			},
		},
		Weight:     6,
		Options:    []string{"Trip"},
		Properties: []string{"Heavy", "Reach", "Two-handed"},
	},

	"greataxe": {
		Name:       "Greataxe",
		Category:   "Martial Melee",
		CostAmount: 30,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       12,
				DamageType:  "slashing",
			},
		},
		Weight:     7,
		Options:    []string{"Disarm", "Hamstring"},
		Properties: []string{"Heavy", "Two-handed"},
	},

	"greatsword": {
		Name:       "Greatsword",
		Category:   "Martial Melee",
		CostAmount: 50,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 2,
				Sides:       6,
				DamageType:  "slashing",
			},
		},
		Weight:     6,
		Options:    []string{"Disarm", "Hamstring"},
		Properties: []string{"Heavy", "Two-handed"},
	},

	"halberd": {
		Name:       "Halberd",
		Category:   "Martial Melee",
		CostAmount: 20,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "slashing",
			},
		},
		Weight:     6,
		Options:    []string{"Trip"},
		Properties: []string{"Heavy", "Reach", "Two-handed"},
	},

	"lance": {
		Name:       "Lance",
		Category:   "Martial Melee",
		CostAmount: 10,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       12,
				DamageType:  "piercing",
			},
		},
		Weight:     6,
		Options:    []string{},
		Properties: []string{"Reach", "Special"},
	},

	"longsword": {
		Name:       "Longsword",
		Category:   "Martial Melee",
		CostAmount: 15,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "slashing",
			},
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "slashing",
			},
		},
		Weight:     3,
		Options:    []string{"Disarm", "Hamstring"},
		Properties: []string{"Versatile"},
	},

	"maul": {
		Name:       "Maul",
		Category:   "Martial Melee",
		CostAmount: 10,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 2,
				Sides:       6,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     10,
		Options:    []string{"Bash", "Disarm"},
		Properties: []string{"Heavy", "Two-handed"},
	},

	"morningstar": {
		Name:       "Morningstar",
		Category:   "Martial Melee",
		CostAmount: 15,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     4,
		Options:    []string{"Disarm"},
		Properties: []string{},
	},

	"pike": {
		Name:       "Pike",
		Category:   "Martial Melee",
		CostAmount: 5,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "piercing",
			},
		},
		Weight:     18,
		Options:    []string{"Trip"},
		Properties: []string{"Heavy", "Reach", "Two-handed"},
	},

	"rapier": {
		Name:       "Rapier",
		Category:   "Martial Melee",
		CostAmount: 25,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     2,
		Options:    []string{"Disarm"},
		Properties: []string{"Finesse"},
	},

	"scimitar": {
		Name:       "Scimitar",
		Category:   "Martial Melee",
		CostAmount: 25,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "slashing",
			},
		},
		Weight:     3,
		Options:    []string{"Hamstring"},
		Properties: []string{"Finesse", "Light"},
	},

	"scythe": {
		Name:       "Scythe",
		Category:   "Martial Melee",
		CostAmount: 20,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 2,
				Sides:       4,
				DamageType:  "slashing",
			},
		},
		Weight:     4,
		Options:    []string{"Pull", "Trip"},
		Properties: []string{"Reach", "Two-handed"},
	},

	"shortsword": {
		Name:       "Shortsword",
		Category:   "Martial Melee",
		CostAmount: 10,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "piercing",
			},
		},
		Weight:     2,
		Options:    []string{"Disarm"},
		Properties: []string{"Finesse", "Light"},
	},

	"trident": {
		Name:       "Trident",
		Category:   "Martial Melee",
		CostAmount: 5,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "piercing",
			},
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     4,
		Options:    []string{"Disarm"},
		Properties: []string{"Thrown (range 20/60 ft.)", "Versatile"},
	},

	"war_pick": {
		Name:       "War pick",
		Category:   "Martial Melee",
		CostAmount: 5,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     2,
		Options:    []string{"Disarm"},
		Properties: []string{},
	},

	"warhammer": {
		Name:       "Warhammer",
		Category:   "Martial Melee",
		CostAmount: 15,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "bludgeoning",
			},
			"two-handed": {
				Identifier:  "Two-Handed",
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "bludgeoning",
			},
		},
		Weight:     2,
		Options:    []string{"Bash", "Disarm"},
		Properties: []string{"Versatile"},
	},

	"whip": {
		Name:       "Whip",
		Category:   "Martial Melee",
		CostAmount: 2,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				TimesToRoll: 1,
				Sides:       4,
				DamageType:  "slashing",
			},
		},
		Weight:     3,
		Options:    []string{"Pull", "Trip"},
		Properties: []string{"Finesse", "Reach"},
	},

	"blowgun": {
		Name:       "Blowgun",
		Category:   "Martial Ranged",
		CostAmount: 10,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 25, Max: 100},
				TimesToRoll: 1,
				Sides:       1,
				DamageType:  "piercing",
			},
		},
		Weight:     1,
		Options:    []string{},
		Properties: []string{"Ammunition (range 25/100 ft.)", "Loading"},
	},

	"hand_crossbow": {
		Name:       "Crossbow, hand",
		Category:   "Martial Ranged",
		CostAmount: 75,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"one-handed": {
				Identifier:  "One-Handed",
				Range:       WeaponRange{Min: 30, Max: 120},
				TimesToRoll: 1,
				Sides:       6,
				DamageType:  "piercing",
			},
		},
		Weight:     3,
		Options:    []string{"Pinning Shot"},
		Properties: []string{"Ammunition (range 30/120 ft.)", "Light", "Loading"},
	},

	"heavy_crossbow": {
		Name:       "Crossbow, heavy",
		Category:   "Martial Ranged",
		CostAmount: 50,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				Range:       WeaponRange{Min: 100, Max: 400},
				TimesToRoll: 1,
				Sides:       10,
				DamageType:  "piercing",
			},
		},
		Weight:     18,
		Options:    []string{"Pinning Shot"},
		Properties: []string{"Ammunition (range 100/400 ft.)", "Heavy", "Loading", "Two-handed"},
	},

	"longbow": {
		Name:       "Longbow",
		Category:   "Martial Ranged",
		CostAmount: 50,
		CostCoin:   "gp",
		Damage: map[string]WeaponDamage{
			"two-handed": {
				Identifier:  "Two-Handed",
				Range:       WeaponRange{Min: 150, Max: 600},
				TimesToRoll: 1,
				Sides:       8,
				DamageType:  "piercing",
			},
		},
		Weight:     2,
		Options:    []string{"Pinning Shot"},
		Properties: []string{"Ammunition (range 150/600 ft.)", "Heavy", "Two-handed"},
	},
}
