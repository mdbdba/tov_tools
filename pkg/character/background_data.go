package character

import "tov_tools/pkg/static_data"

/*
type Background struct {
	Name                         string
	Description                  string
	SkillProficiencies           []string        // predefined
	SkillProficiencyOptions      map[string]ChoiceOptions // choose x from c1, c2, ...
	AdditionalProficiencies      []string        // predefined
	AdditionalProficiencyOptions map[string]ChoiceOptions
	Equipment                    []EquipmentPackContent // predefined
    Money                        Money
	EquipmentOptions             map[string]ChoiceOptions
	TalentOptions                map[string]ChoiceOptions
	Motivations                  map[string]map[int]string // adventuring motivation, artistic expression, secret, etc.
}
*/

var Backgrounds = map[string]Background{
	"adherent": {
		Name: "Adherent",
		Description: "Before you began adventuring, you committed yourself to a faith, belief, or cause. The " +
			"exacting tasks required of this commitment—daily prayers, holy rites, or cryptic " +
			"ceremonies—instilled in you a sense of duty and purpose. Perhaps you were a hopeful inductee " +
			"into the war god’s clergy, a priest excommunicated from a fiend-worshipping sect, or a " +
			"lifelong member of a secret society with global reach. In any case, you still carry the " +
			"teachings and traditions of your devotion.",
		SkillProficiencyOptions: map[string]ChoiceOptions{
			"skills": {
				NumberToSelect: 2,
				Options:        []string{"History", "Investigation", "Religion", "Persuasion"},
			},
		},
		AdditionalProficiencies: []string{"artist tools"},
		AdditionalProficiencyOptions: map[string]ChoiceOptions{
			"tools": {
				NumberToSelect: 1,
				Options:        static_data.GetToolsetNames(),
			},
		},
		Equipment: []static_data.EquipmentPackContent{
			{Name: "holy symbol", Quantity: 1},
			{Name: "incense", Quantity: 1},
			{Name: "vestments", Quantity: 1},
			{Name: "clothes common", Quantity: 1},
			{Name: "pouch", Quantity: 1},
		},
		Money: Money{GoldPieces: 10},
		TalentOptions: map[string]ChoiceOptions{
			"background_related": {
				NumberToSelect: 1,
				Options:        []string{"field medic", "mental fortitude", "ritualist"},
			},
		},
		Motivations: map[string]map[int]string{
			"adventuring": {
				1: "I can test the limits of my devotion out in the wider world through adventuring.",
				2: "Adventuring allows me to learn about and report on other religions and orders.",
				3: "Adventuring frees me to practice more unorthodox methods of worship.",
				4: "I may find others sworn to my order when I am out adventuring.",
				5: "Encountering new people while adventuring lets me share my faith with heretics, pagans, and the uninitiated.",
				6: "When I triumph through adventuring, I will bring glory and notoriety to my order.",
				7: "Adventuring furnishes me with the tithe my order deserves.",
				8: "Staying on the move keeps me from being dragged back to the order from which I narrowly escaped.",
			},
		},
	},
}
