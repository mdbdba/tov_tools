package static_data

import "sort"

type ExampleTask struct {
	Description string
	DC          int
}
type Tool struct {
	Name                string
	Description         string
	AssociatedAbilities []string
	Components          []string
	CraftItems          string
	SpecialUses         map[string]string
	ExampleTasks        []ExampleTask
}

func GetToolsetNames() []string {
	var return_value []string
	for key := range Tools {
		return_value = append(return_value, key)
	}
	sort.Strings(return_value)
	return return_value
}

var Tools = map[string]Tool{
	"alchemist tools": {
		Name: "Alchemist Tools",
		Description: "Alchemist tools are used to perform a variety of tasks such as identifying potions or " +
			"foreign substances, mixing reagents together to cause alchemical reactions, and crafting alchemical " +
			"substances.",
		AssociatedAbilities: []string{"dex", "int"},
		Components: []string{"mixing and storage vessels", "an alembic still",
			"metal frame for holding a mixing vessel above a flame",
			"mortar and pestle",
			"variety of common alchemical reagents",
		},
		CraftItems: "Alchemist tools can be used to make items like those listed in the Alchemical " +
			"Concoctions section of the Adventuring Gear table, with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Purify water for safe consumption", DC: 10},
			{Description: "Start a volatile chemical reaction (such as smoke or fire) or neutralize a toxic substance " +
				"(such as an acid or base)", DC: 15},
			{Description: "Identify a rare or obscure poison", DC: 20},
		},
	},
	"artist tools": {
		Name:                "Artist Tools",
		Description:         "Artist tools are used to create illustrated art objects and ornamental documents, pen illuminated texts, and skillfully paint objects.",
		AssociatedAbilities: []string{"dex", "cha"},
		Components: []string{
			"brushes", "spades", "stirrers",
			"inks", "paints", "a square and triangle",
			"a canvas stretcher and an easel",
			"rolls of vellum, parchment, and canvas",
		},
		CraftItems: "Artist tools can be used to make art objects with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Sketch an accurate image of person, place, or symbol you've seen", DC: 10},
			{Description: "Quickly paint a complex glyph or symbol", DC: 15},
			{Description: "Replicate someone else's art style", DC: 20},
		},
	},
	"charlatan tools": {
		Name:                "Charlatan Tools",
		Description:         "Charlatan tools are used to craft disguises and forge documents.",
		AssociatedAbilities: []string{"dex", "cha"},
		Components: []string{"hair dyes", "small props", "cosmetics",
			"a few choice articles of clothing",
			"selection of wax seals", "gold and silver leaf", "inks",
			"variety of parchments", "sculpting tools to fashion melted wax into seals or prosthetics.",
		},
		SpecialUses: map[string]string{
			"Create Disguise": "Over the course of an hour (which can be done as part of a rest) you can create " +
				"a disguise to make your or someone else look like a different creature. A disguise can make a " +
				"creature seem slightly taller, shorter, fatter, or thinner, but can't conceal true size or basic " +
				"arrangement of limbs. It takes 1 minute to don such a disguise once created. A set of charlatan " +
				"tools only has enough material for one disguise at a time, but you can disassemble a disguise to " +
				"create a new disguise as part of the time spent making a disguise. A creature can use its action to " +
				"inspect a disguised appearance, and must succeed on an int contest versus your dex to discern a disguise.",
			"Forge Documents": "Over the course of one hour (which can be done as part of a rest) you can " +
				"create a simple false document of 5 pages or less - like a personal letter, a party " +
				"invitation, or a sales permit. With the GM's permission you can create longer or more " +
				"complicated documents over the course of several hours or days. A creature can use its " +
				"action to inspect a forged document, and must succeed on an int contest versus your dex " +
				"check to discern a forgery.",
		},
		ExampleTasks: []ExampleTask{
			{Description: "Hide a noteworthy physical feature or minor injury", DC: 10},
			{Description: "Forge a signature from memory", DC: 15},
			{Description: "Make yourself look like a well-known celebrity", DC: 20},
		},
	},
	"clothier tools": {
		Name:                "Clothier Tools",
		Description:         "Clothier tools are used to sew, make, or tailor fabric objects such as clothing or shoes.",
		AssociatedAbilities: []string{"dex", "cha"},
		Components: []string{
			"pins and needles",
			"specially sharpened shears",
			"a small hammer and shoe horn",
			"a variety of dyed threads on spools",
			"thicker waxed threads and yarns",
		},
		CraftItems: "Clothier tools can be used to make items with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Determine a garment's age and origin", DC: 10},
			{Description: "Sew an injury closed", DC: 15},
			{Description: "Patch a torn sail during a raging storm", DC: 20},
		},
	},

	"construction tools": {
		Name:                "Construction Tools",
		Description:         "Construction tools are used to build and repair items primarily made of wood or stone.",
		AssociatedAbilities: []string{"str", "dex"},
		Components: []string{
			"hammers and a mallet",
			"a variety of fasteners",
			"a square and triangle",
			"chisels", "a saw",
			"brushes", "spades",
		},
		CraftItems: "Construction tools can be used to make items with the crafting downtime activity.",
		SpecialUses: map[string]string{
			"Fortify":     "You can spend 1 minute fortifying a chest, door, window, or other sealable object or structure. When you do so, the object's AC increases by an amount equal to your PB or creatures have disadvantage on STR (Athletics) checks made to penetrate the fortification (the GM decides which is more relevant). An item doesn't receive additional benefits from multiple attempts to fortify it.",
			"Build Cover": "You can spend 1 hour (which can be done as part of a long rest) erecting a 5-foot by 5-foot wall panel in an unoccupied space, provided you have the raw materials to do so. The panel provides three-quarters cover for one Medium or smaller creature, and it can't be moved. A panel of cover created with this feature is an object with an AC equal to 10 + your PB and 15 HP if made of wood or 25 HP if made of stone. It is immune to psychic and poison damage.",
		},
		ExampleTasks: []ExampleTask{
			{Description: "Drive a spike into a stone (or similarly hard) surface", DC: 10},
			{Description: "Break down and repurpose materials from a woodshed", DC: 15},
			{Description: "Construct a temporary brace to keep a roof from collapsing", DC: 20},
		},
	},

	"gaming set": {
		Name:                "Gaming Set",
		Description:         "Gaming sets are used to gamble, entertain, or engage in fortune-telling. There are three types: cards, dice, and board games. When you have proficiency in a gaming set, choose one type for your proficiency, not all gaming sets.",
		AssociatedAbilities: []string{"dex", "cha"},
		Components: []string{
			"a full set of cards, a full set of dice with a dice cup, or a collection of special pieces used to play a particular game",
		},
		ExampleTasks: []ExampleTask{
			{Description: "Distract a bored sentry", DC: 10},
			{Description: "Perform a card trick that looks like real magic", DC: 15},
			{Description: "Cheat against an experienced opponent", DC: 20},
		},
	},

	"herbalist tools": {
		Name:                "Herbalist Tools",
		Description:         "Herbalist tools are used to store and brew plants and other organic materials into potions, antidotes, and poisons.",
		AssociatedAbilities: []string{"int", "wis"},
		Components: []string{
			"glass vials",
			"a mortar and pestle",
			"pouches for ingredient storage",
			"clippers",
			"a pair of leather gloves",
			"a small pot for brewing",
		},
		CraftItems: "Herbalist tools can be used to make items like those listed in the Herbal Concoctions section of the Adventuring Gear table, with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Safely harvest and preserve a plant for later use", DC: 10},
			{Description: "Identify a poison with a sample of affected flesh", DC: 15},
			{Description: "Mix a hasty concoction to temporarily halt the effects of a deadly disease", DC: 20},
		},
	},

	"musical instrument": {
		Name:                "Musical Instrument",
		Description:         "Musical instruments come in many shapes and sizes with varying regional and cultural popularity. When you have proficiency in a musical instrument, choose one type: strings, brass, woodwinds, percussion, or keyboard instruments.",
		AssociatedAbilities: []string{"dex", "cha"},
		Components: []string{
			"an instrument",
			"a protective case",
			"any small pieces necessary for play, such as picks or reeds",
		},
		ExampleTasks: []ExampleTask{
			{Description: "Imitate a sound or tune you've heard before", DC: 10},
			{Description: "Provide accompaniment for an ongoing performance", DC: 15},
			{Description: "Outplay a professional (or similarly skilled) musician in direct competition", DC: 20},
		},
	},

	"navigator tools": {
		Name:                "Navigator Tools",
		Description:         "Navigator tools are used to read and draw maps, find your path or prevent you from becoming lost, and determine likely locations on a map for secret doors or hidden features.",
		AssociatedAbilities: []string{"int", "wis"},
		Components: []string{
			"quills, ink, and parchment",
			"two compasses",
			"calipers, and a ruler",
			"a sextant or telescope",
		},
		SpecialUses: map[string]string{
			"Draw Map": "Navigator tools can be used to create accurate maps while traveling, which grant advantage on checks made to prevent becoming lost.",
		},
		ExampleTasks: []ExampleTask{
			{Description: "Plot and track a travel course in a familiar region", DC: 10},
			{Description: "Discover your position according to a map or nautical chart", DC: 15},
			{Description: "Extrapolate a hidden or secret location on an unmarked map", DC: 20},
		},
	},

	"provisioner tools": {
		Name:                "Provisioner Tools",
		Description:         "Provisioner tools are used to prepare meals, make rations, and craft beverages.",
		AssociatedAbilities: []string{"con", "wis"},
		Components: []string{
			"a large jug, several jars",
			"a siphon with several feet of tubing",
			"a rolled set of knives for chopping, sawing, paring, and deboning",
			"a metal soup pot and a large skillet",
			"a small bundle of spoons, spatulas, ladels, and whisks",
			"several mixing bowls, a cutting board",
			"quantities of hops, dried herbs, and powdered spices",
		},
		CraftItems: "Provisioner tools can be used to make items with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Transform basic rations or subpar ingredients into a tavern-worthy meal", DC: 10},
			{Description: "Mask the pungent flavor of a poisonous herb mixed into a drink", DC: 15},
			{Description: "Outcook a professional (or similarly skilled) chef in direct competition", DC: 20},
		},
	},

	"smithing tools": {
		Name:                "Smithing Tools",
		Description:         "Smithing tools are used to build and repair items primarily made of metal. The use of smithing tools typically requires a dedicated forge, furnace, or some other source of intense heat. A campfire is only suitable for the smallest tasks like spot repairs.",
		AssociatedAbilities: []string{"dex", "con"},
		Components: []string{
			"a small anvil",
			"tongs",
			"hammers",
			"heat-resistant hide or cloth towels",
			"calipers",
			"billows",
		},
		CraftItems: "Smithing tools can be used to make items with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Identify the type and origin of a particular kind of metal", DC: 10},
			{Description: "Melt down and repurpose materials from horseshoes", DC: 15},
			{Description: "Safely pull an item out of a roaring fire before it's destroyed", DC: 20},
		},
	},

	"thieves' tools": {
		Name:                "Thieves' Tools",
		Description:         "Thieves' tools are used to pick locks, disable traps, and sabotage unattended items like saddles and doffed armor.",
		AssociatedAbilities: []string{"dex", "int"},
		Components: []string{
			"a roll of lockpicks and probes",
			"a small, telescoping mirror on a pole",
			"several triangular files",
			"a long set of scissors",
			"a pair of pliers",
		},
		SpecialUses: map[string]string{
			"Disarm Traps": "Thieves' tools can be used to disarm traps.",
			"Pick Locks":   "Thieves' tools can be used to pick locks. If no DC to pick a lock is provided, a standard nonmagical lock can be opened with a DC of 15.",
		},
		ExampleTasks: []ExampleTask{
			{Description: "Spot a patrolling guard from around a corner", DC: 10},
			{Description: "Pick a standard nonmagical lock", DC: 15},
			{Description: "Sabotage a ship to slowly fill with water as soon as it hits deep ocean", DC: 20},
		},
	},

	"tinker tools": {
		Name:                "Tinker Tools",
		Description:         "Tinker tools are used to construct mechanical or clockwork devices, repair damaged devices, or determine how to use unfamiliar mechanical or clockwork systems.",
		AssociatedAbilities: []string{"dex", "int"},
		Components: []string{
			"an array of hand tools",
			"an array of files",
			"pliers, tweezers, etching styluses",
			"threads, needles",
			"cloth and leather scraps",
			"glue",
			"loose cogs and pins",
			"a few sheets of thinly hammered metal",
		},
		CraftItems: "Tinker tools can be used to make items with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Provide accurate written or spoken instructions on operating a device you are familiar with", DC: 10},
			{Description: "Craft a temporary replacement for a missing part of a machine or device", DC: 15},
			{Description: "Disarm a complex device that is seconds from exploding", DC: 20},
		},
	},

	"trapper tools": {
		Name:                "Trapper Tools",
		Description:         "Trapper tools are used to set hunting traps, harvest hides or fur, and craft or repair leather goods like armor, pouches, or sturdy garments.",
		AssociatedAbilities: []string{"dex", "wis"},
		Components: []string{
			"a knife",
			"a small shovel",
			"a wooden mallet",
			"a small bundle of edgers, hole punchers, and sturdy needles",
			"thread, leather scraps",
			"a 30-ft. length of strong cord",
			"a quantity of salt",
			"tanning oil",
		},
		CraftItems: "Trapper tools can be used to make items with the crafting downtime activity.",
		ExampleTasks: []ExampleTask{
			{Description: "Identify a leather good, hide, or fur's age and origin", DC: 10},
			{Description: "Perfectly field strip the remains of creature with an unusually hard or magical hide—like that of a Dragon", DC: 15},
			{Description: "Secure a damaged leather saddle strap while atop a galloping horse", DC: 20},
		},
	},
}
