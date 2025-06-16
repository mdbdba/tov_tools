package character

type ArmorClassCalculator struct {
	BaseAC               int
	AddDexterityModifier bool
	DexterityModifierMax int
}
type ArmorPiece struct {
	Name         string
	Description  string
	Category     string
	CostAmount   int
	CostCoin     string
	ArmorClass   ArmorClassCalculator
	Weight       float64 // in lbs
	Properties   []string
	Prerequisite func(c *Character) bool // A function to check if a character meets the prerequisite
}

var Armor = map[string]ArmorPiece{
	"padded": {
		Name:        "Padded",
		Description: "This full-body outfit consists of quilted layers of cloth and batting.",
		Category:    "Light",
		CostAmount:  5,
		CostCoin:    "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               11,
			AddDexterityModifier: true,
		},
		Weight:     8,
		Properties: []string{"may be Noisy"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"leather": {
		Name: "Leather",
		Description: "The breastplate and shoulder protectors of this armor are made of leather that has " +
			"been stiffened by being boiled in oil. The rest of the armor is made of softer and more " +
			"flexible materials.",
		Category:   "Light",
		CostAmount: 10,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               11,
			AddDexterityModifier: true,
		},
		Weight:     10,
		Properties: []string{"Natural Materials"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},

	"studded_leather": {
		Name:        "Studded leather",
		Description: "Made from tough but flexible leather, studded leather is reinforced with close-set rivets or spikes.",
		Category:    "Light",
		CostAmount:  45,
		CostCoin:    "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               12,
			AddDexterityModifier: true,
		},
		Weight:     13,
		Properties: []string{},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"brigandine": {
		Name:        "Brigandine",
		Description: "This knee-length coat is made of heavy cloth or canvas lined with small metal plates.",
		Category:    "Light",
		CostAmount:  50,
		CostCoin:    "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               13,
			AddDexterityModifier: true,
		},
		Weight:     25,
		Properties: []string{"Noisy"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"hide": {
		Name:        "Hide",
		Description: "This full-body suit of armor consists of thick furs and pelts.",
		Category:    "Medium",
		CostAmount:  10,
		CostCoin:    "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               12,
			AddDexterityModifier: true,
			DexterityModifierMax: 2,
		},
		Weight:     12,
		Properties: []string{"Natural Materials"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"chain_shirt": {
		Name: "Chain Shirt",
		Description: "A chain shirt is made of interlocking metal rings that are worn between layers of clothing " +
			"or leather. This armor protects the wearer’s upper body and the outer layers muffle the sound of the " +
			"rings rubbing against one another.",
		Category:   "Medium",
		CostAmount: 50,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               13,
			AddDexterityModifier: true,
			DexterityModifierMax: 2,
		},
		Weight:     20,
		Properties: []string{},
		Prerequisite: func(c *Character) bool {
			return true
		},
	},
	"scale_mail": {
		Name: "Scale mail",
		Description: "This armor consists of a coat and leggings (and perhaps a separate skirt) of leather covered " +
			"with overlapping pieces of metal, much like the scales of a fish.",
		Category:   "Medium",
		CostAmount: 50,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               14,
			AddDexterityModifier: true,
			DexterityModifierMax: 2,
		},
		Weight:     45,
		Properties: []string{"Noisy"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"breastplate": {
		Name: "Breastplate",
		Description: "This armor consists of a fitted metal chest lined with supple leather. This armor leaves " +
			"limbs unprotected but provides good protection for vital organs and allows for easier movement than " +
			"most medium armor.",
		Category:   "Medium",
		CostAmount: 400,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               14,
			AddDexterityModifier: true,
			DexterityModifierMax: 2,
		},
		Weight:     20,
		Properties: []string{},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"half_plate": {
		Name: "Half plate",
		Description: "Half plate consists of shaped metal plates that cover most of the wearer's body. It doesn’t " +
			"include leg protection beyond greaves attached with leather straps.",
		Category:   "Medium",
		CostAmount: 750,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               15,
			AddDexterityModifier: true,
			DexterityModifierMax: 2,
		},
		Weight:     40,
		Properties: []string{"Noisy"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"ring_mail": {
		Name:        "Ring mail",
		Description: "This leather armor has heavy rings sewn into it. The rings help reinforce the armor against attacks.",
		Category:    "Heavy",
		CostAmount:  30,
		CostCoin:    "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               15,
			AddDexterityModifier: false,
		},
		Weight:     40,
		Properties: []string{"Noisy"},
		Prerequisite: func(c *Character) bool {
			// no prerequisite for combat casting
			return true
		},
	},
	"chain_mail": {
		Name: "Chain mail",
		Description: "This is a head-to-toe suit of armor made of quilted fabric worn underneath the mail to prevent " +
			"chafing and to cushion the impact of blows.",
		Category:   "Heavy",
		CostAmount: 75,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               16,
			AddDexterityModifier: false,
		},
		Weight:     55,
		Properties: []string{"Cumbersome (STR 13)", "Noisy"},
		Prerequisite: func(c *Character) bool {
			// str 13 or higher
			if c.Abilities.Values["str"] >= 13 {
				return true
			}
			return false
		},
	},
	"splint": {
		Name: "Splint",
		Description: "This armor is made of narrow vertical strips of metal riveted to a backing of leather that is " +
			"worn over cloth padding. Flexible chain mail protects the joints.",
		Category:   "Heavy",
		CostAmount: 200,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               17,
			AddDexterityModifier: false,
		},
		Weight:     60,
		Properties: []string{"Cumbersome (STR 15)", "Noisy"},
		Prerequisite: func(c *Character) bool {
			// str 15 or higher
			if c.Abilities.Values["str"] >= 15 {
				return true
			}
			return false
		},
	},
	"plate": {
		Name: "Plate",
		Description: "Plate consists of shaped, interlocking metal plates to cover the entire body. A suit of plate " +
			"includes gauntlets, heavy leather boots, a visored helmet, and thick layers of padding underneath the " +
			"armor. Buckles and straps distribute the weight over the body.",
		Category:   "Heavy",
		CostAmount: 1500,
		CostCoin:   "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               18,
			AddDexterityModifier: false,
		},
		Weight:     65,
		Properties: []string{"Cumbersome (STR 16)", "Noisy"},
		Prerequisite: func(c *Character) bool {
			// str 16 or higher
			if c.Abilities.Values["str"] >= 16 {
				return true
			}
			return false
		},
	},
	"shield": {
		Name:        "Shield",
		Description: "This broad piece of wood and metal is held by a handle attached to one side.",
		Category:    "Shield",
		CostAmount:  10,
		CostCoin:    "gp",
		ArmorClass: ArmorClassCalculator{
			BaseAC:               2,
			AddDexterityModifier: false,
		},
		Weight:     6,
		Properties: []string{},
	},
}
