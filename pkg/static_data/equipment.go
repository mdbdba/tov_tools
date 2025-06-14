package static_data

type ObjectArmorClass struct {
	ArmorClass int
}

var ArmorClassForObjects = map[string]ObjectArmorClass{
	"cloth":      {ArmorClass: 11},
	"paper":      {ArmorClass: 11},
	"rope":       {ArmorClass: 11},
	"crystal":    {ArmorClass: 13},
	"glass":      {ArmorClass: 13},
	"ice":        {ArmorClass: 13},
	"wood":       {ArmorClass: 15},
	"bone":       {ArmorClass: 15},
	"stone":      {ArmorClass: 17},
	"iron":       {ArmorClass: 19},
	"steel":      {ArmorClass: 19},
	"mithral":    {ArmorClass: 21},
	"adamantine": {ArmorClass: 23},
}

type EquipmentPackContent struct {
	Name     string
	Quantity int
}
type EquipmentPack struct {
	Name       string
	Contents   []EquipmentPackContent
	CostAmount int
	CostCoin   string // gp, sp, cp
}

var EquipmentPacks = map[string]EquipmentPack{
	"burglar": {
		Name: "Burglar’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "ball bearings", Quantity: 1},
			{Name: "bell", Quantity: 1},
			{Name: "candles", Quantity: 5},
			{Name: "crowbar", Quantity: 1},
			{Name: "hammer", Quantity: 1},
			{Name: "pitons", Quantity: 10},
			{Name: "lantern hooded", Quantity: 1},
			{Name: "oil", Quantity: 2},
			{Name: "rations", Quantity: 5},
			{Name: "tinderbox", Quantity: 1},
			{Name: "waterskin", Quantity: 1},
			{Name: "rope", Quantity: 1},
		},
		CostAmount: 16,
		CostCoin:   "gp",
	},
	"diplomat": {
		Name: "Diplomat’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "chest", Quantity: 1},
			{Name: "cases for maps and scrolls", Quantity: 2},
			{Name: "clothes fine", Quantity: 1},
			{Name: "ink", Quantity: 1},
			{Name: "pen", Quantity: 1},
			{Name: "lamp", Quantity: 1},
			{Name: "oil", Quantity: 2},
			{Name: "paper", Quantity: 5},
			{Name: "perfume", Quantity: 1},
			{Name: "sealing wax", Quantity: 1},
			{Name: "soap", Quantity: 1},
		},
		CostAmount: 39,
		CostCoin:   "gp",
	},
	"dungeoneer": {
		Name: "Dungeoneer’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "crowbar", Quantity: 1},
			{Name: "hammer", Quantity: 1},
			{Name: "pitons", Quantity: 10},
			{Name: "torches", Quantity: 10},
			{Name: "tinderbox", Quantity: 1},
			{Name: "rations", Quantity: 10},
			{Name: "waterskin", Quantity: 1},
			{Name: "rope", Quantity: 1},
		},
		CostAmount: 12,
		CostCoin:   "gp",
	},
	"entertainer": {
		Name: "Entertainer’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "bedroll", Quantity: 1},
			{Name: "clothes costume", Quantity: 2},
			{Name: "candles", Quantity: 5},
			{Name: "rations", Quantity: 5},
			{Name: "waterskin", Quantity: 1},
			{Name: "cosmetics", Quantity: 1},
		},
		CostAmount: 40,
		CostCoin:   "gp",
	},
	"explorer": {
		Name: "Explorer’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "bedroll", Quantity: 1},
			{Name: "mess kit", Quantity: 1},
			{Name: "tinderbox", Quantity: 1},
			{Name: "torch", Quantity: 10},
			{Name: "rations", Quantity: 10},
			{Name: "waterskin", Quantity: 1},
			{Name: "rope", Quantity: 1},
		},
		CostAmount: 10,
		CostCoin:   "gp",
	},
	"priest": {
		Name: "Priest’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "blanket", Quantity: 1},
			{Name: "candles", Quantity: 10},
			{Name: "tinderbox", Quantity: 1},
			{Name: "alms box", Quantity: 1},
			{Name: "incense", Quantity: 2},
			{Name: "censer", Quantity: 1},
			{Name: "vestments", Quantity: 1},
			{Name: "rations", Quantity: 2},
			{Name: "waterskin", Quantity: 1},
		},
		CostAmount: 19,
		CostCoin:   "gp",
	},
	"scholar": {
		Name: "Scholar’s Pack",
		Contents: []EquipmentPackContent{
			{Name: "backpack", Quantity: 1},
			{Name: "ink", Quantity: 1},
			{Name: "book", Quantity: 1},
			{Name: "pen", Quantity: 1},
			{Name: "paper", Quantity: 10},
			{Name: "sand", Quantity: 1},
			{Name: "knife", Quantity: 1},
		},
		CostAmount: 40,
		CostCoin:   "gp",
	},
}

