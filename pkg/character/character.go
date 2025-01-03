package character

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"regexp"
	"sort"
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
				returnValue := 0
				if c.AbilitySkills["investigation"].Proficient {
					returnValue += c.GetBaseProficiencyBonus()
				}
				return c.GetAbility("int") + returnValue
			},
		},
	},
	"wis": {
		Name:            "wis",
		DependentSkills: []string{"animal handling", "insight", "medicine", "perception", "survival"},
		DependentValues: map[string]func(*Character) int{
			"PassivePerception": func(c *Character) int {
				returnValue := 0
				if c.AbilitySkills["perception"].Proficient {
					returnValue += c.GetBaseProficiencyBonus()
				}
				return c.GetAbility("wis") + returnValue
			},
		},
	},
	"cha": {
		Name:            "cha",
		DependentSkills: []string{"deception", "intimidation", "performance", "persuasion"},
		DependentValues: map[string]func(*Character) int{
			"PassiveInsight": func(c *Character) int {
				returnValue := 0
				if c.AbilitySkills["insight"].Proficient {
					returnValue += c.GetBaseProficiencyBonus()
				}
				return c.GetAbility("cha") + returnValue
			},
		},
	},
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

// Character represents a character in the game
type Character struct {
	Name                         string
	Level                        int
	CharacterClass               string
	CharacterSubClassToImplement Subclass // store subclass in case the pc is < 3rd level
	CharacterSubClass            Subclass
	HitDice                      []HitDie
	Lineage                      Lineage
	Heritage                     Heritage
	Background                   string
	ChosenSize                   string
	ChosenTraits                 map[string]string
	BaseSkills                   map[string]int
	BaseSkillBonus               map[string]int
	Abilities                    AbilityArray
	AbilitySaveModifiers         map[string]int
	RollingOption                string
	MaxHitPoints                 int
	CurrentHitPoints             int
	InitiativeBonus              int
	PassiveInvestigation         int
	PassivePerception            int
	PassiveInsight               int
	Talents                      map[string]Talent
	SpellBook                    []string
	SkillProficiencies           []AbilitySkillProficiency
	SkillBonus                   map[string]map[string]AbilitySkillBonus
	TotalSkillModifiers          map[string]int
	MovementBase                 map[string]MovementValue
	MovementBonus                map[string]map[string]MovementValue
	TotalMovement                map[string]MovementValue
	AbilitySkills                map[string]AbilitySkill
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
	for i := range c.Abilities.Modifiers {
		c.AbilitySaveModifiers[i] = c.GetAbilityModifier(i)
		for _, proficiency := range Classes[c.CharacterClass].SaveProficiencies {
			if i == proficiency {
				// fmt.Printf("Proficiency: %s\n", proficiency)
				// fmt.Printf("PB: %d\n", c.GetBaseProficiencyBonus())
				c.AbilitySaveModifiers[i] += c.GetBaseProficiencyBonus()
			}
		}
	}

	// fmt.Printf("Ability Modifiers: %v\n", c.Abilities.Modifiers)
	// fmt.Printf("Ability Save Modifiers: %v\n", c.AbilitySaveModifiers)
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
		runningTotal += c.GetBaseProficiencyBonus()
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
			runningTotal += c.GetBaseProficiencyBonus()
		}

		for i := range c.SkillBonus[skill] {
			runningTotal += c.SkillBonus[skill][i].Bonus
		}

		abilitySkill.Value = runningTotal + c.Abilities.Modifiers[abilitySkill.Ability]
		c.AbilitySkills[skill] = abilitySkill
		c.TotalSkillModifiers[skill] = runningTotal
	}
}

func (c *Character) AddTalent(t Talent) error {
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
	c.Talents[t.Name] = t
	return nil
}

func (c *Character) GetBaseProficiencyBonus() int {
	return c.Level/4 + 2
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
	c.BaseSkillBonus[skillName] += int(float64(c.GetBaseProficiencyBonus()) * multiplier)
}

func (c *Character) AddAbilityBonus(ability string, reason string, bonus int) {
	c.Abilities.AdjustBonuses(ability, reason, bonus, zapLogger)

}

