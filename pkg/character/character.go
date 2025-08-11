package character

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
	"tov_tools/pkg/dice"
	"tov_tools/pkg/helpers"
	"tov_tools/pkg/static_data"

	"go.uber.org/zap"
)

var zapLogger *zap.SugaredLogger

type AbilityDependency struct {
	Name            string
	DependentSkills []string
	DependentValues map[string]func(*Character) int // Functions to recalculate
}

var DependencyLookup = map[string]AbilityDependency{
	"str": {
		Name:            "str",
		DependentSkills: []string{"athletics"},
		DependentValues: map[string]func(*Character) int{},
	},
	"dex": {
		Name:            "dex",
		DependentSkills: []string{"acrobatics", "sleight of hand", "stealth"},
		DependentValues: map[string]func(*Character) int{
			"InitiativeBonus": func(c *Character) int { return c.Abilities.Modifiers["dex"] },
		},
	},
	"con": {
		Name:            "con",
		DependentSkills: []string{},
		DependentValues: map[string]func(*Character) int{},
	},
	"int": {
		Name:            "int",
		DependentSkills: []string{"arcana", "history", "investigation", "nature", "religion"},
		DependentValues: map[string]func(*Character) int{
			"PassiveInvestigation": func(c *Character) int {
				returnValue := 10
				if c.AbilitySkills["investigation"].Proficient {
					returnValue += c.GetProficiencyBonus()
				}
				return c.GetAbilityModifier("int") + returnValue
			},
		},
	},
	"wis": {
		Name:            "wis",
		DependentSkills: []string{"animal handling", "insight", "medicine", "perception", "survival"},
		DependentValues: map[string]func(*Character) int{
			"PassivePerception": func(c *Character) int {
				returnValue := 10
				if c.AbilitySkills["perception"].Proficient {
					returnValue += c.GetProficiencyBonus()
				}
				return c.GetAbilityModifier("wis") + returnValue
			},
			"PassiveInsight": func(c *Character) int {
				returnValue := 10
				if c.AbilitySkills["insight"].Proficient {
					returnValue += c.GetProficiencyBonus()
				}
				return c.GetAbilityModifier("wis") + returnValue
			},
		},
	},
	"cha": {
		Name:            "cha",
		DependentSkills: []string{"deception", "intimidation", "performance", "persuasion"},
		DependentValues: map[string]func(*Character) int{},
	},
}

// ConditionEffects returns a map of names describing
// conditions that can happen to a character.
var ConditionEffects = func() map[string][]string {
	return map[string][]string{
		"blinded": {"A blinded character can't see and automatically fails any ability check that requires sight.",
			"Attack rolls against a blinded character have advantage, and a blinded creature's attack rolls have disadvantage."},
		"charmed": {"A charmed creature can't attack the charmer or target the charmer with harmful Abilities or magical effects.",
			"The charmer has advantage on any ability check to interact socially with the charmed creature."},
		"deafened":   {"A deafened character can't hear and automatically fails any ability check that requires hearing."},
		"exhaustion": {"Measured in levels, a character's exhaustion level effects combine the higher the level, the greater the exhaustion."},
		"frightened": {"A frightened character has disadvantage on ability checks and attack rolls while the source of its fear is within line of sight.",
			"The character can't willingly move closer to the source of its fear."},
		"grappled": {"A grappled creature's speed becomes 0, and it can't benefit from any bonus to its speed.",
			"The condition ends if the grappler is incapacitated (see the condition incapacitated).",
			"The condition also ends if an effect removes the grappled creature from the reach of the grappler or moves the grappler closer than 5 feet to the creature."},
		"incapacitated": {"An incapacitated creature can't take actions or reactions."},
		"invisible": {"An invisible creature is impossible to see without the aid of magic or a special sense. For the purpose of hiding, the creature is heavily obscured.",
			"The character has advantage on Dexterity (Stealth) checks and can take the Hide action as a bonus action.",
			"Attack rolls against an invisible creature have disadvantage, and an invisible creature's attack rolls have advantage"},
		"paralyzed": {"A paralyzed character is incapacitated (see the condition incapacitated), can't move, or speak.",
			"The character automatically fails Strength and Dexterity saves.",
			"Attack rolls against a paralyzed character have advantage",
			"Any attack that hits the character is a critical hit if the attacker is within 5 feet of the character."},
		"petrified": {"A petrified character is transformed, along with any nonmagical object it is wearing or carrying, into a solid inanimate substance (usually stone). Its weight increases by a factor of ten, and it ceases aging.",
			"The character is incapacitated, can't move, or speak.",
			"Attack rolls against a petrified creature have advantage",
			"The character automatically fails Strength and Dexterity saves.",
			"The character has resistance to all damage.",
			"The character is immune to poison and disease, although a poison or disease already in its system is suspended, not neutralized."},
		"poisoned": {"A poisoned character has disadvantage on attack rolls and ability checks."},
		"prone": {"A prone character's only movement options are to crawl, or to stand up. Standing up costs an amount of movement equal to half its speed and ends the condition.",
			"The character has disadvantage on attack rolls",
			"An attack roll against a prone creature has advantage if the attacker is within 5 feet of the creature. Otherwise, the attack roll has disadvantage."},
		"restrained": {"A restrained character's speed becomes 0, and it can't benefit from any bonus to its speed.",
			"Attack rolls against a restrained creature have advantage, and a restrained creature's attack rolls have disadvantage",
			"The character has disadvantage on Dexterity saving throws."},
		"stunned": {"A stunned character is incapacitated (see the condition incapacitated), can't move, and speak only faintly.",
			"The character automatically fails Strength and Dexterity saves.",
			"Attack rolls against a stunned creature have advantage"},
		"surprised": {"A surprised character can't move or take an action on its first turn of combat, and it can't take a reaction until after its first turn ends."},
		"unconscious": {"An unconscious character is incapacitated (see the condition incapacitated), can't move or speak, and is unaware of its surroundings.",
			"The character drops whatever it's holding and falls prone.",
			"The character automatically fails Strength and Dexterity saves.",
			"Attack rolls against an unconscious character have advantage",
			"Any attack that hits the character is a critical hit if the attacker is within 5 feet of the character."},
		"dead": {"Magical healing or a WIS(Medicine) check can't help a dead character. ",
			"Only a spell like revivify or resurrection can bring the character back to life."},
	}
}

