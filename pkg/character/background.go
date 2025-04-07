package character

type Background struct {
	name                         string
	skillProficiencies           []string // predefined
	skillProficiencyOptions      []ChoiceOptions
	additionalProficiencies      []string // predefined
	additionalProficiencyOptions []ChoiceOptions
	equipment                    []string // predefined
	equipmentOptions             []ChoiceOptions
	talentOptions                []ChoiceOptions
	motivations                  map[string]map[int]string // adventuring motivation, artistic expression, etc.
}
