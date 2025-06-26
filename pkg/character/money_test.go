package character

import (
	"testing"
)

func TestMoney_AddCoin(t *testing.T) {
	tests := []struct {
		name        string
		initial     Money
		amount      int
		coinType    string
		expected    Money
		expectError bool
	}{
		{
			name:        "Add gold coins",
			initial:     Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			amount:      3,
			coinType:    "gold",
			expected:    Money{GoldPieces: 8, SilverPieces: 10, CopperPieces: 20},
			expectError: false,
		},
		{
			name:        "Add silver coins",
			initial:     Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			amount:      5,
			coinType:    "silver",
			expected:    Money{GoldPieces: 5, SilverPieces: 15, CopperPieces: 20},
			expectError: false,
		},
		{
			name:        "Add copper coins",
			initial:     Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			amount:      15,
			coinType:    "copper",
			expected:    Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 35},
			expectError: false,
		},
		{
			name:        "Invalid coin type",
			initial:     Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			amount:      5,
			coinType:    "platinum",
			expected:    Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// IMPORTANT: Note that this test will fail with the current implementation
			// because AddCoin is using a value receiver instead of a pointer receiver
			// and won't modify the original Money struct
			m := tt.initial
			err := m.AddCoin(tt.amount, tt.coinType)

			if (err != nil) != tt.expectError {
				t.Errorf("AddCoin() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if err == nil && (m.GoldPieces != tt.expected.GoldPieces ||
				m.SilverPieces != tt.expected.SilverPieces ||
				m.CopperPieces != tt.expected.CopperPieces) {
				t.Errorf("AddCoin() = %+v, expected %+v", m, tt.expected)
			}
		})
	}
}

func TestMoney_RemoveCoin(t *testing.T) {
	tests := []struct {
		name        string
		initial     Money
		amount      int
		coinType    string
		expected    Money
		expectError bool
	}{
		{
			name:        "Remove gold coins with sufficient balance",
			initial:     Money{GoldPieces: 10, SilverPieces: 5, CopperPieces: 15},
			amount:      3,
			coinType:    "gold",
			expected:    Money{GoldPieces: 7, SilverPieces: 5, CopperPieces: 15},
			expectError: false,
		},
		{
			name:        "Remove gold coins with insufficient balance",
			initial:     Money{GoldPieces: 2, SilverPieces: 5, CopperPieces: 15},
			amount:      5,
			coinType:    "gold",
			expected:    Money{GoldPieces: 2, SilverPieces: 5, CopperPieces: 15},
			expectError: true,
		},
		{
			name:        "Remove silver coins with sufficient balance",
			initial:     Money{GoldPieces: 5, SilverPieces: 15, CopperPieces: 20},
			amount:      10,
			coinType:    "silver",
			expected:    Money{GoldPieces: 5, SilverPieces: 5, CopperPieces: 20},
			expectError: false,
		},
		{
			name:        "Remove silver coins with insufficient silver but sufficient gold",
			initial:     Money{GoldPieces: 5, SilverPieces: 5, CopperPieces: 20},
			amount:      10,
			coinType:    "silver",
			expected:    Money{GoldPieces: 4, SilverPieces: 5, CopperPieces: 20},
			expectError: false,
		},
		{
			name:        "Remove silver coins with insufficient balance",
			initial:     Money{GoldPieces: 0, SilverPieces: 5, CopperPieces: 20},
			amount:      10,
			coinType:    "silver",
			expected:    Money{GoldPieces: 0, SilverPieces: 5, CopperPieces: 20},
			expectError: true,
		},
		{
			name:        "Remove copper coins with sufficient balance",
			initial:     Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 30},
			amount:      15,
			coinType:    "copper",
			expected:    Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 15},
			expectError: false,
		},
		{
			name:        "Remove copper coins with insufficient copper but sufficient silver",
			initial:     Money{GoldPieces: 5, SilverPieces: 5, CopperPieces: 5},
			amount:      15,
			coinType:    "copper",
			expected:    Money{GoldPieces: 5, SilverPieces: 4, CopperPieces: 0},
			expectError: false,
		},
		{
			name:        "Remove copper coins with insufficient copper and silver but sufficient gold",
			initial:     Money{GoldPieces: 1, SilverPieces: 0, CopperPieces: 5},
			amount:      15,
			coinType:    "copper",
			expected:    Money{GoldPieces: 0, SilverPieces: 9, CopperPieces: 0},
			expectError: false,
		},
		{
			name:        "Remove copper coins with insufficient balance",
			initial:     Money{GoldPieces: 0, SilverPieces: 0, CopperPieces: 5},
			amount:      15,
			coinType:    "copper",
			expected:    Money{GoldPieces: 0, SilverPieces: 0, CopperPieces: 5},
			expectError: true,
		},
		{
			name:        "Invalid coin type",
			initial:     Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			amount:      5,
			coinType:    "platinum",
			expected:    Money{GoldPieces: 5, SilverPieces: 10, CopperPieces: 20},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// IMPORTANT: Note that this test will fail with the current implementation
			// because RemoveCoin is using a value receiver instead of a pointer receiver
			// and won't modify the original Money struct
			m := tt.initial
			err := m.RemoveCoin(tt.amount, tt.coinType)

			if (err != nil) != tt.expectError {
				t.Errorf("RemoveCoin() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if err == nil && (m.GoldPieces != tt.expected.GoldPieces ||
				m.SilverPieces != tt.expected.SilverPieces ||
				m.CopperPieces != tt.expected.CopperPieces) {
				t.Errorf("RemoveCoin() = %+v, expected %+v", m, tt.expected)
			}
		})
	}
}