type AbilitySkillProficiency struct {
	Skill  string
	Source string
}

type AbilitySkillBonus struct {
	Bonus  int
	Source string
}

type AbilitySkill struct {
	Ability    string
	Proficient bool
	Value      int
}

type HitDie struct {
	SourceClass string
	DiceType    string
	Max         int
	Used        int
}

type DamageAudit struct {
	ID              string
	DamageType      string
	HitPointsBefore int
	BaseAmount      int
	Adjustments     map[string]int
	TotalAmount     int
	HitPointsAfter  int
}

type MovementValue struct {
	Speed int
}

type MovementRateMultiplier struct {
	Type       string
	Multiplier float64
	Source     string
}

var Movement = func(baseMovement float64) map[string]MovementValue {
	return map[string]MovementValue{
		"walking": {
			int(baseMovement),
		},
		"running": {
			int(2 * baseMovement),
		},
		"climbing": {
			int(baseMovement * 0.5),
		},
		"crawling": {
			int(baseMovement * 0.5),
		},
		"difficult terrain": {
			int(baseMovement * 0.5),
		},
		"running long jump": {
			int(baseMovement * 0.5),
		},
		"standing long jump": {
			int(baseMovement * 0.25),
		},
		"swimming": {
			int(baseMovement * 0.5),
		},
		"flying": {0},
	}
}

var InitMovementBonus = func() map[string]map[string]MovementValue {
	return map[string]map[string]MovementValue{
		"walking":            {"init": {0}},
		"running":            {"init": {0}},
		"climbing":           {"init": {0}},
		"crawling":           {"init": {0}},
		"difficult terrain":  {"init": {0}},
		"running long jump":  {"init": {0}},
		"standing long jump": {"init": {0}},
		"swimming":           {"init": {0}},
		"flying":             {"init": {0}},
	}
}

type VantageType string

func (v VantageType) IsValid() bool {
	switch v {
	case ADV, DIS, NRM:
		return true
	}
	return false
}

const (
	ADV VantageType = "advantage"
	DIS VantageType = "disadvantage"
	NRM VantageType = "normal" // For straight rolls
)

// ConditionAdjustment ConditionAdjustments are for modifying rolls with
// advantage or disadvantage that relate to the character's
// condition.  Created to allow for anointed Heritage to
// give death save rolls advantage.
type ConditionAdjustment struct {
	Condition string      // DeathSaves
	Vantage   VantageType // advantage or disadvantage
	Source    string      // what gave the advantage or disadvantage
}

type DeathSaveAudit struct {
	SaveSuccess     bool
	CriticalSuccess bool
	CriticalFailure bool
	RollData        dice.Roll
}

// AuditEntry represents a change to a character field
type AuditEntry struct {
	Field     string      // Name of the field that changed
	OldValue  interface{} // Previous value
	NewValue  interface{} // New value
	Source    string      // What caused this change (e.g., "Level Up", "Background", "Item")
	Timestamp time.Time   // When the change occurred
}

type SelectedTraits struct {
	Name                string
	Description         string
	addToCharacter      func(c *Character) error
	removeFromCharacter func(c *Character) error
}

type LineageSelectionstruct struct {
	LineageName string
	Traits      map[string]SelectedTraits
}

type CharacterDescription struct {
	Age          int    `json:"age"`
	Size         string `json:"size"`
	HeightFeet   int    `json:"height_feet"`
	HeightInches int    `json:"height_inches"`
	WeightPounds int    `json:"weight_pounds"`
	Gender       string `json:"gender"`
	EyeColor     string `json:"eye_color"`
	HairColor    string `json:"hair_color"`
	SkinColor    string `json:"skin_color"`
	Description  string `json:"description"`
}

// Character represents a character in the game
type Character struct {
	UserId                       string
	ID                           string
	Name                         string
	Description                  *CharacterDescription
	OverallLevel                 int
	CharacterLevels              map[string]int
	CharacterClassStr            string // if multiclassing this will be class 1/class 2/class 3/etc
	CharacterClassBuildType      ClassBuildType
	CharacterSubClassToImplement Subclass // store subclass in case the pc is < 3rd level
	CharacterSubClass            Subclass
	DamageTypeAdjustments        map[string]string
	HitDice                      []HitDie
	Lineage                      Lineage
	LineageChoices               map[string][]string
	LineageInputRequired         bool
	Heritage                     Heritage
	HeritageChoices              map[string][]string
	HeritageInputRequired        bool
	KnownLanguages               []string
	KnownLanguagesInputRequired  bool
	Background                   Background
	BackgroundChoices            map[string][]string
	BackgroundInputRequired      bool
	Traits                       map[string]string
	TraitChoices                 map[string][]string
	BaseSkills                   map[string]int
	BaseSkillBonus               map[string]int
	Abilities                    AbilityArray
	AbilitySaveModifiers         map[string]int
	RollingOption                string
	HitPointBonuses              map[string]int
	TotalHitPointBonuses         int
	MaxHitPoints                 int
	TemporaryHitPoints           int
	CurrentHitPoints             int
	InitiativeBonus              int
	PassiveInvestigation         int
	PassivePerception            int
	PassiveInsight               int
	Talents                      map[string]Talent
	TalentsChoices               map[string][]string
	TalentsInputRequired         bool
	DeathSaves                   [3]int
	SpellcastingAbility          string
	SpellBook                    []string
	SkillProficiencies           map[string]AbilitySkillProficiency
	SkillBonus                   map[string]map[string]AbilitySkillBonus
	ProficiencyBonusBonus        map[string]AbilitySkillBonus
	Tools                        map[string]static_data.Tool
	TotalSkillModifiers          map[string]int
	Equipment                    []string
	MovementBase                 map[string]MovementValue
	MovementBonus                map[string]map[string]MovementValue
	TotalMovement                map[string]MovementValue
	AbilitySkills                map[string]AbilitySkill
	ConditionAdjustments         map[string][]ConditionAdjustment
	Conditions                   map[string]string // key = condition name, value = note
	AbilityScoreOrderPreference  []string
	KeyAbilities                 []string
	History                      *HistoryAudit
}

