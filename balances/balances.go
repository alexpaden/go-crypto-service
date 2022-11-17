package balances

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Wallet struct {
	ADDRESS  string
	BALANCES []Balance
}

type Balance struct {
	CHAINID int
	BALANCE *big.Int
}

func Hello(test *string) {
	*test = "other"
	fmt.Println("within aftr " + *test)
}

func goGetDotEnv(key string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetDefaultBalance(address string, chainId int) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	//goerli, polygon-mainnet, polygon-mumbai, rinkeby, ropsten, kovan, mainnet
	client, err := ethclient.Dial(infuraStringMaker(1))
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: 1, BALANCE: balance})
	return wallet
}

func infuraStringMaker(chainId int) string {
	url := "https://"
	switch chainId {
	case 1:
		url = url + "mainnet"
	case 3:
		url = url + "ropsten"
	case 4:
		url = url + "rinkeby"
	case 5:
		url = url + "goerli"
	case 42:
		url = url + "kovan"
	case 137:
		url = url + "polygon-mainnet"
	case 80001:
		url = url + "polygon-mumbai"
	default:
		url = url + "mainnet"
	}

	return url + ".infura.io/v3/" + goGetDotEnv("INFURA_KEY")
}
