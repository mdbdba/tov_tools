package dice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasic(t *testing.T) {
	assertions := assert.New(t)
	opts := []string{"drop lowest 1"}
	pSides := 6
	pTimesToRoll := 4
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Basic", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != 4 {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != 3 {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.Equal("drop lowest: 1; ", rollObj.Options)
	assertions.GreaterOrEqual(rollObj.Result, 3)
	assertions.LessOrEqual(rollObj.Result, 18)
}

func TestDoubleAdvantage(t *testing.T) {
	assertions := assert.New(t)
	opts := []string{"advantage", "advantage"}
	pSides := 20
	pTimesToRoll := 1
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Double Advantage", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != 2 {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != 1 {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.GreaterOrEqual(rollObj.RollsGenerated[0], rollObj.RollsGenerated[1])
	assertions.Equal("vantage: advantage; ", rollObj.Options)
	assertions.GreaterOrEqual(rollObj.Result, 1)
	assertions.LessOrEqual(rollObj.Result, 20)
}

func TestPerformDisadvantage(t *testing.T) {
	assertions := assert.New(t)
	opts := []string{"disadvantage"}
	pSides := 20
	pTimesToRoll := 1
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Disadvantage", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != 2 {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != 1 {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.LessOrEqual(rollObj.RollsGenerated[0], rollObj.RollsGenerated[1])
	assertions.Equal("vantage: disadvantage; ", rollObj.Options)
	assertions.GreaterOrEqual(rollObj.Result, 1)
	assertions.LessOrEqual(rollObj.Result, 20)
}

func TestDoubleDisadvantage(t *testing.T) {
	// should be exactly the same as disadvantage
	assertions := assert.New(t)
	opts := []string{"disadvantage", "disadvantage"}
	pSides := 20
	pTimesToRoll := 1
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Double Disadvantage", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != 2 {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != 1 {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.LessOrEqual(rollObj.RollsGenerated[0], rollObj.RollsGenerated[1])
	assertions.Equal("vantage: disadvantage; ", rollObj.Options)
	assertions.GreaterOrEqual(rollObj.Result, 1)
	assertions.LessOrEqual(rollObj.Result, 20)
}

func TestAdvDis(t *testing.T) {
	// should end up being normal
	assertions := assert.New(t)
	opts := []string{"advantage", "disadvantage"}
	pSides := 20
	pTimesToRoll := 1
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Adv Dis", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != 1 {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != 1 {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.Equal("vantage: normal; ", rollObj.Options)
	assertions.GreaterOrEqual(rollObj.Result, 1)
	assertions.LessOrEqual(rollObj.Result, 20)
}

func TestDisAdv(t *testing.T) {
	// should end up being normal
	assertions := assert.New(t)
	opts := []string{"disadvantage", "advantage"}
	pSides := 20
	pTimesToRoll := 1
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Dis Adv", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != 1 {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != 1 {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.Equal("vantage: normal; ", rollObj.Options)
	assertions.GreaterOrEqual(rollObj.Result, 1)
	assertions.LessOrEqual(rollObj.Result, 20)
}

func TestAdd(t *testing.T) {
	assertions := assert.New(t)
	opts := []string{"add 3"}
	pSides := 6
	pTimesToRoll := 5
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Add", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != pTimesToRoll {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != pTimesToRoll {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	likelyResult := 0
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		likelyResult += rollObj.RollsUsed[i]
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.Equal("add: 3; ", rollObj.Options)
	assertions.Equal(rollObj.Result, likelyResult+3)
	assertions.GreaterOrEqual(rollObj.Result, pTimesToRoll+3)
	assertions.LessOrEqual(rollObj.Result, (pTimesToRoll*pSides)+3)
}

func TestSubtract(t *testing.T) {
	assertions := assert.New(t)
	opts := []string{"subtract 3"}
	pSides := 6
	pTimesToRoll := 5
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Subtract", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != pTimesToRoll {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != pTimesToRoll {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	likelyResult := 0
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		likelyResult += rollObj.RollsUsed[i]
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.Equal("subtract: 3; ", rollObj.Options)
	assertions.Equal(rollObj.Result, likelyResult-3)
	assertions.GreaterOrEqual(rollObj.Result, pTimesToRoll-3)
	assertions.LessOrEqual(rollObj.Result, (pTimesToRoll*pSides)-3)
}

func TestMultiAdds(t *testing.T) {
	assertions := assert.New(t)
	opts := []string{"subtract 3", "add 10", "subtract 4"}
	pSides := 6
	pTimesToRoll := 5
	rollObj, err := Perform(pSides, pTimesToRoll, "Test Multi Adds", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(rollObj.ToPrettyString())

	assertions.Equal(pSides, rollObj.Sides)
	assertions.Equal(pTimesToRoll, rollObj.TimesToRoll)
	if len(rollObj.RollsGenerated) != pTimesToRoll {
		t.Errorf("wrong number of generated received: %d",
			len(rollObj.RollsGenerated))
	}
	for i := 0; i < len(rollObj.RollsGenerated); i++ {
		if rollObj.RollsGenerated[i] < 1 ||
			rollObj.RollsGenerated[i] > rollObj.Sides {
			t.Errorf("Value for RollsGenerated outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsGenerated[i], i)
		}
	}
	if len(rollObj.RollsUsed) != pTimesToRoll {
		t.Errorf("wrong number of used received: %d",
			len(rollObj.RollsGenerated))
	}
	likelyResult := 0
	for i := 0; i < len(rollObj.RollsUsed); i++ {
		likelyResult += rollObj.RollsUsed[i]
		if rollObj.RollsUsed[i] < 1 ||
			rollObj.RollsUsed[i] > rollObj.Sides {
			t.Errorf("Value for RollsUsed outside range(1-%d): %d "+
				"Position: %d", rollObj.Sides, rollObj.RollsUsed[i], i)
		}
	}
	assertions.Equal(3, rollObj.AdditiveValue)
	assertions.Equal(rollObj.Result, likelyResult+3)
	assertions.GreaterOrEqual(rollObj.Result, pTimesToRoll+3)
	assertions.LessOrEqual(rollObj.Result, (pTimesToRoll*pSides)+3)
}
