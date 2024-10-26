package character

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func TestAbilityDescriptions(t *testing.T) {
	actual := AbilityDescriptions()
	assert.Equal(t, 6, len(actual))
}

func TestAbilityScoreModifier(t *testing.T) {
	actual := AbilityScoreModifier()
	assert.Equal(t, 30, len(actual))
	assert.Equal(t, 0, actual[0])
	assert.Equal(t, 10, actual[30])
}

func TestAbilityAssign(t *testing.T) {
	actual := AbilityAssign()
	actualKeys := GetAbilityRollingOptions()
	assert.Equal(t, 8, len(actual))
	assert.Equal(t, 8, len(actualKeys))
}

func TestAbilityScorePointCost(t *testing.T) {
	actual := AbilityScorePointCost()
	assert.Equal(t, 8, len(actual))
	assert.Equal(t, 2, actual[10])
}

func TestSkillAbilityLookup(t *testing.T) {
	actual := SkillAbilityLookup()
	assert.Equal(t, 18, len(actual))
	assert.Equal(t, "int", actual["history"])
}

func TestAbilityArrayTemplate(t *testing.T) {
	actual := AbilityArrayTemplate()
	assert.Equal(t, 6, len(actual))
	assert.Equal(t, 0, actual["cha"])
}

func TestGetPreGeneratedBaseAbilityArray(t *testing.T) {
	actual, sortOrder := GetPreGeneratedBaseAbilityArray([]int{18, 17, 16, 15, 14, 13})
	assert.Equal(t, 6, len(actual))
	expected := map[string]int{
		"str": 18,
		"dex": 17,
		"con": 16,
		"int": 15,
		"wis": 14,
		"cha": 13,
	}
	assert.Equal(t, expected, actual)
	expectedSortOrder := []string{
		"str",
		"dex",
		"con",
		"int",
		"wis",
		"cha",
	}
	assert.Equal(t, expectedSortOrder, sortOrder)
}

func TestGetBaseAbilityArray(t *testing.T) {
	// Given
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()

	sortOrder := []string{"dex", "con", "str",
		"cha", "wis", "int"}
	rollingOption := "standard"

	// When
	actual, r, err := GetBaseAbilityArray(sortOrder, rollingOption, observedLoggerSugared)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, 6, len(r))
	assert.Equal(t, []int{15, 14, 13, 12, 10, 8}, r)
	expected := map[string]int{
		"str": 13,
		"dex": 15,
		"con": 14,
		"int": 8,
		"wis": 10,
		"cha": 12,
	}
	assert.Equal(t, expected, actual)
	require.Equal(t, 1, observedLogs.Len())
}

func TestGetBaseAbilityArrayWithRolls(t *testing.T) {
	// Given
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	sortOrder := []string{"dex", "con", "str",
		"cha", "wis", "int"}
	rollingOption := "common"

	// When
	actual, r, err := GetBaseAbilityArray(sortOrder, rollingOption, observedLoggerSugared)

	// Then
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, len(r))
	assert.GreaterOrEqual(t, actual["dex"], actual["con"])
	assert.GreaterOrEqual(t, actual["con"], actual["str"])
	assert.GreaterOrEqual(t, actual["str"], actual["cha"])
	assert.GreaterOrEqual(t, actual["cha"], actual["wis"])
	assert.GreaterOrEqual(t, actual["wis"], actual["int"])
	require.Equal(t, 6, len(actual))
	require.Equal(t, 7, observedLogs.Len()) // 6 dice rolls and the sorted map
}

func TestGetPreGeneratedAbilityArray(t *testing.T) {
	Raw := []int{18, 17, 16, 15, 14, 13}
	ArchetypeBonus := AbilityArrayTemplate()
	ArchetypeBonus["cha"] = 2
	ArchetypeBonus["int"] = 1
	ArchetypeBonusIgnored := false
	LevelChangeIncrease := AbilityArrayTemplate()
	LevelChangeIncrease["dex"] = 2
	AdditionalBonus := AbilityArrayTemplate()
	AdditionalBonus["str"] = 2
	ctxRef := "TestGetPreGeneratedAbilityArray"
	isMonsterOrGod := false
	a := GetPreGeneratedAbilityArray(Raw, ArchetypeBonus,
		ArchetypeBonusIgnored, LevelChangeIncrease,
		AdditionalBonus, ctxRef, isMonsterOrGod)
	// fmt.Println(a.ToPrettyString())
	assert.Equal(t, 20, a.Values["str"])
	assert.Equal(t, 15, a.Values["cha"])
	assert.Equal(t, 19, a.Values["dex"])
	assert.Equal(t, 16, a.Values["int"])
	actual, _ := a.GetModifier("str")
	assert.Equal(t, 5, actual)
	actual, _ = a.GetModifier("int")
	assert.Equal(t, 3, actual)
}

