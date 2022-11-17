package balances

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type wallet struct {
	ADDRESS  string    `json:"walletAddress"`
	BALANCES []balance `json:"balances"`
}

type balance struct {
	CHAINID int `json:"chain_id"`
	Balance int `json:"balance"`
}

var balances = []balance{
	{CHAINID: 1, Balance: 0.00},
	{CHAINID: 5, Balance: 0.00},
	{CHAINID: 137, Balance: 0.00},
}

func Hello(test *string) {
	*test = "other"
	fmt.Println("within aftr " + *test)
}

func GetBalances(c *gin.Context) []balance {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/")
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	//goerli
	client2, err := ethclient.Dial("https://polygon-mainnet.infura.io/v3/")

	balance2, err := client2.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	balances[0].Balance = int(balance.Int64())
	balances[2].Balance = int(balance2.Int64())

	return balances
}
