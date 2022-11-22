package balances

import (
	"context"
	"log"
	"os"

	util "github.com/alexpaden/go-crypto-service/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ADDRESS  string
	BALANCES []Balance
}

type Balance struct {
	CHAINID int
	BALANCE decimal.Decimal
}

// retrieves env variables from ./.env file
func goGetDotEnv(key string) string {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// creates an infura connection string by chainId
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

// returns a wallet with default balance for given chain
func GetDefaultBalance(address string, chainId int) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	//goerli, polygon-mainnet, polygon-mumbai, rinkeby, ropsten, kovan, mainnet
	client, err := ethclient.Dial(infuraStringMaker(chainId))
	wei, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	eth := util.ToDecimal(wei, 18)
	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: eth})

	client.Close()

	return wallet
}

func GetAllBalances(address string) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	//goerli, polygon-mainnet, polygon-mumbai, rinkeby, ropsten, kovan, mainnet
	chainIds := []int{1, 5, 137}

	for _, chainId := range chainIds {
		client, err := ethclient.Dial(infuraStringMaker(chainId))
		wei, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}
		eth := util.ToDecimal(wei, 18)
		wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: eth})
	}

	return wallet
}

func GetBalanceToken(address string, chainId int, tokenAddress string) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	//goerli, polygon-mainnet, polygon-mumbai, rinkeby, ropsten, kovan, mainnet
	client, err := ethclient.Dial(infuraStringMaker(chainId))
	wei, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	eth := util.ToDecimal(wei, 18)
	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: eth})

	client.Close()

	return wallet
}