func TestGetAbilityArray(t *testing.T) {
	// Given
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"
	sortOrder := []string{"dex", "con", "str",
		"cha", "wis", "int"}
	ArchetypeBonus := AbilityArrayTemplate()
	ArchetypeBonus["cha"] = 2
	ArchetypeBonus["int"] = 1
	ArchetypeBonusIgnored := true
	LevelChangeIncrease := AbilityArrayTemplate()
	LevelChangeIncrease["dex"] = 2
	AdditionalBonus := AbilityArrayTemplate()
	AdditionalBonus["str"] = 2
	ctxRef := "TestGetAbilityArray"
	isMonsterOrGod := true

	// When
	a, err := GetAbilityArray(rollingOption, sortOrder, ArchetypeBonus,
		ArchetypeBonusIgnored, LevelChangeIncrease, AdditionalBonus,
		ctxRef, isMonsterOrGod, observedLoggerSugared)

	// Then
	assert.Equal(t, nil, err)
	// fmt.Println(a.ToPrettyString())
	assert.Equal(t, 15, a.Values["str"])
	assert.Equal(t, 12, a.Values["cha"])
	assert.Equal(t, 17, a.Values["dex"])
	assert.Equal(t, 8, a.Values["int"])
	actual, _ := a.GetModifier("str")
	assert.Equal(t, 2, actual)
	actual, _ = a.GetModifier("int")
	assert.Equal(t, -1, actual)
	allLogs := observedLogs.All()

	ctxMap := allLogs[0].ContextMap()
	tStr, _ := ctxMap["rawValues"].(string)
	assert.Equal(t, "[15, 14, 13, 12, 10, 8]", tStr)
	tStr, _ = ctxMap["sortedValues"].(string)
	assert.Equal(t, "{\"Str\": 13, \"Dex\": 15, \"Con\": 14, "+
		"\"Int\":  8, \"Wis\": 10, \"Cha\": 12}", tStr)

	assert.Equal(t, "GetAbilityArray", allLogs[len(allLogs)-1].Message)

}

func TestAdjustValues(t *testing.T) {
	// Given
	observedZapCore, observedLogs := observer.New(zap.InfoLevel)
	observedLoggerSugared := zap.New(observedZapCore).Sugar()
	rollingOption := "standard"
	sortOrder := []string{"dex", "con", "str",
		"cha", "wis", "int"}
	ArchetypeBonus := AbilityArrayTemplate()
	ArchetypeBonusIgnored := false
	LevelChangeIncrease := AbilityArrayTemplate()
	AdditionalBonus := AbilityArrayTemplate()
	ctxRef := "TestAdjustValues"
	isMonsterOrGod := false

	a, err := GetAbilityArray(rollingOption, sortOrder, ArchetypeBonus,
		ArchetypeBonusIgnored, LevelChangeIncrease, AdditionalBonus,
		ctxRef, isMonsterOrGod, observedLoggerSugared)
	assert.Equal(t, nil, err)
	a.AdjustValues("ArchetypeBonus", "cha",
		2, observedLoggerSugared)
	a.AdjustValues("ArchetypeBonus", "int",
		1, observedLoggerSugared)
	assert.Equal(t, 14, a.Values["cha"])
	assert.Equal(t, 9, a.Values["int"])
	a.AdjustValues("LevelChangeIncrease", "dex",
		2, observedLoggerSugared)
	assert.Equal(t, 17, a.Values["dex"])
	a.AdjustValues("AdditionalBonus", "str",
		2, observedLoggerSugared)
	assert.Equal(t, 15, a.Values["str"])
	actual, _ := a.GetModifier("str")
	assert.Equal(t, 2, actual)
	actual, _ = a.GetModifier("int")
	assert.Equal(t, -1, actual)
	// fmt.Println(a.ToPrettyString())

	allLogs := observedLogs.All()
	assert.Equal(t, "AdjustValues", allLogs[len(allLogs)-1].Message)
}
