package character

import (
	"fmt"
	"math"
)

type Money struct {
	GoldPieces   int
	SilverPieces int
	CopperPieces int
}

func (m *Money) AddCoin(Amount int, Type string) error {
	switch Type {
	case "gold":
		m.GoldPieces += Amount
	case "silver":
		m.SilverPieces += Amount
	case "copper":
		m.CopperPieces += Amount
	default:
		return fmt.Errorf("invalid coin type: %s", Type)
	}
	return nil
}

func (m *Money) RemoveCoin(Amount int, Type string) error {
	switch Type {
	case "gold":
		if Amount > m.GoldPieces {
			return fmt.Errorf("not enough gold pieces")
		}
		m.GoldPieces -= Amount
	case "silver":
		if Amount > m.SilverPieces {
			requiredGold := int(math.Ceil((float64(Amount - m.SilverPieces)) / 10.0))
			if m.GoldPieces >= requiredGold {
				m.GoldPieces -= requiredGold
				m.SilverPieces += requiredGold * 10
				m.SilverPieces -= Amount
				return nil
			}
			return fmt.Errorf("not enough gold or silver pieces")
		}
		m.SilverPieces -= Amount
	case "copper":
		if Amount > m.CopperPieces {
			requiredSilver := int(math.Ceil((float64(Amount - m.CopperPieces)) / 10.0))
			if m.SilverPieces >= requiredSilver {
				m.SilverPieces -= requiredSilver
				m.CopperPieces += requiredSilver * 10
				m.CopperPieces -= Amount
				return nil
			}
			requiredGold := int(math.Ceil((float64(Amount - m.CopperPieces)) / 100.0))

			if m.GoldPieces >= requiredGold {
				m.GoldPieces -= requiredGold
				m.SilverPieces += (requiredGold * 10) - requiredSilver
				m.CopperPieces += requiredSilver * 10
				m.CopperPieces -= Amount
				return nil
			}
			return fmt.Errorf("not enough gold, silver, or copper pieces")
		}
		m.CopperPieces -= Amount
	default:
		return fmt.Errorf("invalid coin type: %s", Type)
	}
	return nil
}
