package character

// Heritage represents upbringing and cultural traits
type Heritage struct {
	Name               string
	SkillProficiencies []string // e.g., ["Stealth", "Arcana"]
	Languages          []string
	CulturalTraits     map[string]string // e.g., "City Navigation": "Bonus to find your way in big cities"
}
