package balances

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/alexpaden/go-crypto-service/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Wallet struct {
	ADDRESS  string
	BALANCES []Balance
}

type Balance struct {
	CHAINID  int
	BALANCE  *big.Float
	CONTRACT string
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

	//eth := util.ToDecimal(wei, 18)
	eth := new(big.Float)
	eth.SetString(wei.String())
	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: eth})

	client.Close()

	return wallet
}

func GetAllDefaultBalances(address string) Wallet {
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

		//eth := util.ToDecimal(wei, 18)
		eth := new(big.Float)
		eth.SetString(wei.String())
		wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: eth})
	}

	return wallet
}

func GetTokenBalance(walletAddress string, chainId int, contractAddress string) Wallet {

	client, err := ethclient.Dial(infuraStringMaker(chainId))
	tokenAddress := common.HexToAddress(contractAddress)
	address := common.HexToAddress(walletAddress)
	wallet := Wallet{
		ADDRESS: address.String(),
	}

	// create a new instance of the token contract bound to a specific deployed contract
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// call BalanceOf function that takes an address and returns a *big.Int
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	// name of token contract
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// symbol of token contract
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// num decimals of token contract
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"

	// convert wei to eth
	//balance := util.ToDecimal(bal, 18)
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	balance := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", balance) // "balance: 74605500.647409"

	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: balance, CONTRACT: contractAddress})

	return wallet
}