// HistoryAudit represents an auditable record for the character.  When there's
//
//	a change to a character value, the attr name that changed will be the map key
type HistoryAudit struct {
	CharacterId  string
	Audits       map[string][]AuditEntry
	DamageAudits []DamageAudit
}

func (c *Character) SetConditionAdjustment(condition string, vantage VantageType, source string) {
	c.ConditionAdjustments[condition] = append(c.ConditionAdjustments[condition], ConditionAdjustment{
		Condition: condition,
		Vantage:   vantage,
		Source:    source,
	})
}

func (c *Character) CalculateMovement() {
	c.TotalMovement = make(map[string]MovementValue)
	for key, movement := range c.MovementBonus {
		runningTotal := 0
		for _, bonus := range movement {
			runningTotal += bonus.Speed
		}
		c.TotalMovement[key] = MovementValue{
			Speed: c.MovementBase[key].Speed + runningTotal,
		}
	}
}

func (c *Character) SetAbilitySaveModifiers() {
	c.AbilitySaveModifiers = AbilityArrayTemplate()
	// being able to multiclass means that we need to make sure
	// proficiencies aren't duplicated and warn if they are.
	proficiencies := map[string]bool{}
	for _, h := range helpers.GetMapKeys(c.CharacterLevels) {
		for _, val := range Classes[h].SaveProficiencies {
			if _, ok := proficiencies[val]; ok {
				fmt.Printf("Warning: %s has a duplicate proficiency %s\n", h, val)
			} else {
				proficiencies[val] = true
			}
		}
	}
	for i := range c.Abilities.Modifiers {
		c.AbilitySaveModifiers[i] = c.GetAbilityModifier(i)
		for _, proficiency := range helpers.GetMapKeys(proficiencies) {
			if i == proficiency {
				c.AbilitySaveModifiers[i] += c.GetProficiencyBonus()
			}
		}
	}

}

func (c *Character) IsProficientIn(skill string) bool {
	for _, proficiency := range c.SkillProficiencies {
		if proficiency.Skill == skill {
			return true
		}
	}
	return false
}

func (c *Character) CalculateTotalSkillBonus(skill string) int {
	runningTotal := c.Abilities.Modifiers[SkillAbilityLookup()[skill]]
	if c.IsProficientIn(skill) {
		runningTotal += c.GetProficiencyBonus()
	}

	for _, bonus := range c.SkillBonus[skill] {
		runningTotal += bonus.Bonus
	}

	return runningTotal
}

func (c *Character) UpdateDependencies(ability string) {
	if dep, ok := DependencyLookup[ability]; ok {
		for _, skill := range dep.DependentSkills {

			c.AbilitySkills[skill] = AbilitySkill{
				Ability:    ability,
				Proficient: c.IsProficientIn(skill),
				Value:      c.CalculateTotalSkillBonus(skill),
			}
		}
		// Update other dependent values
		for key, calculationFunc := range dep.DependentValues {
			switch key {
			case "InitiativeBonus":
				c.InitiativeBonus = calculationFunc(c)
			case "PassiveInvestigation":
				c.PassiveInvestigation = calculationFunc(c)
			case "PassivePerception":
				c.PassivePerception = calculationFunc(c)
			case "PassiveInsight":
				c.PassiveInsight = calculationFunc(c)
			}
		}
	}
}

func (c *Character) SetAbilitySkills() {

	c.AbilitySkills = map[string]AbilitySkill{}
	c.TotalSkillModifiers = map[string]int{}
	// init
	for skill, ability := range SkillAbilityLookup() {
		c.AbilitySkills[skill] = AbilitySkill{
			Ability:    ability,
			Proficient: false,
			Value:      0,
		}
	}
	// set proficiencies
	for _, i := range c.SkillProficiencies {
		// Retrieve the AbilitySkill struct, modify it, and write it back to the map
		abilitySkill := c.AbilitySkills[i.Skill]
		abilitySkill.Proficient = true
		c.AbilitySkills[i.Skill] = abilitySkill
	}

	// set values
	for skill := range c.AbilitySkills {

		// Retrieve the actual struct
		abilitySkill := c.AbilitySkills[skill]

		runningTotal := 0 // c.Abilities.Modifiers[abilitySkill.Ability]
		if abilitySkill.Proficient {
			runningTotal += c.GetProficiencyBonus()
		}

		for i := range c.SkillBonus[skill] {
			runningTotal += c.SkillBonus[skill][i].Bonus
		}

		abilitySkill.Value = runningTotal + c.Abilities.Modifiers[abilitySkill.Ability]
		c.AbilitySkills[skill] = abilitySkill
		c.TotalSkillModifiers[skill] = runningTotal
	}
}

func (c *Character) AddTalent(t Talent, source string) error {
	// Check prerequisite
	if !t.Prerequisite(c) {
		return fmt.Errorf("character does not meet the prerequisites for talent: %s", t.Name)
	}
	// Apply all benefits of the Talent
	for _, benefit := range t.Benefits {
		if err := benefit.Apply(c); err != nil {
			return fmt.Errorf("failed to apply benefit '%s': %v", benefit.Description(), err)
		}
	}
	// Record audit before adding
	audit := AuditEntry{
		Field:     "Talents",
		OldValue:  len(c.Talents),
		NewValue:  len(c.Talents) + 1,
		Source:    source,
		Timestamp: time.Now(),
	}

	c.Talents[t.Name] = t
	// Record the audit
	c.History.Audits["Talents"] = append(c.History.Audits["Talents"], audit)

	return nil
}