func TestMoney_CurrencyConversion(t *testing.T) {
	// This test specifically checks the currency conversion scenarios

	t.Run("Converting gold to silver when removing silver", func(t *testing.T) {
		m := Money{GoldPieces: 1, SilverPieces: 5, CopperPieces: 0}
		err := m.RemoveCoin(10, "silver")

		if err != nil {
			t.Errorf("RemoveCoin() error = %v, expected nil", err)
			return
		}

		expected := Money{GoldPieces: 0, SilverPieces: 5, CopperPieces: 0}
		if m.GoldPieces != expected.GoldPieces ||
			m.SilverPieces != expected.SilverPieces ||
			m.CopperPieces != expected.CopperPieces {
			t.Errorf("RemoveCoin() = %+v, expected %+v", m, expected)
		}
	})

	t.Run("Converting silver to copper when removing copper", func(t *testing.T) {
		m := Money{GoldPieces: 0, SilverPieces: 1, CopperPieces: 5}
		err := m.RemoveCoin(10, "copper")

		if err != nil {
			t.Errorf("RemoveCoin() error = %v, expected nil", err)
			return
		}

		expected := Money{GoldPieces: 0, SilverPieces: 0, CopperPieces: 5}
		if m.GoldPieces != expected.GoldPieces ||
			m.SilverPieces != expected.SilverPieces ||
			m.CopperPieces != expected.CopperPieces {
			t.Errorf("RemoveCoin() = %+v, expected %+v", m, expected)
		}
	})

	t.Run("Converting gold to copper (via silver) when removing copper", func(t *testing.T) {
		m := Money{GoldPieces: 1, SilverPieces: 0, CopperPieces: 5}
		err := m.RemoveCoin(50, "copper")

		if err != nil {
			t.Errorf("RemoveCoin() error = %v, expected nil", err)
			return
		}

		expected := Money{GoldPieces: 0, SilverPieces: 5, CopperPieces: 5}
		if m.GoldPieces != expected.GoldPieces ||
			m.SilverPieces != expected.SilverPieces ||
			m.CopperPieces != expected.CopperPieces {
			t.Errorf("RemoveCoin() = %+v, expected %+v", m, expected)
		}
	})
}