// ValidateName checks if the name is not empty and does not contain invalid characters.
func ValidateName(name string) error {

	// Ensure the name is not empty
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
	name string,
	level int,
	characterClass string,
	characterSubClass string,
	lineage Lineage,
	heritage Heritage,
	chosenSize string,
	rollingOption string,
	chosenTraits map[string]string,
	chosenTalents []string,
	classBuildType string,
	ctxRef string,
	logger *zap.SugaredLogger) (*Character, error) {

	zapLogger = logger
	useClass := Class{}

	err := error(nil)

	err = ValidateName(name)
	if err != nil {
		return nil, fmt.Errorf("name is invalid: %v", err)
	}

	if characterClass != "" {
		useClass, err = GetClassByName(characterClass)
		if err != nil {
			fmt.Printf("Error getting class '%s': %v\n", characterClass, err)
			fmt.Println("Using random selection instead.")
			useClass = RandomClass()
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

	selectedSubclass := Subclass{}
	implementedSubclass := Subclass{}
	subclass, err := useClass.GetSubclass(characterSubClass)
	if err != nil {
		fmt.Printf("The %s subclass is not valid for the %s class. Ignoring.\n",
			characterSubClass,
			characterClass)
	} else {
		selectedSubclass = subclass
		if level >= 3 {
			implementedSubclass = selectedSubclass
		}
	}

	testLineage := Lineage{}
	if lineage.Name == testLineage.Name {
		return nil, fmt.Errorf("The %s lineage is not valid.\n",
			lineage.Name)
	}

	for _, talent := range chosenTalents {
		_, ok := Talents[talent]
		if !ok {
			fmt.Printf("Could not find the talent: %s. Ignoring.\n", talent)
		}
	}
	// TODO: marry up chosen Talents with reasons to choose them!

	// TODO: apply any class modifiers here, like humans getting a talent, etc

	AbilityScoreOrderPreference := useClass.ClassBuildTypes[classBuildType].AbilityScoreOrderPreference
	BonusArray := BonusArrayTemplate()

	// It would be a good idea to walk the Talents slice for changes to the ability bonuses before getting the account

	a, err := GetAbilityArray(rollingOption, AbilityScoreOrderPreference, BonusArray,
		ctxRef, false, logger)

	if err != nil {
		return nil, fmt.Errorf("failed to get ability array: %v", err)
	}

	character := &Character{
		Name:                         name,
		Level:                        level,
		CharacterClass:               characterClass,
		CharacterSubClassToImplement: selectedSubclass,
		CharacterSubClass:            implementedSubclass,
		MovementBase:                 Movement(float64(lineage.Speed)),
		MovementBonus:                InitMovementBonus(),
		Lineage:                      lineage,
		Heritage:                     heritage,
		ChosenSize:                   chosenSize,
		ChosenTraits:                 chosenTraits,
		Abilities:                    *a,
		Talents:                      map[string]Talent{},
	}

	character.SetAbilitySkills()
	character.SetAbilitySaveModifiers()
	character.CalculateMovement()

	return character, nil
}

// PrintDetails prints detailed information about the character
func (c *Character) PrintDetails() {
	abilityOrder := []string{"str", "dex", "con", "int", "wis", "cha"}
	fmt.Printf("Character: %s\n", c.Name)
	fmt.Printf("Class: %s\n", c.CharacterClass)
	fmt.Printf("Subclass To Implement: %s\n", c.CharacterSubClassToImplement.Name)
	fmt.Printf("Subclass: %s\n", c.CharacterSubClass.Name)
	fmt.Printf("Level: %d\n", c.Level)
	c.Lineage.PrintDetails()
	fmt.Printf("Chosen Size: %s\n", c.ChosenSize)
	fmt.Println("Chosen Traits:")
	for traitType, trait := range c.ChosenTraits {
		fmt.Printf("  %s: %s\n", traitType, trait)
	}
	fmt.Printf("Heritage: %s, Languages: %v, Cultural Trait: %s\n",
		c.Heritage.Name, c.Heritage.Languages, c.Heritage.CulturalTraits)

	fmt.Printf("Rolling Option Used: %s\n", c.RollingOption)
	fmt.Printf("Max Hit Points: %d\n", c.MaxHitPoints)
	fmt.Printf("Current Hit Points: %d\n", c.CurrentHitPoints)
	for x := range c.Talents {
		fmt.Printf("Talent: %s\n", c.Talents[x].Name)
		// fmt.Printf("  Prerequisite: %s\n", c.Talents[x].Prerequisite)
		//fmt.Printf("  Benefits: %v\n", c.Talents[x].Benefits)
	}
	fmt.Printf("Spell Book: %s\n", c.SpellBook)
	headerWidth := 4 // Length of the longest ability name (e.g., "cha") + padding
	valueWidth := 4  // Padding for consistent alignment
	abilityHeaderStr := "            "
	abilityValueStr := "Values     "
	abilityModifierStr := "Modifiers  "
	abilitySaveModifierStr := "Save Mods  "
	for i := 0; i < 5; i++ {
		abilityHeaderStr += fmt.Sprintf("%-*s", headerWidth, abilityOrder[i])
		abilityValueStr += fmt.Sprintf("%*d", valueWidth, c.Abilities.Values[abilityOrder[i]])
		abilityModifierStr += fmt.Sprintf("%*d", valueWidth, c.Abilities.Modifiers[abilityOrder[i]])
		abilitySaveModifierStr += fmt.Sprintf("%*d", valueWidth, c.AbilitySaveModifiers[abilityOrder[i]])
	}

	fmt.Printf("Abilities:\n    %s\n", abilityHeaderStr)
	fmt.Printf("    %s\n", abilityValueStr)
	fmt.Printf("    %s\n", abilityModifierStr)
	fmt.Printf("    %s\n", abilitySaveModifierStr)
	// for x := range c.Abilities.Modifiers {
	// 	fmt.Printf("Ability: %s, Modifier: %d\n", x, c.Abilities.Modifiers[x])
	// }
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

}