func (c *Character) GetProficiencyBonus() int {
	base := c.OverallLevel/4 + 2
	bonus := 0
	for i := range c.ProficiencyBonusBonus {
		bonus += c.ProficiencyBonusBonus[i].Bonus
	}
	return base + bonus
}

func (c *Character) GetAbility(ability string) int {
	return c.Abilities.Values[ability]
}
func (c *Character) GetAbilityModifier(ability string) int {
	return c.Abilities.Modifiers[ability]
}
func (c *Character) GetSkillBonus(skill string) int {
	return c.AbilitySkills[skill].Value
}
func (c *Character) IncreaseAbility(ability string) error {
	abilityMax := 20
	if c.Abilities.IsMonsterOrGod {
		abilityMax = 30
	}
	if c.Abilities.Values[ability]+1 <= abilityMax {
		c.Abilities.Base[ability]++
	} else {
		return fmt.Errorf("character has already reached the maximum for %s", ability)
	}
	c.Abilities.setValuesAndModifiers()
	c.UpdateDependencies(ability)
	return nil
}

func (c *Character) AddSkillBonusMultiplier(skillName string, multiplier float64) {
	c.BaseSkillBonus[skillName] += int(float64(c.GetProficiencyBonus()) * multiplier)
}

func (c *Character) AddAbilityBonus(ability string, reason string, bonus int) {
	c.Abilities.AdjustBonuses(ability, reason, bonus, zapLogger)

}

// ValidateName checks if the Name is not empty and does not contain invalid characters.
func ValidateName(name string) error {

	// Ensure the Name is not empty
	if name == "" {
		return errors.New("name cannot be empty")
	}

	// Regular expression to match invalid characters
	invalidCharsPattern := `[@#$%^&*()+\-\'\"<>?/=_1234567890]`
	matched, err := regexp.MatchString(invalidCharsPattern, name)
	if err != nil {
		return err // Handle regex error (unlikely, but good practice)
	}

	// Check if any invalid character is found
	if matched {
		return errors.New("name contains invalid characters")
	}

	// If no error found, return nil
	return nil
}

// ValidateSize checks if a chosen size is possible for a given Lineage
func ValidateSize(size string, lineage Lineage) error {
	if len(lineage.SizeOptions) == 1 {
		if size == lineage.SizeOptions[0] {
			return nil
		} else {
			return fmt.Errorf("size %s is not valid for %s", size, lineage.Name)
		}
	} else {
		for _, s := range lineage.SizeOptions {
			if s == size {
				return nil
			}
		}
		// didn't find it.
		return fmt.Errorf("size %s is not valid for %s", size, lineage.Name)
	}
}

// ValidateLevel makes sure that the level for the character is acceptable
func ValidateLevel(level int) error {
	if level < 1 {
		return errors.New("level must be greater than 0")
	}
	if level > 20 {
		return errors.New("level must be less than 21")
	}
	return nil
}

func (c *Character) GetTotalHitPoints() int {
	return c.CurrentHitPoints + c.TemporaryHitPoints
}

func (c *Character) GetHitPointBonusTotal() int {
	total := 0
	for _, bonus := range c.HitPointBonuses {
		total += bonus
	}
	c.TotalHitPointBonuses = total + c.Abilities.Modifiers["con"]
	return total
}

func (c *Character) AddHitPointsForLevel(nbrOfLevels int, sides int, startingLevel int) {
	Bonuses := c.GetHitPointBonusTotal()
	levelMessage := fmt.Sprintf("Character.AddHitPointsForLevel for levels %d through %d",
		startingLevel,
		(startingLevel-1)+nbrOfLevels)
	if nbrOfLevels == 1 {
		levelMessage = fmt.Sprintf("Character.AddHitPointsForLevel for level %d", startingLevel)
	}
	opts := []string{fmt.Sprintf("add %d", Bonuses*(nbrOfLevels))}
	results, err := dice.Perform(sides, nbrOfLevels, levelMessage, opts...)
	if err != nil {
		panic(err)
	}
	beforeCurrentHP := c.CurrentHitPoints
	beforeMaxHP := c.MaxHitPoints
	c.MaxHitPoints += results.Result
	c.CurrentHitPoints = c.MaxHitPoints
	c.History.Audits["MaxHitPoints"] = append(c.History.Audits["MaxHitPoints"],
		AuditEntry{
			Field:    "MaxHitPoints",
			OldValue: beforeMaxHP,
			NewValue: dice.Roll{
				ID:             results.ID,
				Options:        results.Options,
				Sides:          results.Sides,
				TimesToRoll:    results.TimesToRoll,
				RollsGenerated: results.RollsGenerated,
				AdditiveValue:  results.AdditiveValue,
				Result:         results.Result,
				RollsUsed:      results.RollsUsed,
				CtxRef:         results.CtxRef,
			},
			Source:    "Character.AddHitPointsForLevel",
			Timestamp: time.Now(),
		},
	)

	c.History.Audits["CurrentHitPoints"] = append(c.History.Audits["CurrentHitPoints"],
		AuditEntry{
			Field:    "CurrentHitPoints",
			OldValue: beforeCurrentHP,
			NewValue: dice.Roll{
				ID:             results.ID,
				Options:        results.Options,
				Sides:          results.Sides,
				TimesToRoll:    results.TimesToRoll,
				RollsGenerated: results.RollsGenerated,
				AdditiveValue:  results.AdditiveValue,
				Result:         results.Result,
				RollsUsed:      results.RollsUsed,
				CtxRef:         results.CtxRef,
			},
			Source:    "Character.AddHitPointsForLevel",
			Timestamp: time.Now(),
		},
	)
}

