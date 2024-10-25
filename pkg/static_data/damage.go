package character

// DamageModifiers returns a map of names describing how susceptible
// a character is to damage providing a name and modifier
var DamageModifiers = func() map[string]float32 {
	return map[string]float32{
		"vulnerable": 2.0,
		"resistant":  0.5,
		"immune":     0.0,
	}
}

// DamageType returns a map of names describing the different types
// of damage allowed and a description of each
var DamageType = func() map[string]string {
	return map[string]string{
		"acid": `The corrosive spray of a black dragon's breath and ` +
			`the dissolving enzymes secreted by a black pudding deal acid damage.`,
		"bludgeoning": `Blunt force attacks–hammers, falling, constriction, ` +
			`and the like–deal bludgeoning damage.`,
		"cold": `The infernal chill radiating from an ice devil's spear and ` +
			`the frigid blast of a white dragon's breath deal cold damage.`,
		"fire": `Red dragons breathe fire, and many spells conjure flames to deal fire damage.`,
		"force": `Force is pure magical energy focused into a damaging form. ` +
			`Most effects that deal force damage are spells, including magic missile and spiritual weapon.`,
		"lightning": `A lightning bolt spell and a blue dragon's breath deal lightning damage.`,
		"necrotic": `Necrotic damage, dealt by certain undead and a spell such as chill touch, ` +
			`withers matter and even the soul.`,
		"piercing": `Puncturing and impaling attacks, including spears and monsters' bites, deal piercing damage.`,
		"poison":   `Venomous stings and the toxic gas of a green dragon's breath deal poison damage.`,
		"psychic":  `Mental abilities such as a mind flayer's psionic blast deal psychic damage.`,
		"radiant": `Radiant damage, dealt by a cleric's flame strike spell or an angel's smiting weapon, ` +
			`sears the flesh like fire and overloads the spirit with power.`,
		"slashing": `Swords, axes, and monsters' claws deal slashing damage.`,
		"thunder":  `A concussive burst of sound, such as the effect of the thunderwave spell, deals thunder damage.`,
	}
}
