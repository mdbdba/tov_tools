package dice

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"tov_tools/pkg/helpers"
	"tov_tools/pkg/logging"
)

// Roll struct containing everything you wanted to know about a roll
type Roll struct {
	Options        string
	Sides          int
	TimesToRoll    int
	RollsGenerated []int
	RollsUsed      []int
	AdditiveValue  int
	Result         int
	CtxRef         string
}

func (r *Roll) ToJson() string {
	j, err := json.Marshal(r)
	if err != nil {
		panic("Issue converting Roll to json object")
	}
	return string(j)
}

func (r *Roll) ToPrettyString() string {
	return r.ConvertToString(true)
}

func (r *Roll) ToString() string {
	return r.ConvertToString(false)
}

func (r *Roll) ConvertToString(p bool) (s string) {
	usedStr := helpers.IntSliceToString(r.RollsUsed)
	genStr := helpers.IntSliceToString(r.RollsGenerated)
	pStr := ""
	if p {
		pStr = "\n\t"
	}
	s = fmt.Sprintf("ROLL -- %sSides: %d, %sTimesToRoll: %d, "+
		"%sOptions: [%s], %sAdditiveValue: %d, %sResult: %d, %sRollsUsed: %s, "+
		"%sRollsGenerated: %s, %sCtxRef: %s\n",
		pStr, r.Sides,
		pStr, r.TimesToRoll,
		pStr, strings.TrimSpace(r.Options),
		pStr, r.AdditiveValue,
		pStr, r.Result,
		pStr, usedStr,
		pStr, genStr,
		pStr, r.CtxRef,
	)
	return
}

func getRolls(sides int, timesToRoll int) (*[]int, error) {
	var rolls []int
	for i := 0; i < timesToRoll; i++ {
		value, err := rand.Int(rand.Reader, big.NewInt(int64(sides)))
		if err != nil {
			return nil, err
		}
		t := int(value.Int64()) + 1 // +1 because dice start at 1, not 0
		rolls = append(rolls, t)
	}
	return &rolls, nil
}

// Perform - Internal perform function to handle the core logic.
// Options
//
//	[keep | drop] [highest | lowest] timesToRoll
//	[advantage | disadvantage]
//	[add | subtract] value
//
// Expectations:
//   - advantage and disadvantage cancel each other out.
//     If advantage and disadvantage are both passed, then the
//     result will be a normal Roll.
//   - advantage and disadvantage do not stack.
//   - it is assumed that rolling with advantage or disadvantage the
//     number of rolls is = 1. If something other than 1 is passed in this
//     scenario an error will be returned.
//   - using the variadic function for the Options parameter will allow us
//     to simplify all the different combinations by just evaluating them here.
func Perform(sides int, timesToRoll int, CtxRef string, options ...string) (r *Roll, err error) {
	canonical := logging.New("Dice.Roll.Perform")
	var reqLogStr string // boil down all the Options to an easy-to-read string
	var vantageLogStr string
	var keepLogStr string
	var additiveLogStr string
	keepValue := timesToRoll // the total number of rolls to keep.
	evalValue := timesToRoll // the total number of rolls to evaluate.
	// e.g.  rolling with advantage/disadvantage evaluates 2 rolls but keep 1
	//       keep / drop will have eval & keep numbers that differ as well
	sortDirection := "descending"
	additiveValue := 0 // value to add or subtract from the result.
	vantageTrack := "normal"
	for _, opt := range options {
		optSlice := strings.Split(opt, " ")
		switch optSlice[0] {
		case "keep":
			if optSlice[1] == "highest" {
				sortDirection = "ascending"
			} else if optSlice[1] == "lowest" {
				sortDirection = "descending"
			} else {
				panic("Unrecognized string for which values to keep.")
			}
			keepValue, err = strconv.Atoi(optSlice[2])
			if err != nil {
				panic(err)
			}
			keepLogStr = fmt.Sprintf("keep %s: %d; ", optSlice[1], keepValue)
		case "drop":
			if optSlice[1] == "highest" {
				sortDirection = "ascending"
			} else if optSlice[1] == "lowest" {
				sortDirection = "descending"
			} else {
				panic("Unrecognized string for which values to keep.")
			}
			var tmpInt int
			tmpInt, err = strconv.Atoi(optSlice[2])
			if err != nil {
				panic(err)
			}

			keepLogStr = fmt.Sprintf("drop %s: %d; ", optSlice[1], tmpInt)
			if keepValue > tmpInt {
				keepValue -= tmpInt
			} else {
				panic("Tried to drop more rolls than requested.")
			}
		case "add":
			var tValue int
			tValue, err = strconv.Atoi(optSlice[1])
			if err != nil {
				panic(err)
			}
			additiveValue += tValue
			additiveLogStr = fmt.Sprintf("%sadd: %d; ", additiveLogStr, tValue)
		case "subtract":
			var tValue int
			tValue, err = strconv.Atoi(optSlice[1])
			additiveValue -= tValue
			if err != nil {
				panic(err)
			}

			additiveLogStr = fmt.Sprintf("%ssubtract: %d; ", additiveLogStr, tValue)
		case "advantage":
			if timesToRoll != 1 {
				panic("advantage cannot be used with multiple rolls")
			}
			if vantageTrack == "normal" {
				vantageTrack = "advantage"
				evalValue = 2
			} else if vantageTrack == "disadvantage" {
				vantageTrack = "normal"
				evalValue = 1
				sortDirection = "descending"
			}
			vantageLogStr = fmt.Sprintf("vantage: %s; ", vantageTrack)
		case "disadvantage":
			if timesToRoll != 1 {
				panic("disadvantage cannot be used with multiple rolls")
			}

			if vantageTrack == "normal" {
				vantageTrack = "disadvantage"
				sortDirection = "ascending"
				evalValue = 2
			} else if vantageTrack == "advantage" {
				vantageTrack = "normal"
				evalValue = 1
				sortDirection = "descending"
			}
			vantageLogStr = fmt.Sprintf("vantage: %s; ", vantageTrack)
		}
	}
	rolls, err := getRolls(sides, evalValue)
	if err != nil {
		panic(err)
	}
	if sortDirection == "descending" {
		helpers.SortDescendingIntSlice(*rolls)
	} else {
		helpers.SortAscendingIntSlice(*rolls)
	}
	usedSlice := *rolls
	if evalValue != keepValue {
		usedSlice = usedSlice[0:keepValue]
	}
	result := 0
	for i := 0; i < len(usedSlice); i++ {
		result = result + usedSlice[i]
	}
	result += additiveValue
	reqLogStr = fmt.Sprintf("%s%s%s", vantageLogStr, keepLogStr, additiveLogStr)

	RollObj := Roll{
		Options:        reqLogStr,
		Sides:          sides,
		TimesToRoll:    timesToRoll,
		RollsGenerated: *rolls,
		RollsUsed:      usedSlice,
		AdditiveValue:  additiveValue,
		Result:         result,
		CtxRef:         CtxRef,
	}

	logging.LogUnitOfWork(canonical, &RollObj, "Perform")
	return &RollObj, nil
}