func (c *Character) InitHitPoints() {
	// hitPoints := 0
	sides := 0

	Bonuses := c.GetHitPointBonusTotal()

	// c.CurrentHitPointsAudit = []dice.Roll{}
	levelCounter := 0

	for i := range c.HitDice {
		toRoll := c.HitDice[i].DiceType
		switch toRoll {
		case "d4":
			sides = 4
		case "d6":
			sides = 6
		case "d8":
			sides = 8
		case "d10":
			sides = 10
		case "d12":
			sides = 12
		case "d20":
			sides = 20
		}
		if i == 0 {

			levelCounter += c.HitDice[i].Max
			c.MaxHitPoints = sides + Bonuses
			c.CurrentHitPoints = sides + Bonuses
			tmpID, err := helpers.GenerateRandomString(13)
			if err != nil {
				panic(err)
			}
			c.History.Audits["CurrentHitPoints"] = append(c.History.Audits["CurrentHitPoints"],
				AuditEntry{
					Field:    "CurrentHitPoints",
					OldValue: 0,
					NewValue: dice.Roll{
						ID:             tmpID,
						Options:        "",
						Sides:          sides,
						TimesToRoll:    1,
						RollsGenerated: []int{sides},
						AdditiveValue:  Bonuses,
						Result:         sides + Bonuses,
						RollsUsed:      []int{sides},
						CtxRef:         "First OverallLevel Auto Populate",
					},
					Source:    "Character.InitHitPoints",
					Timestamp: time.Now(),
				})
			if c.HitDice[i].Max > 1 {
				c.AddHitPointsForLevel(c.HitDice[i].Max-1, sides, 2)
			}
		} else {
			c.AddHitPointsForLevel(c.HitDice[i].Max, sides, levelCounter)
			levelCounter += c.HitDice[i].Max
		}
	}
}

// adjustDamageForType returns an adjusted amount for a character based on the damage type
func (c *Character) adjustDamageForType(data *DamageAudit) {
	value, exists := c.DamageTypeAdjustments[data.DamageType]
	if exists {

		switch value {
		case "vulnerable":
			data.Adjustments[data.DamageType] = data.BaseAmount
			data.TotalAmount = data.BaseAmount * 2
		case "resistant":
			data.Adjustments[data.DamageType] = (data.BaseAmount / 2) * -1
			data.TotalAmount = data.BaseAmount / 2
		case "immune":
			data.Adjustments[data.DamageType] = data.BaseAmount * -1
			data.TotalAmount = 0
		default:
			// No adjustment for normal damage
		}
	}
}

func (c *Character) Damage(amount int, damageType string) {
	tmpID, err := helpers.GenerateRandomString(13)
	if err != nil {
		panic(err)
	}

	audit := DamageAudit{
		ID:              tmpID,
		DamageType:      damageType,
		HitPointsBefore: c.GetTotalHitPoints(),
		BaseAmount:      amount,
		Adjustments:     make(map[string]int),
		TotalAmount:     amount,
	}

	c.adjustDamageForType(&audit)
	workingAmount := audit.TotalAmount
	if audit.TotalAmount > 0 {
		if c.TemporaryHitPoints >= workingAmount {
			c.TemporaryHitPoints -= workingAmount
			audit.Adjustments["temporary hit points"] = -workingAmount
			audit.TotalAmount = 0
			workingAmount = 0
		} else if c.TemporaryHitPoints > 0 {
			audit.Adjustments["temporary hit points"] = -c.TemporaryHitPoints
			workingAmount -= c.TemporaryHitPoints
			c.TemporaryHitPoints = 0
			audit.TotalAmount = workingAmount
		}
	}

	fmt.Printf("Damage: %d\n", workingAmount)

	if workingAmount > 0 {
		startingValue := c.CurrentHitPoints
		c.CurrentHitPoints -= workingAmount
		if c.CurrentHitPoints <= 0 {
			// Handle unconscious or death (not implemented here, but stub logic included)
			if -c.CurrentHitPoints >= c.MaxHitPoints {
				fmt.Println("Character has suffered instant death!")
				c.History.Audits["CurrentHitPoints"] = append(c.History.Audits["CurrentHitPoints"],
					AuditEntry{
						Field:     "CurrentHitPoints",
						OldValue:  fmt.Sprintf("%d/%d", startingValue, c.MaxHitPoints),
						NewValue:  fmt.Sprintf("%d < - %d = Instant Death", c.CurrentHitPoints, c.MaxHitPoints),
						Source:    "Character.Damage - Instant Death",
						Timestamp: time.Now(),
					},
				)
				c.CurrentHitPoints = 0
				c.MaxHitPoints = 0
				c.Conditions["dead"] = fmt.Sprintf("character suffered instant death at %v", time.Now().Format(time.RFC3339))
			} else {
				c.Conditions["unconscious"] = fmt.Sprintf("character fell unconscious at %v", time.Now().Format(time.RFC3339))

			}
		}

	}
	audit.HitPointsAfter = c.GetTotalHitPoints()
	c.History.DamageAudits = append(c.History.DamageAudits, audit)
}

func (c *Character) ModifyTemporaryHitPoints(amount int) {
	c.TemporaryHitPoints += amount
}

func (c *Character) UpdateAllDependencies() {
	for i := range c.Abilities.Base {
		c.UpdateDependencies(i)
	}
}

// AddClassLevel adds a level to a class and records the change
func (c *Character) AddClassLevel(className string, source string) {
	// Save old values
	oldLevel := c.OverallLevel
	oldClassLevel := c.CharacterLevels[className]

	// Update the class level
	c.CharacterLevels[className]++

	// Recalculate total level
	totalLevel := 0
	for _, level := range c.CharacterLevels {
		totalLevel += level
	}
	c.OverallLevel = totalLevel

	// Record class level audit
	classAudit := AuditEntry{
		Field:     "CharacterLevels",
		OldValue:  oldClassLevel,
		NewValue:  c.CharacterLevels[className],
		Source:    source,
		Timestamp: time.Now(),
	}
	c.History.Audits["CharacterLevels"] = append(c.History.Audits["CharacterLevels"], classAudit)

	// Record total level audit
	levelAudit := AuditEntry{
		Field:     "OverallLevel",
		OldValue:  oldLevel,
		NewValue:  c.OverallLevel,
		Source:    source,
		Timestamp: time.Now(),
	}
	c.History.Audits["OverallLevel"] = append(c.History.Audits["OverallLevel"], levelAudit)

	// Apply level-up benefits based on the new class level
	// todo: implement this through the classes themselves instead of
	//    writing this into the character class.
	// c.applyLevelBenefits(className, c.Classes[className])
}

