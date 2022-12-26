package balances

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Panicf("Error loading .env file")
	}
}

func TestRetrieveSingleBal(t *testing.T) {
	t.Parallel()
	addresses := []string{
		"0x0000000000000000000000000000000000000000",
		"0x000000000000000000000000000000000000001",
		"0x00000000000000000000000000000000000000002",
		"0x0000000000000000000000000000000000000000",
	}

	{
		wallet, err := RetrieveSingleBal(addresses[0], 1)
		if wallet == nil || err != nil {
			t.Errorf("Unexpected error for address %s", addresses[0])
		}
	}

	{
		wallet, err := RetrieveSingleBal(addresses[1], 1)
		if wallet != nil || err == nil {
			t.Errorf("Expected error for address %s", addresses[1])
		}
	}

	{
		wallet, err := RetrieveSingleBal(addresses[2], 1)
		if wallet != nil || err == nil {
			t.Errorf("Expected error for address %s", addresses[2])
		}
	}

	{
		// invalid chainId
		wallet, err := RetrieveSingleBal(addresses[3], 2)
		if wallet != nil || err == nil {
			t.Errorf("Expected error for address %s", addresses[2])
		}
	}
}

func TestRetrieveTokenBal(t *testing.T) {
	address := "0x0000000000000000000000000000000000000000"
	chainId := 1
	chainIdF := 137
	tokenF := "0x0000000000000000000000000000000000000000"
	tokenS := "0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"

	t.Parallel()
	{
		wallet, err := RetrieveTokenBal(address, chainId, tokenF)
		if wallet != nil || err == nil {
			t.Errorf("Expected error for  %s", tokenF)
		}
	}

	{
		wallet, err := RetrieveTokenBal(address, chainId, tokenS)
		if wallet == nil || err != nil {
			t.Errorf("Didn't expect error for token %s", tokenS)
		}
	}

	{
		wallet, err := RetrieveTokenBal(address, chainIdF, tokenS)
		if wallet != nil || err == nil {
			t.Errorf("Expected error for chainId %d", chainIdF)
		}
	}
}

func TestRetrieveManyBalances(t *testing.T) {
	t.Parallel()
	addresses := []string{
		"0x0000000000000000000000000000000000000000",
		"0x000000000000000000000000000000000000001",
		"0x00000000000000000000000000000000000000002",
	}

	{
		wallet, err := RetrieveManyBals(addresses[0])
		if wallet == nil || err != nil {
			t.Errorf("Unexpected error for address %s", addresses[0])
		}
	}

	{
		wallet, err := RetrieveManyBals(addresses[1])
		if wallet != nil || err == nil {
			t.Errorf("Unexpected error for address %s", addresses[1])
		}
	}

	{
		wallet, err := RetrieveManyBals(addresses[2])
		if wallet != nil || err == nil {
			t.Errorf("Unexpected error for address %s", addresses[2])
		}
	}
}
