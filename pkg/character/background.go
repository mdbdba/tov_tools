package character

import "tov_tools/pkg/static_data"

type Background struct {
	Name                         string
	Description                  string
	SkillProficiencies           []string                 // predefined
	SkillProficiencyOptions      map[string]ChoiceOptions // choose x from c1, c2, ...
	AdditionalProficiencies      []string                 // predefined
	AdditionalProficiencyOptions map[string]ChoiceOptions
	Equipment                    []static_data.EquipmentPackContent // predefined
	Money                        Money
	EquipmentOptions             map[string]ChoiceOptions
	TalentOptions                map[string]ChoiceOptions
	Motivations                  map[string]map[int]string // adventuring motivation, artistic expression, secret, etc.
}