// GetFieldHistory returns the audit history for a specific field
func (c *Character) GetFieldHistory(fieldName string) []AuditEntry {
	// Construct the audit field Name
	auditFieldName := fieldName + "Audit"

	// Use reflection to get the audit field
	v := reflect.ValueOf(c).Elem() // Get the Character struct value (dereference pointer)
	auditField := v.FieldByName(auditFieldName)

	// Check if the field exists
	if !auditField.IsValid() {
		return []AuditEntry{} // Return empty slice if field doesn't exist
	}

	// Convert the reflected value to []AuditEntry
	if auditField.Type().String() == "[]character.AuditEntry" {
		// Convert to interface then to concrete type
		auditEntries := auditField.Interface().([]AuditEntry)
		return auditEntries
	}

	// If we get here, the field exists but isn't the right type
	return []AuditEntry{}
}

// updateWithAudit is a helper to update any field with an audit trail
func (c *Character) updateWithAudit(
	fieldName string,
	oldValue interface{},
	newValue interface{},
	source string,
	auditSlice *[]AuditEntry) {

	// Create audit entry
	audit := AuditEntry{
		Field:     fieldName,
		OldValue:  oldValue,
		NewValue:  newValue,
		Source:    source,
		Timestamp: time.Now(),
	}

	// Append to the provided audit slice
	*auditSlice = append(*auditSlice, audit)
}

// InitializeAuditFields initializes all audit fields using reflection
func (c *Character) InitializeAuditFields() {
	v := reflect.ValueOf(c).Elem()
	t := v.Type()

	// Iterate through all fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// Check if this is an audit field (ends with "Audit")
		if strings.HasSuffix(field.Name, "Audit") && fieldValue.Kind() == reflect.Slice {
			// Make sure the field is exported (starts with uppercase)
			if field.Name[0] >= 'A' && field.Name[0] <= 'Z' {
				// Create a new empty slice of the appropriate type
				newSlice := reflect.MakeSlice(field.Type, 0, 0)

				// Set the field value to the new empty slice
				fieldValue.Set(newSlice)
			}
		}
	}
}

