package character

import (
	"fmt"
	"go.uber.org/zap"
)

// Character represents a character in the game
type Character struct {
	Name           string
	Level          int
	CharacterClass string
	Lineage        Lineage
	Heritage       Heritage
	ChosenSize     string
	ChosenTraits   map[string]string
	BaseSkills     map[string]int
	BaseSkillBonus map[string]int
	Abilities      AbilityArray
	RollingOption  string
	HitPoints      int
	Talents        map[string]Talent
	SpellBook      []string
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

func (c *Character) GetBaseProficiencyBonus() float64 {
	return float64(c.Level/4 + 2)
}

func (c *Character) GetAbility(ability string) int {
	return c.Abilities.Values[ability]
}

func (c *Character) AddSkillBonusMultiplier(skillName string, multiplier float64) {
	c.BaseSkillBonus[skillName] += int(c.GetBaseProficiencyBonus() * multiplier)
}

func (c *Character) AddAbilityBonus(ability string, bonus int) {
	c.Abilities.AdditionalBonus[ability] += bonus
	c.Abilities.setValuesAndModifiers()
}

// NewCharacter Method to create a new character with default properties
func NewCharacter(
	name string,
	level int,
	characterClass string,
	lineage Lineage,
	heritage Heritage,
	chosenSize string,
	rollingOption string,
	chosenTraits map[string]string,
	classBuildType string,
	ctxRef string,
	logger *zap.SugaredLogger) *Character {

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

	AbilityScoreOrderPreference := useClass.ClassBuildTypes[classBuildType].AbilityScoreOrderPreference
	AdditionalBonus := AbilityArrayTemplate()

	// TODO: Implement this
	LevelChangeIncrease := AbilityArrayTemplate()

	// It would be a good idea to walk the Talents slice for changes to the ability bonuses before getting the account

	a, err := GetAbilityArray(rollingOption, AbilityScoreOrderPreference, LevelChangeIncrease,
		AdditionalBonus, ctxRef, false, logger)

	if err != nil {
		panic(err)
	}

	character := &Character{
		Name:           name,
		Level:          level,
		CharacterClass: characterClass,
		Lineage:        lineage,
		Heritage:       heritage,
		ChosenSize:     chosenSize,
		ChosenTraits:   chosenTraits,
		Abilities:      *a,
		Talents:        map[string]Talent{},
	}
	return character
}

// Initializes a default AbilityArray,
// Modify rolling options and other parameters based on your RPG rules.
func createDefaultAbilities() AbilityArray {
	// Consider default parameters, adjust or prompt user/UI for input
	rollingOption := "standard"                                     // Default to standard or based on user choice
	sortOrder := []string{"str", "dex", "con", "int", "wis", "cha"} // Customize order as needed

	// Creating a default AbilityArray; passing empty values as per your initial setup
	abilityArray, _ := GetAbilityArray(rollingOption, sortOrder, map[string]int{},
		map[string]int{}, "initial character creation", false, nil)
	return *abilityArray
}

// PrintDetails prints detailed information about the character
func (c *Character) PrintDetails() {
	fmt.Printf("Character: %s\n", c.Name)
	c.Lineage.PrintDetails()
	fmt.Printf("Chosen Size: %s\n", c.ChosenSize)
	fmt.Println("Chosen Traits:")
	for traitType, trait := range c.ChosenTraits {
		fmt.Printf("  %s: %s\n", traitType, trait)
	}
	fmt.Printf("Heritage: %s, Languages: %v, Cultural Trait: %s\n",
		c.Heritage.Name, c.Heritage.Languages, c.Heritage.CulturalTraits["City Navigation"])
}