type Gear struct {
	Name        string
	Category    string // general, Alchemical Concoctions, Ammunition, etc
	Description string // some gear is special and has a description
	CostAmount  int
	CostCoin    string  // gp, sp, cp
	WeightEach  float64 // in pounds.  So, .5, 1, 25.75, etc
	Quantity    int
}

var AdventuringGear = map[string]map[string]Gear{
	"general": {
		"abacus": {
			Name:       "Abacus",
			Category:   "General", // General, Alchemical Concoctions, Ammunition, etc
			CostAmount: 2,
			CostCoin:   "gp", // gp, sp, cp
			WeightEach: 2,
			Quantity:   1,
		},
		"alms box": {
			Name:       "Alms Box",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "cp",
			WeightEach: 2,
			Quantity:   1,
		},
		"ball bearings": {
			Name:     "Ball Bearings (bag)",
			Category: "General",
			Description: "As an action, you can spill these tiny metal balls from their pouch to cover a level, square " +
				"area, 10 feet on a side. A creature moving in this area must succeed on a DC 10 DEX save or " +
				"fall prone. A creature moving in the area at half speed doesn’t need to make the save.",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1000,
		},
		"bedroll": {
			Name:       "Bedroll",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 7,
			Quantity:   1,
		},
		"bell": {
			Name:       "Bell",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"blanket": {
			Name:       "Blanket",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 3,
			Quantity:   1,
		},
		"block and tackle": {
			Name:     "Block and Tackle",
			Category: "General",
			Description: "This is a set of pulleys with a cable threaded through them and a hook to attach " +
				"to items. A block and tackle allows you to hoist up to four times the weight you can " +
				"normally lift.",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 5,
			Quantity:   1,
		},
		"book": {
			Name:     "Book",
			Category: "General",
			Description: "A book might contain poetry, historical accounts, information pertaining to a field " +
				"of lore, diagrams, notes on contraptions, or just about anything that can be represented using " +
				"text or pictures. A spellbook is a separate item described later in this section.",
			CostAmount: 25,
			CostCoin:   "gp",
			WeightEach: 5,
			Quantity:   1,
		},
		"caltrops": {
			Name:     "Caltrops (bag)",
			Category: "General",
			Description: "As an action, you can spread a bag of caltrops to cover a 5-foot-square area. " +
				"A Large or smaller creature that enters the area must succeed on a DC 15 DEX save or " +
				"stop moving this turn and take 1 piercing damage. Taking this damage reduces the " +
				"creature’s walking speed by 10 feet until the creature regains at least 1 hit point. " +
				"A creature moving through the area at half speed doesn’t need to make the save.",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   20,
		},
		"candle": {
			Name:        "Candle",
			Category:    "General",
			Description: "For 1 hour, a candle sheds bright light in a 5-foot radius and dim light for an additional 5 feet.",
			CostAmount:  1,
			CostCoin:    "cp",
			WeightEach:  0,
			Quantity:    1,
		},
		"case": {
			Name:        "Case, map or scroll",
			Category:    "General",
			Description: "This cylindrical leather case can hold\nup to ten rolled-up sheets of paper.",
			CostAmount:  1,
			CostCoin:    "gp",
			WeightEach:  1,
			Quantity:    1,
		},
		"censer": {
			Name:        "Censer",
			Category:    "General",
			Description: "Used for burning incense",
			CostAmount:  1,
			CostCoin:    "cp",
			WeightEach:  0.5,
			Quantity:    1,
		},
		"chain": {
			Name:     "Chain (10 feet)",
			Category: "General",
			Description: "A chain is an object that has AC 20 and 10 HP. It can be broken as an action with a " +
				"successful DC 20 STR (Athletics) check.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 10,
			Quantity:   1,
		},
		"chalk": {
			Name:       "Chalk",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "cp",
			WeightEach: 0,
			Quantity:   1,
		},
		"climber's kit": {
			Name:     "Climber's Kit",
			Category: "General",
			Description: "A climber’s kit includes special pitons, boot tips, gloves, and a harness. " +
				"You can use the climber’s kit as an action to anchor yourself to a solid surface within " +
				"5 feet of you (this surface can’t be part of a creature); when you do, you can’t fall " +
				"more than 25 feet from your anchor point, and you can’t climb more than 25 feet away from " +
				"that point without undoing the anchor.",
			CostAmount: 25,
			CostCoin:   "gp",
			WeightEach: 12,
			Quantity:   1,
		},
		"clothes common": {
			Name:       "Clothes, common",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 3,
			Quantity:   1,
		},
		"clothes costume": {
			Name:       "Clothes, costume",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 4,
			Quantity:   1,
		},
		"clothes fine": {
			Name:       "Clothes, fine",
			Category:   "General",
			CostAmount: 15,
			CostCoin:   "gp",
			WeightEach: 6,
			Quantity:   1,
		},
		"vestments": {
			Name:       "Vestments",
			Category:   "General",
			CostAmount: 15,
			CostCoin:   "gp",
			WeightEach: 6,
			Quantity:   1,
		},
		"component pouch": {
			Name:     "Component Pouch",
			Category: "General",
			Description: "A component pouch is a small, watertight leather belt pouch that has compartments " +
				"to hold all the material components and other special items you need to cast your spells " +
				"(see Components in Chapter 7), except for those material components that have a listed gp " +
				"cost (as indicated in a spell’s description).",
			CostAmount: 25,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"cosmetics": {
			Name:        "Cosmetics",
			Category:    "General",
			Description: "A variety of cosmetics that can be used to enhance appearance.",
			CostAmount:  2,
			CostCoin:    "gp",
			WeightEach:  0.5,
			Quantity:    1,
		},
		"crowbar": {
			Name:        "Crowbar",
			Category:    "General",
			Description: "Using a crowbar grants advantage to STR checks where the crowbar’s leverage can be applied.",
			CostAmount:  2,
			CostCoin:    "gp",
			WeightEach:  5,
			Quantity:    1,
		},
		"fishing tackle": {
			Name:     "Fishing Tackle",
			Category: "General",
			Description: "This kit includes a wooden rod, silken line, corkwood bobbers, iron hooks, " +
				"lead sinkers, velvet lures, and netting.",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 4,
			Quantity:   1,
		},
		"hammer": {
			Name:       "Hammer",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 3,
			Quantity:   1,
		},
		"healer's kit": {
			Name:     "Healer's Kit",
			Category: "General",
			Description: "This kit is a leather pouch containing bandages, salves, and splints. " +
				"The kit has 10 uses. As an action, you can expend one use of the kit to stabilize a " +
				"creature that has 0 HP without needing to make a WIS (Medicine) check.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 3,
			Quantity:   1,
		},
		"holy water": {
			Name:     "Holy Water (flask)",
			Category: "General",
			Description: "As an action, you can splash the contents of this flask onto a creature within " +
				"5 feet of you or throw the flask up to 20 feet, shattering it on impact. In either " +
				"case, make a ranged attack against a target creature, treating the holy water as an " +
				"improvised weapon. If the target is a Fiend or Undead, it takes 2d6 radiant " +
				"damage—otherwise, the water has no effect. A creature that can cast at least one " +
				"1st-circle Divine spells can create holy water by performing a special ritual. The " +
				"ritual takes 1 hour to perform, uses 25 gp worth of powdered silver, and requires the " +
				"caster to expend a 1st circle spell slot.",
			CostAmount: 25,
			CostCoin:   "gp",
			WeightEach: 1,
			Quantity:   1,
		},
		"hourglass": {
			Name:       "Hourglass",
			Category:   "General",
			CostAmount: 25,
			CostCoin:   "gp",
			WeightEach: 1,
			Quantity:   1,
		},
		"hunting trap": {
			Name:     "Hunting Trap, basic",
			Category: "General",
			Description: "As an action, you can set a trap in an unoccupied space within 5 feet of you. " +
				"When set, this trap forms a saw-toothed steel ring that snaps shut when a creature " +
				"steps on a pressure plate in the center. A Large or smaller creature that enters the " +
				"trap’s space must succeed on a DC 13 DEX save or take 1d4 piercing damage and become " +
				"grappled until freed. A creature can use its action to make a DC 13 STR (Athletics) " +
				"check, freeing itself or another creature within its reach on a success. Each failed " +
				"check deals 1 piercing damage to the trapped creature.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 25,
			Quantity:   1,
		},
		"ink": {
			Name:       "Ink (1-ounce bottle)",
			Category:   "General",
			CostAmount: 10,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"ink pen": {
			Name:       "Ink Pen",
			Category:   "General",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: 0,
			Quantity:   1,
		},
		"incense": {
			Name:       "Incense (block)",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "cp",
			WeightEach: 0,
			Quantity:   1,
		},
		"knife": {
			Name:        "Knife",
			Category:    "General",
			Description: "A pen knife used for cutting and slicing.",
			CostAmount:  1,
			CostCoin:    "cp",
			WeightEach:  0,
			Quantity:    1,
		},
		"ladder": {
			Name:       "Ladder (10-foot)",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "sp",
			WeightEach: 25,
			Quantity:   1,
		},
		"lamp": {
			Name:     "Lamp",
			Category: "General",
			Description: "While lit, a lamp casts bright light in a 15-foot radius and dim light for an " +
				"additional 30 feet. Once lit, it burns for 6 hours on a flask (1 pint) of oil.",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 1,
			Quantity:   1,
		},
		"lantern bullseye": {
			Name:     "Lantern, bullseye",
			Category: "General",
			Description: "While lit, a bullseye lantern casts bright light in a 60-foot cone and dim " +
				"light for an additional 60 feet. Once lit, it burns for 6 hours on a flask (1 pint) of oil.",
			CostAmount: 10,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"lantern hooded": {
			Name:     "Lantern, hooded",
			Category: "General",
			Description: "While lit, a hooded lantern casts bright light in a 30-foot radius and dim " +
				"light for an additional 30 feet. Once lit, it burns for 6 hours on a flask (1 pint) of " +
				"oil. As an action, you can lower the hood, reducing the light to dim light in a 5-foot " +
				"radius.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"lock": {
			Name:     "Lock",
			Category: "General",
			Description: "A key is provided with the lock. Without the key, a creature proficient with " +
				"thieves’ tools (see Tools in this chapter) can pick this lock with a successful DC 15 " +
				"DEX Thieves’ Tools) check. Your GM can decide that better locks are available for " +
				"higher prices.",
			CostAmount: 10,
			CostCoin:   "gp",
			WeightEach: 1,
			Quantity:   1,
		},
		"magnifying glass": {
			Name:     "Magnifying Glass",
			Category: "General",
			Description: "This lens allows a closer look at small items. It is also useful as a substitute " +
				"for flint and steel when starting fires. Lighting a fire with a magnifying glass " +
				"requires sunlight to focus, tinder to light, and about 5 minutes for the fire to ignite. " +
				"A magnifying glass grants advantage on any ability check made to appraise or inspect an " +
				"item that is small or highly detailed.",
			CostAmount: 100,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"manacles": {
			Name:     "Manacles",
			Category: "General",
			Description: "These metal restraints can bind a Small or Medium creature. To escape the manacles, " +
				"a creature must use their action to make a successful DC 20 DEX (Sleight of Hand) check or " +
				"break them with a successful DC 20 STR (Athletics) check. Each set of manacles comes with one " +
				"key. Without the key, a creature proficient with thieves’ tools can pick the manacles’ lock " +
				"with a successful DC 15 DEX (Thieves’ Tools) check. Manacles are an object that has AC 19 and " +
				"15 HP.",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 6,
			Quantity:   1,
		},
		"mess kit": {
			Name:     "Mess Kit",
			Category: "General",
			Description: "This tin box contains a cup and simple cutlery. The box clamps together; one " +
				"side can be used as a cooking pan and the other as a plate or shallow bowl.",
			CostAmount: 2,
			CostCoin:   "sp",
			WeightEach: 1,
			Quantity:   1,
		},
		"mirror": {
			Name:       "Mirror, compact",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 0.5,
			Quantity:   1,
		},
		"net": {
			Name:     "Net",
			Category: "General",
			Description: "As an action, you can throw a net at a Large or smaller creature within 15 feet of you. The " +
				"net has no effect on creatures that are formless or incorporeal (like ghosts). Make a ranged " +
				"attack against the creature, treating the net as an improvised weapon. On a hit, the target is " +
				"restrained until it is freed. A creature can use its action to make a DC 10 STR (Athletics) check, " +
				"freeing itself or another creature within its reach on a success. Dealing 5 slashing damage to the " +
				"net (AC 10) also frees the creature without harming it, ending the effect and destroying the net.",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 3,
			Quantity:   1,
		},
		"paper": {
			Name:       "Paper (one sheet)",
			Category:   "General",
			CostAmount: 2,
			CostCoin:   "sp",
			WeightEach: 0,
			Quantity:   1,
		},
		"miner's pick": {
			Name:       "Pick, miner's",
			Category:   "General",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 10,
			Quantity:   1,
		},
		"piton": {
			Name:       "Piton",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "cp",
			WeightEach: 0.25,
			Quantity:   1,
		},
		"pole": {
			Name:       "Pole (10-foot)",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "cp",
			WeightEach: 7,
			Quantity:   1,
		},
		"potion of healing": {
			Name:     "Potion of Healing",
			Category: "General",
			Description: "A character who drinks the magical\nred fluid in this vial regains 2d4 + 2 hit points. " +
				"Drinking or administering a potion takes an action. For more information about this item, " +
				"see potion of healing in the Magic Item Descriptions section in this chapter.",
			CostAmount: 50,
			CostCoin:   "gp",
			WeightEach: 0.5,
			Quantity:   1,
		},
		"quiver": {
			Name:        "Quiver",
			Category:    "General",
			Description: "A quiver can hold up to 20 arrows or 20 crossbow bolts.",
			CostAmount:  1,
			CostCoin:    "gp",
			WeightEach:  1,
			Quantity:    1,
		},
		"portable ram": {
			Name:     "Ram, portable",
			Category: "General",
			Description: "You can use a portable ram to break down doors. You have advantage on relevant STR (Athletic) " +
				"checks made to do so.",
			CostAmount: 4,
			CostCoin:   "gp",
			WeightEach: 35,
			Quantity:   1,
		},
		"rations": {
			Name:     "Rations (1 day)",
			Category: "General",
			Description: "Rations consist of dry foods suitable for extended travel, including jerky, dried fruit, " +
				"hardtack, and nuts. One unit of rations is enough to feed one Medium or smaller creature for one day.",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 2,
			Quantity:   1,
		},
		"rope": {
			Name:        "Rope (50 feet)",
			Category:    "General",
			Description: "Rope is an object with AC 11 and 2 HP. It can be burst with a DC 17 STR (Athletics) check.",
			CostAmount:  1,
			CostCoin:    "gp",
			WeightEach:  10,
			Quantity:    1,
		},
		"sealing wax": {
			Name:       "Sealing Wax",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 0,
			Quantity:   1,
		},
		"merchant's scale": {
			Name:     "Scale, merchant's",
			Category: "General",
			Description: "A scale includes a small balance, pans, and a suitable assortment of weights up to 2 lb. " +
				"With it, you can measure the exact weight of suitably sized objects, such as raw precious metals " +
				"or trade goods, to help determine their worth.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 3,
			Quantity:   1,
		},
		"sand": {
			Name:       "Sand (small bag)",
			Category:   "General",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: .25,
			Quantity:   1,
		},
		"shovel": {
			Name:       "Shovel",
			Category:   "General",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 5,
			Quantity:   1,
		},
		"signal whistle": {
			Name:       "Signal Whistle",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "cp",
			WeightEach: 0,
			Quantity:   1,
		},
		"signet ring": {
			Name:       "Signet Ring",
			Category:   "General",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"soap": {
			Name:       "Soap",
			Category:   "General",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: 0,
			Quantity:   1,
		},
		"spellbook": {
			Name:        "Spellbook",
			Category:    "General",
			Description: "A spellbook is a leather‑bound tome with 100 blank pages suitable for recording spells.",
			CostAmount:  50,
			CostCoin:    "gp",
			WeightEach:  3,
			Quantity:    1,
		},
		"iron spikes": {
			Name:       "Spikes, iron",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 0.5,
			Quantity:   10,
		},
		"spyglass": {
			Name:     "Spyglass",
			Category: "General",
			Description: "Objects viewed through a spyglass are magnified to twice their size. A spyglass grants " +
				"advantage on any ability check made to view or inspect items that are far away.",
			CostAmount: 1000,
			CostCoin:   "gp",
			WeightEach: 1,
			Quantity:   1,
		},
		"tent": {
			Name:        "Tent, two-person",
			Category:    "General",
			Description: "A simple, portable canvas shelter that sleeps two Medium or smaller creatures.",
			CostAmount:  2,
			CostCoin:    "gp",
			WeightEach:  20,
			Quantity:    1,
		},
		"tinderbox": {
			Name:     "Tinderbox",
			Category: "General",
			Description: "This small container holds flint, fire steel, and tinder (usually dry cloth soaked in " +
				"light oil) used to kindle a fire. Using it to light a torch—or anything else with abundant, " +
				"exposed fuel—takes an action. Lighting any other fire takes 1 minute.",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 1,
			Quantity:   1,
		},
		"torch": {
			Name:     "Torch",
			Category: "General",
			Description: "A torch burns for 1 hour, providing bright light in a 20-foot radius and dim light for an " +
				"additional 20 feet. If you make a melee attack with a burning torch and hit, it deals 1 fire damage.",
			CostAmount: 1,
			CostCoin:   "cp",
			WeightEach: 1,
			Quantity:   1,
		},
		"whetstone": {
			Name:       "Whetstone",
			Category:   "General",
			CostAmount: 1,
			CostCoin:   "cp",
			WeightEach: 1,
			Quantity:   1,
		},
	},

	"alchemical concoctions": {
		"acid (vial)": {
			Name:     "Acid (vial)",
			Category: "Alchemical Concoctions", // General, Alchemical Concoctions, Ammunition, etc
			Description: "As an action, you can splash the contents of this vial onto a creature within 5 " +
				"feet of you or throw the vial up to 20 feet, shattering it on impact. In either case, " +
				"make a ranged attack against a creature or object, treating the acid as an improvised weapon. " +
				"On a hit, the target takes 2d6acid damage.",
			CostAmount: 25,
			CostCoin:   "gp", // gp, sp, cp
			WeightEach: 1,
			Quantity:   1,
		},
		"alchemist's fire": {
			Name:     "Alchemist's fire (flask)",
			Category: "Alchemical Concoctions",
			Description: "This sticky, adhesive fluid ignites when exposed to air. As an action, you can " +
				"throw this flask up to 20 feet, shattering it on impact. Make a ranged attack against a " +
				"creature or object, treating the alchemist’s fire as an improvised weapon. On a hit, the " +
				"target takes 1d4 fire damage at the start of each of its turns. A creature can end this " +
				"damage by using its action to make a DC 10 DEX (Sleight of Hand) check to extinguish the " +
				"flames.",
			CostAmount: 50,
			CostCoin:   "gp",
			WeightEach: 1,
			Quantity:   1,
		},
		"oil": {
			Name:     "Oil (flask)",
			Category: "Alchemical Concoctions",
			Description: "Oil usually comes in a clay flask that holds 1 pint. As an action, you can splash oil onto " +
				"a creature within 5 feet of you or throw the flask up to 20 feet, shattering on impact. Make a " +
				"ranged attack against a target creature or object, treating the oil as an improvised weapon. On a " +
				"hit, the target is covered in oil. If the target takes any fire damage before the oil dries " +
				"(after 1 minute), the target takes an additional 5 fire damage from the burning oil. You can " +
				"also pour a flask of oil on the ground to cover a 5-foot-square area, provided that the surface is " +
				"level. If lit, the oil burns for 2 rounds and deals 5 fire damage to any creature that enters the " +
				"area or ends its turn in the area. A creature can take this damage only once per turn.",
			CostAmount: 1,
			CostCoin:   "sp",
			WeightEach: 1,
			Quantity:   1,
		},
	},
	"ammunition": {
		"arrows": {
			Name:       "Arrows (20)",
			Category:   "Ammunition", // General, Alchemical Concoctions, Ammunition, etc
			CostAmount: 1,
			CostCoin:   "gp", // gp, sp, cp
			WeightEach: 1,
			Quantity:   20,
		},
		"blowgun needles": {
			Name:       "Blowgun needles (50)",
			Category:   "Ammunition",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 1,
			Quantity:   50,
		},
		"crossbow bolts": {
			Name:       "Crossbow bolts (20)",
			Category:   "Ammunition",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 1.5,
			Quantity:   20,
		},
		"sling bullets": {
			Name:       "Sling bullets (20)",
			Category:   "Ammunition",
			CostAmount: 4,
			CostCoin:   "cp",
			WeightEach: 1.5,
			Quantity:   20,
		},
	},
	"containers": {
		"backpack": {
			Name:       "Backpack",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 5,
			Quantity:   1,
		},
		"barrel": {
			Name:       "Barrel",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 70,
			Quantity:   1,
		},
		"basket": {
			Name:       "Basket",
			Category:   "Containers",
			CostAmount: 4,
			CostCoin:   "sp",
			WeightEach: 2,
			Quantity:   1,
		},
		"bottle": {
			Name:       "Bottle, glass",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"chest": {
			Name:       "Chest",
			Category:   "Containers",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 25,
			Quantity:   1,
		},
		"flask": {
			Name:       "Flask",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: 1,
			Quantity:   1,
		},
		"Tankard": {
			Name:       "Tankard",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: 1,
			Quantity:   1,
		},
		"jug": {
			Name:       "Jug",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: 4,
			Quantity:   1,
		},
		"pitcher": {
			Name:       "Pitcher",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "cp",
			WeightEach: 4,
			Quantity:   1,
		},
		"cooking pot": {
			Name:       "Pot, cooking",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "gp",
			WeightEach: 10,
			Quantity:   1,
		},
		"pouch": {
			Name:     "Pouch",
			Category: "Containers",
			Description: "A cloth or leather pouch can hold up to 20 sling bullets or 50 blowgun needles, " +
				"among other things. A compartmentalized pouch for holding spell components is called a component " +
				"pouch (described earlier in this section).",
			CostAmount: 5,
			CostCoin:   "sp",
			WeightEach: 1,
			Quantity:   1,
		},
		"sack": {
			Name:       "Sack",
			Category:   "Containers",
			CostAmount: 1,
			CostCoin:   "cp",
			WeightEach: 0.5,
			Quantity:   1,
		},
		"vial": {
			Name:       "Vial",
			Category:   "Containers",
			CostAmount: 1,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"waterskin": {
			Name:       "Waterskin (full)",
			Category:   "Containers",
			CostAmount: 2,
			CostCoin:   "sp",
			WeightEach: 5,
			Quantity:   1,
		},
	},
	"herbal concoctions": {
		"antitoxin": {
			Name:     "Antitoxin (vial)",
			Category: "Herbal Concoctions",
			Description: "A creature that drinks this vial of liquid gains advantage on saves against " +
				"poison for 1 hour. It confers no benefit to Undead or Constructs.",
			CostAmount: 50,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"perfume": {
			Name:       "Perfume (vial)",
			Category:   "Herbal Concoctions",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"poison basic": {
			Name:     "Poison, basic (vial)",
			Category: "Herbal Concoctions",
			Description: "You can use the poison in this vial to coat one weapon or up to three pieces of ammunition " +
				"that deal piercing or slashing damage. Applying the poison takes an action. A creature hit by " +
				"the poisoned weapon or ammunition must succeed on a DC 10 CON save or take 1d4 poison damage " +
				"(in addition to the weapon’s normal damage). Once applied, the poison retains potency for 1 minute before drying.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"poison ether": {
			Name:     "Poison, essence of ether (vial)",
			Category: "Herbal Concoctions",
			Description: "A creature subjected to this poison must succeed on a DC 15 CON save or be poisoned for " +
				"8 hours. The poisoned creature is unconscious. The unconscious effect ends if the creature takes " +
				"damage or if another creature takes an action to shake it awake.",
			CostAmount: 300,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"poison last gasp": {
			Name:     "Poison, last gasp (vial)",
			Category: "Herbal Concoctions",
			Description: "A creature subjected to this poison must succeed on a DC 13 CON save or be poisoned for " +
				"1 minute. The poisoned creature is paralyzed. The creature can repeat the save at the end of each " +
				"of its turns, ending both effects on a success.",
			CostAmount: 200,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
		"poison midnight tears": {
			Name:     "Poison, midnight tears (vial)",
			Category: "Herbal Concoctions",
			Description: "A creature who ingests this poison suffers no effect until the stroke of midnight. If the " +
				"poison has not been neutralized before then, the creature must succeed on a DC 17 CON save, " +
				"taking 31(9d6) poison damage on a failure or half as much damage on a success.",
			CostAmount: 1500,
			CostCoin:   "gp",
			WeightEach: 0,
			Quantity:   1,
		},
	},
	"spellcasting foci": {
		"arcane focus": {
			Name:     "Arcane Focus",
			Category: "Spellcasting Foci",
			Description: "This item can be used as a spellcasting\nfocus to channel Arcane spells. Examples " +
				"include an orb, a crystal, a rod, a specially made staff, or a wooden wand.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"holy symbol": {
			Name:     "Holy Symbol",
			Category: "Spellcasting Foci",
			Description: "This item can be used as a spellcasting focus to channel Divine spells. Examples include " +
				"an amulet depicting a deity’s symbol, that same symbol engraved or inlaid as an emblem on a " +
				"shield, or a tiny box holding a fragment of a sacred relic.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"primordial focus": {
			Name:     "Primordial Focus",
			Category: "Spellcasting Foci",
			Description: "This item can be used as a spellcasting\nfocus to channel Primordial spells. Examples " +
				"include a totem made from natural materials like feather or bone, a wooden staff carved from a " +
				"living tree, or a yew wand.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
		"wyrd focus": {
			Name:     "Wyrd Focus",
			Category: "Spellcasting Foci",
			Description: "This item can be used as a spellcasting focus to channel Wyrd spells. " +
				"Examples include an amulet carved from bone, a charm bag filled with rare herbs and crystals, " +
				"or a wand made of starmetal.",
			CostAmount: 5,
			CostCoin:   "gp",
			WeightEach: 2,
			Quantity:   1,
		},
	},
}
