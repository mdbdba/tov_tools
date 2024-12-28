package character

import (
	"fmt"
	"go.uber.org/zap"
	"sort"
)

var zapLogger *zap.SugaredLogger

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

// Character represents a character in the game
type Character struct {
	Name                         string
	Level                        int
	CharacterClass               string
	CharacterSubClassToImplement Subclass // store subclass in case the pc is < 3rd level
	CharacterSubClass            Subclass
	Lineage                      Lineage
	Heritage                     Heritage
	Background                   string
	ChosenSize                   string
	ChosenTraits                 map[string]string
	BaseSkills                   map[string]int
	BaseSkillBonus               map[string]int
	Abilities                    AbilityArray
	RollingOption                string
	MaxHitPoints                 int
	CurrentHitPoints             int
	Talents                      map[string]Talent
	SpellBook                    []string
	AbilityProficiencies         []AbilitySkillProficiency
	AbilityBonus                 map[string]map[string]AbilitySkillBonus
	TotalAbilityBonus            map[string]int
	AbilitySkills                map[string]AbilitySkill
}

func (c *Character) SetAbilitySkills() {

	c.AbilitySkills = map[string]AbilitySkill{}
	c.TotalAbilityBonus = map[string]int{}
	// init
	for skill, ability := range SkillAbilityLookup() {
		c.AbilitySkills[skill] = AbilitySkill{
			Ability:    ability,
			Proficient: false,
			Value:      0,
		}
	}
	// set proficiencies
	for _, i := range c.AbilityProficiencies {
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
		// if skill == "acrobatics" {
		// 	fmt.Printf("modifier [%s]: %v\n", abilitySkill.Ability,
		// 		c.Abilities.Modifiers[abilitySkill.Ability])
		// 	fmt.Printf("runningTotal: %d\n", runningTotal)
		// }
		for i := range c.AbilityBonus[skill] {
			runningTotal += c.AbilityBonus[skill][i].Bonus

			// if skill == "acrobatics" {
			// 	fmt.Printf("runningTotal: %d\n", runningTotal)
			// }
		}

		abilitySkill.Value = runningTotal + c.Abilities.Modifiers[abilitySkill.Ability]
		c.AbilitySkills[skill] = abilitySkill
		c.TotalAbilityBonus[skill] = runningTotal
	}
	//fmt.Println(c.AbilitySkills)
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

func (c *Character) AddSkillBonusMultiplier(skillName string, multiplier float64) {
	c.BaseSkillBonus[skillName] += int(float64(c.GetBaseProficiencyBonus()) * multiplier)
}

func (c *Character) AddAbilityBonus(ability string, reason string, bonus int) {
	c.Abilities.AdjustBonuses(ability, reason, bonus, zapLogger)

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
	logger *zap.SugaredLogger) *Character {

	zapLogger = logger
	useClass := Class{}

	err := error(nil)

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
		panic(err)
	}

	character := &Character{
		Name:                         name,
		Level:                        level,
		CharacterClass:               characterClass,
		CharacterSubClassToImplement: selectedSubclass,
		CharacterSubClass:            implementedSubclass,
		Lineage:                      lineage,
		Heritage:                     heritage,
		ChosenSize:                   chosenSize,
		ChosenTraits:                 chosenTraits,
		Abilities:                    *a,
		Talents:                      map[string]Talent{},
	}

	return character
}

/*
// Initializes a default AbilityArray,
// Modify rolling options and other parameters based on your RPG rules.
func createDefaultAbilities() AbilityArray {
	// Consider default parameters, adjust or prompt user/UI for input
	rollingOption := "standard"                                     // Default to standard or based on user choice
	sortOrder := []string{"str", "dex", "con", "int", "wis", "cha"} // Customize order as needed

	// Creating a default AbilityArray; passing empty values as per your initial setup
	abilityArray, _ := GetAbilityArray(rollingOption, sortOrder, BonusArrayTemplate(),
		"initial character creation", false, nil)
	return *abilityArray
}
*/

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
	for i := 0; i < 5; i++ {
		abilityHeaderStr += fmt.Sprintf("%-*s", headerWidth, abilityOrder[i])
		abilityValueStr += fmt.Sprintf("%*d", valueWidth, c.Abilities.Values[abilityOrder[i]])
		abilityModifierStr += fmt.Sprintf("%*d", valueWidth, c.Abilities.Modifiers[abilityOrder[i]])

	}

	fmt.Printf("Abilities:\n    %s\n", abilityHeaderStr)
	fmt.Printf("    %s\n", abilityValueStr)
	fmt.Printf("    %s\n", abilityModifierStr)
	// for x := range c.Abilities.Modifiers {
	// 	fmt.Printf("Ability: %s, Modifier: %d\n", x, c.Abilities.Modifiers[x])
	// }
	tmpStr := ""
	separator := ""
	for x := range c.AbilityProficiencies {
		tmpStr += fmt.Sprintf("%s%s (%s)", separator,
			c.AbilityProficiencies[x].Skill, c.AbilityProficiencies[x].Source)
		separator = ", "
	}
	fmt.Printf("\nAbility Proficiencies: %s\n", tmpStr)

	tmpStr = ""
	for x := range c.AbilityBonus {
		separator = ""
		tmpStr += fmt.Sprintf("\n%18s: ", x)
		for y := range c.AbilityBonus[x] {
			tmpStr += fmt.Sprintf("%s%d (%s)",
				separator,
				c.AbilityBonus[x][y].Bonus,
				c.AbilityBonus[x][y].Source)
			separator = ", "
		}
	}
	fmt.Printf("Ability Bonus: %s\n\n", tmpStr)

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
			c.TotalAbilityBonus[key],
			c.AbilitySkills[key].Value)
	}
	fmt.Printf("Ability Skills:\nP      Name (ability)       Base Bonus Total\n")
	fmt.Printf("---------------------------------------------\n%s\n", tmpStr)

}