// ValidateLineage makes sure a valid lineage is used.
func ValidateLineage(lineageName string) bool {
	if lineageName != "" {
		_, err := GetLineageByName(lineageName)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func ValidateLanguages(languages []string) bool {
	for lang := range languages {
		if _, ok := Languages()[languages[lang]]; !ok {
			return false
		}
	}
	return true
}

// NewCharacter Method to create a new character with default properties.
//
//	rollingOptions:
//	  "predefined" - you manually rolled and are passing the values in.
//	  "common" - for each ability roll 4x, drop lowest
//	  "strict" - for each ability roll 3x
//	  "standard" - Use standard array - {15, 14, 13, 12, 10, 8}
//	  "pointbuy_even"     -  {13, 13, 13, 12, 12, 12}
//	  "pointbuy_onemax"   -  {15, 12, 12, 12, 11, 11}
//	  "pointbuy_twomax"   -  {15, 15, 11, 10, 10, 10}
//	  "pointbuy_threemax" -  {15, 15, 15, 8, 8, 8}
func NewCharacter(
	userId string, // user that created the character
	name string,
	level int,
	characterClassName string,
	selectedSubclassName string,
	lineageName string,
	heritageName string,
	backgroundName string,
	rollingOption string,
	chosenTraits map[string]string,
	chosenTalents []string,
	chosenLanguages []string,
	classBuildType string,
	manualBuildType ClassBuildType,
	description CharacterDescription,
	ctxRef string,
	logger *zap.SugaredLogger) (*Character, error) {

	zapLogger = logger
	useClass := Class{}
	useLineage := Lineage{}
	useHeritage := Heritage{}
	useBackground := Background{}
	id, err := helpers.GenerateRandomString(13)
	if err != nil {
		return nil, err
	}
	id = "pc" + id
	Audit := &HistoryAudit{
		CharacterId:  id,
		Audits:       make(map[string][]AuditEntry),
		DamageAudits: make([]DamageAudit, 0),
	}

	if len(name) == 0 {
		return nil, fmt.Errorf("name cannot be empty")
	}
	err = ValidateName(name)
	if err != nil {
		return nil, fmt.Errorf("name is invalid: %v", err)
	}
	err = ValidateLevel(level)
	if err != nil {
		return nil, fmt.Errorf("level is invalid: %v", err)
	}

	if rollingOption == "" {
		rollingOption = "standard"
	} else {
		err = ValidateRollingOption(rollingOption)
		if err != nil {
			return nil, fmt.Errorf("rolling option %s is invalid: %v", rollingOption, err)
		}
	}

	if characterClassName != "" {
		useClass, err = GetClassByName(characterClassName)
		if err != nil {
			return nil, fmt.Errorf("Error getting class '%s': %v\n", characterClassName, err)
		}
	} else {
		fmt.Println("No class specified. Using random selection instead.")
		useClass = RandomClass()
	}

	if classBuildType != "" {
		if !ValidateClassBuildType(classBuildType, useClass.ClassBuildTypes) {
			fmt.Printf("Class build type '%s' is invalid. Using Random Selection\n", classBuildType)
			classBuildType = RandomClassBuildType(useClass.ClassBuildTypes)
		}
	} else {
		if len(useClass.ClassBuildTypes) == 1 {
			fmt.Println("No class build type specified. Using Standard selection instead.")
			classBuildType = "Standard"
		} else {
			fmt.Println("No class build type specified. Using random selection instead.")
			classBuildType = RandomClassBuildType(useClass.ClassBuildTypes)
		}
	}
	classBuildInfo := useClass.ClassBuildTypes[classBuildType]

	selectedSubclass := Subclass{}
	implementedSubclass := Subclass{}
	subclass, err := useClass.GetSubclass(selectedSubclassName)
	if err != nil {
		fmt.Printf("The %s subclass is not valid for the %s class. Ignoring.\n",
			selectedSubclassName,
			characterClassName)
	} else {
		selectedSubclass = subclass
		if level >= 3 {
			implementedSubclass = selectedSubclass
		}
	}

	if ValidateLineage(lineageName) {
		useLineage, err = GetLineageByName(lineageName)
		if err != nil {
			return nil, fmt.Errorf("Error getting lineage '%s': %v\n", lineageName, err)
		}
	} else {
		return nil, fmt.Errorf("a valid lineage was not specified")
	}

	if heritageName != "" {
		useHeritage, err = GetHeritageByName(heritageName)
		if err != nil {
			return nil, fmt.Errorf("The %s Heritage is not valid.: %v\n", heritageName, err)
		}
	} else {
		return nil, fmt.Errorf("No Heritage specified. ")
	}

	if description.Size != "" {
		err = ValidateSize(description.Size, useLineage)
		if err != nil {
			return nil, fmt.Errorf("The %s size is not valid for %s: %v\n",
				description.Size, useLineage.Name, err)
		}
	} else {
		fmt.Println("No character size specified. Using random selection instead.")
		description.Size = RandomSize(useLineage)
	}

	useBackground, err = GetBackgroundByName(backgroundName)
	if err != nil {
		return nil, fmt.Errorf("'%s' is not a valid background: %v\n", backgroundName, err)
	}

	useLanguages := useHeritage.LanguageDefaults
	if len(chosenLanguages) > 0 {
		if ValidateLanguages(chosenLanguages) == false {
			return nil, fmt.Errorf("languages are invalid: %v", err)
		}
		useLanguages = chosenLanguages
	}

	for _, talent := range chosenTalents {
		_, ok := Talents[talent]
		if !ok {
			fmt.Printf("Could not find the talent: %s. Ignoring.\n", talent)
		}
	}

	AbilityScoreOrderPreference := useClass.ClassBuildTypes[classBuildType].AbilityScoreOrderPreference
	BonusArray := BonusArrayTemplate()

	// It would be a good idea to walk the Talents slice for changes to the ability bonuses before getting the account

	a, err := GetAbilityArray(rollingOption, AbilityScoreOrderPreference, BonusArray,
		ctxRef, false, logger)

	if err != nil {
		return nil, fmt.Errorf("failed to get ability array: %v", err)
	}

	hd := []HitDie{
		{
			SourceClass: characterClassName,
			DiceType:    useClass.HitDie,
			Max:         level,
			Used:        0,
		},
	}

	// KnownLanguages := useHeritage.LanguageDefaults

	var useAbilityScoreOrderPreference []string
	var useKeyAbilities []string
	if val, ok := useClass.ClassBuildTypes[classBuildType]; ok {
		useAbilityScoreOrderPreference = val.AbilityScoreOrderPreference
		useKeyAbilities = val.KeyAbilities
	} else {
		useAbilityScoreOrderPreference = manualBuildType.AbilityScoreOrderPreference
		useKeyAbilities = manualBuildType.KeyAbilities
	}

	// saves a pointer to the description struct
	Description := &description

	character := &Character{
		UserId:                       userId,
		ID:                           id,
		Name:                         name,
		Description:                  Description,
		OverallLevel:                 level,
		CharacterClassStr:            characterClassName,
		CharacterClassBuildType:      classBuildInfo,
		CharacterSubClassToImplement: selectedSubclass,
		CharacterSubClass:            implementedSubclass,
		Background:                   useBackground,
		HitDice:                      hd,
		DamageTypeAdjustments:        make(map[string]string),
		HitPointBonuses:              make(map[string]int),
		TemporaryHitPoints:           0,
		ProficiencyBonusBonus:        make(map[string]AbilitySkillBonus),
		MovementBase:                 Movement(float64(useLineage.Speed)),
		MovementBonus:                InitMovementBonus(),
		Lineage:                      useLineage,
		KnownLanguages:               useLanguages,
		Heritage:                     useHeritage,
		Traits:                       chosenTraits,
		Abilities:                    *a,
		Talents:                      map[string]Talent{},
		SpellcastingAbility:          string(useClass.SpellcastingAbility),
		AbilityScoreOrderPreference:  useAbilityScoreOrderPreference,
		KeyAbilities:                 useKeyAbilities,
		History:                      Audit,
	}
	character.InitializeAuditFields()
	character.SetAbilitySkills()
	character.SetAbilitySaveModifiers()
	character.CalculateMovement()
	character.UpdateAllDependencies()
	character.InitHitPoints()

	return character, nil
}

// PrintDetails prints detailed information about the character
func (c *Character) PrintDetails() {
	abilityOrder := []string{"str", "dex", "con", "int", "wis", "cha"}
	fmt.Printf("Character: %s\n", c.Name)
	fmt.Printf("Class: %s\n", c.CharacterClassStr)
	fmt.Printf("Subclass To Implement: %s\n", c.CharacterSubClassToImplement.Name)
	fmt.Printf("Subclass: %s\n", c.CharacterSubClass.Name)
	fmt.Printf("OverallLevel: %d\n", c.OverallLevel)
	fmt.Printf("Proficiency Bonus: %d\n", c.GetProficiencyBonus())
	fmt.Printf("Initiative Bonus: %d\n", c.InitiativeBonus)
	fmt.Printf("Passive Insight: %d\n", c.PassiveInsight)
	fmt.Printf("Passive Investigation: %d\n", c.PassiveInvestigation)
	fmt.Printf("Passive Perception: %d\n", c.PassivePerception)

	fmt.Printf("Rolling Option Used: %s\n", c.RollingOption)
	headerWidth := 4 // Length of the longest ability Name (e.g., "cha") + padding
	valueWidth := 4  // Padding for consistent alignment
	abilityHeaderStr := "            "
	abilityValueStr := "Values     "
	abilityModifierStr := "Modifiers  "
	abilitySaveModifierStr := "Save Mods  "
	for i := 0; i < 6; i++ {
		abilityHeaderStr += fmt.Sprintf("%-*s", headerWidth, abilityOrder[i])
		abilityValueStr += fmt.Sprintf("%*d", valueWidth, c.Abilities.Values[abilityOrder[i]])
		abilityModifierStr += fmt.Sprintf("%*d", valueWidth, c.Abilities.Modifiers[abilityOrder[i]])
		abilitySaveModifierStr += fmt.Sprintf("%*d", valueWidth, c.AbilitySaveModifiers[abilityOrder[i]])
	}
	fmt.Printf("Abilities:\n    %s\n", abilityHeaderStr)
	fmt.Printf("    %s\n", abilityValueStr)
	fmt.Printf("    %s\n", abilityModifierStr)
	fmt.Printf("    %s\n", abilitySaveModifierStr)

	fmt.Printf("Hit Dice:")
	for i := range c.HitDice {
		fmt.Printf(" %d@%s (%s)", c.HitDice[i].Max, c.HitDice[i].DiceType, c.HitDice[i].SourceClass)
	}
	fmt.Printf("\nMax Hit Points: %d\n", c.MaxHitPoints)
	fmt.Printf("Current Hit Points: %d\n", c.CurrentHitPoints)

	c.Lineage.PrintDetails()
	fmt.Printf("Chosen Size: %s\n", c.Description.Size)
	fmt.Println("Chosen Traits:")
	for traitType, trait := range c.Traits {
		fmt.Printf("  %s: %s\n", traitType, trait)
	}
	fmt.Printf("Heritage: %s, Languages: %v,  StaticTraits: %s, ChoiceTraits: %s\n",
		c.Heritage.Name, c.KnownLanguages, helpers.GetMapKeys(c.Heritage.Traits),
		helpers.GetMapKeys(c.Heritage.TraitOptions))

	fmt.Printf("Max Hit Points: %d\n", c.MaxHitPoints)
	fmt.Printf("Current Hit Points: %d\n", c.CurrentHitPoints)
	fmt.Printf("Total Hit Points: %d\n", c.GetTotalHitPoints())
	for x := range c.Talents {
		fmt.Printf("Talent: %s\n", c.Talents[x].Name)
		// fmt.Printf("  Prerequisite: %s\n", c.Talents[x].Prerequisite)
		//fmt.Printf("  Benefits: %v\n", c.Talents[x].Benefits)
	}
	fmt.Printf("Spell Book: %s\n", c.SpellBook)

	tmpStr := ""
	separator := ""
	if len(c.SkillProficiencies) > 0 {
		for x := range c.SkillProficiencies {
			tmpStr += fmt.Sprintf("%s%s (%s)", separator,
				c.SkillProficiencies[x].Skill, c.SkillProficiencies[x].Source)
			separator = ", "
		}
		fmt.Printf("\nAbility Proficiencies: %s\n", tmpStr)
	} else {
		fmt.Printf("\nNo Ability Proficiencies\n")
	}

	if len(c.SkillBonus) > 0 {
		tmpStr = ""
		for x := range c.SkillBonus {
			separator = ""
			tmpStr += fmt.Sprintf("\n%18s: ", x)
			for y := range c.SkillBonus[x] {
				tmpStr += fmt.Sprintf("%s%d (%s)",
					separator,
					c.SkillBonus[x][y].Bonus,
					c.SkillBonus[x][y].Source)
				separator = ", "
			}
		}
		fmt.Printf("Ability Bonus: %s\n\n", tmpStr)
	} else {
		fmt.Printf("\nNo Ability Bonus\n")
	}

	tmpStr = ""
	keys := make([]string, 0, len(c.AbilitySkills))
	for x := range c.AbilitySkills {
		keys = append(keys, x)
	}
	sort.Strings(keys)
	for _, key := range keys {
		proficientChar := "-"
		if c.AbilitySkills[key].Proficient {
			proficientChar = "+"
		}
		tmpStr += fmt.Sprintf("%s%18s (%3s): %4d %4d %4d\n",
			proficientChar, key, c.AbilitySkills[key].Ability,
			c.Abilities.Modifiers[c.AbilitySkills[key].Ability],
			c.TotalSkillModifiers[key],
			c.AbilitySkills[key].Value)
	}
	fmt.Printf("Ability Skills:\nP      Name (ability)       Base Bonus Total\n")
	fmt.Printf("---------------------------------------------\n%s\n", tmpStr)

	movementHeaderStr := fmt.Sprintf("%14s           %3s %12s    %s", "Type", "Base", "Bonus", "Total")

	fmt.Printf("Movement:\n  %s\n", movementHeaderStr)

	outputSlice := make(map[string]string, len(c.MovementBase))
	movementKeys := make([]string, 0, len(c.MovementBase))

	for x := range c.MovementBase {
		movementKeys = append(movementKeys, x)
	}
	sort.Strings(movementKeys)
	for x := range c.MovementBase {
		separator := ""
		tmpStr := "("
		for y := range c.MovementBonus[x] {
			if c.MovementBonus[x][y].Speed == 0 {
				tmpStr += fmt.Sprintf("%s%d", separator, c.MovementBonus[x][y].Speed)
			} else {
				tmpStr += fmt.Sprintf("%s%d [%s]", separator, c.MovementBonus[x][y].Speed, y)
			}
			separator = " + "
		}
		tmpStr += ")"
		outputSlice[x] = fmt.Sprintf("%18s: %6d + %12s = %3d",
			x,
			c.MovementBase[x].Speed,
			tmpStr,
			c.TotalMovement[x].Speed)
	}
	fmt.Printf("    %s\n", outputSlice["walking"])
	for _, value := range movementKeys {
		if value != "walking" {
			fmt.Printf("    %s\n", outputSlice[value])
		}
	}
	fmt.Printf("\nHit Point Audit:\n")
	for _, value := range c.History.Audits["CurrentHitPoints"] {
		fmt.Printf("%v %v %s\n", value.OldValue,
			value.NewValue, value.Source)
	}

}
