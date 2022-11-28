package util

import (
	"math/big"
	"testing"
)

func TestIsValidAddress(t *testing.T) {
	t.Parallel()
	validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	invalidAddress := "0xabc"
	invalidAddress2 := "323b5d4c32345ced77393b3530b1eed0f346429d"
	{
		got := IsValidAddress(validAddress)
		expected := true

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress2)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestToDecimal(t *testing.T) {
	t.Parallel()
	weiAmount := big.NewInt(0)
	weiAmount.SetString("20000000000000000", 10)
	ethAmount := ToDecimal(weiAmount, 18)
	f64, _ := ethAmount.Float64()
	expected := 0.02
	if f64 != expected {
		t.Errorf("%v does not equal expected %v", ethAmount, expected)
	}
}
