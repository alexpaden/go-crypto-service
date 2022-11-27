package balances

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/alexpaden/go-crypto-service/token"
	util "github.com/alexpaden/go-crypto-service/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct {
	ADDRESS  string
	BALANCES []Balance
}

// should be lowercase?
type Balance struct {
	CHAINID int
	BALANCE *big.Float
}

// returns a wallet with default balance for given chain
func GetDefaultBalance(address string, chainId int) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	//goerli, polygon-mainnet, polygon-mumbai, rinkeby, ropsten, kovan, mainnet
	client, err := ethclient.Dial(util.InfuraStringMaker(chainId))
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

func getMagicBal(address string, chainId int, ch chan Balance) {
	balance := GetDefaultBalance(address, chainId)
	ch <- balance.BALANCES[0]
}

func GetAllDefaultBalances(address string) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	chainIds := []int{1, 5, 137}

	ch := make(chan Balance)

	for _, chainId := range chainIds {
		go getMagicBal(address, chainId, ch)
	}
	balances := make([]Balance, len(chainIds))

	for i := range balances {
		balances[i] = <-ch
	}

	wallet.BALANCES = balances

	return wallet
}

/*
func GetAllDefaultBalances(address string) Wallet {
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	chainIds := []int{1, 5, 137}

	for _, chainId := range chainIds {
		client, err := ethclient.Dial(util.InfuraStringMaker(chainId))
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
*/

func GetTokenBalance(walletAddress string, chainId int, contractAddress string) Wallet {

	client, err := ethclient.Dial(util.InfuraStringMaker(chainId))
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

	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: balance})

	return wallet
}
