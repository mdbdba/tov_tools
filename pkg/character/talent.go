package character

import (
	"fmt"
)

type Benefit interface {
	Apply(c *Character) error // Applies the benefit to the character
	Description() string      // Returns a human-readable description of the benefit
}

type Talent struct {
	Name         string                  // The name of the talent
	Category     string                  // magic, martial, or technical
	Description  string                  // A description of what the talent represents or does
	Prerequisite func(c *Character) bool // A function to check if a character meets the prerequisite
	Benefits     []Benefit               // A list of benefits provided by the talent
	// Source       string                  // What granted this talent, was it a specific background or a human getting
	// an extra talent, etc.
}

type SkillBonusMultiplierBenefit struct {
	SkillName       string
	BonusMultiplier float64
}

func (b *SkillBonusMultiplierBenefit) Apply(c *Character) error {
	// Logic to apply the bonus to the character's skill
	// This assumes "AddSkillBonus" is a method to grant custom bonuses
	c.AddSkillBonusMultiplier(b.SkillName, b.BonusMultiplier)

	return nil
}

func (b *SkillBonusMultiplierBenefit) Description() string {
	return fmt.Sprintf("Increase your proficiency bonus for any ability check that uses the %s skill by %f times", b.SkillName, b.BonusMultiplier)
}
func (b *SkillBonusMultiplierBenefit) Label() string {
	return fmt.Sprintf("SkillBonusMultiplier")
}

type FlatBonusBenefit struct {
	Attribute string // e.g., "strength" or "dexterity"
	Bonus     int
}

func (b *FlatBonusBenefit) Apply(c *Character) error {
	// Logic to add the bonus to the character's attribute
	c.AddAbilityBonus(b.Attribute, b.Label(), b.Bonus)
	return nil
}

func (b *FlatBonusBenefit) Description() string {
	return fmt.Sprintf("Gain a +%d bonus to %s", b.Bonus, b.Attribute)
}
func (b *FlatBonusBenefit) Label() string {
	return fmt.Sprintf("FlatBonusBenefit")
}

type SpellSwapBenefit struct {
	OldSpell string
	NewSpell string
}

func (b *SpellSwapBenefit) Apply(c *Character) error {
	// Ensure the character knows the old spell
	found := false
	for i, spell := range c.SpellBook {
		if spell == b.OldSpell {
			// Swap the spell
			c.SpellBook[i] = b.NewSpell
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("character does not know the spell '%s'", b.OldSpell)
	}

	// The character now has the new spell
	return nil
}

func (b *SpellSwapBenefit) Description() string {
	return fmt.Sprintf("Replace one spell you know (%s) with a new spell (%s)", b.OldSpell, b.NewSpell)
}
