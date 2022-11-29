package balances

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
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
	}

	{
		wallet, _ := RetrieveSingleBal(addresses[0], 1)
		if wallet == nil {
			t.Errorf("Unexpected error for address %s", addresses[0])
		}
	}

	{
		wallet, err := RetrieveSingleBal(addresses[1], 1)
		if wallet != nil || err == nil {
			t.Errorf("Unexpected error for address %s", addresses[1])
		}
	}

	{
		wallet, err := RetrieveSingleBal(addresses[2], 1)
		if wallet != nil || err == nil {
			t.Errorf("Unexpected error for address %s", addresses[2])
		}
	}
}

func TestRetrieveTokenBal(t *testing.T) {

}
