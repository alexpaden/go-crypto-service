package balances

import (
	"context"
	"errors"
	"log"
	"math/big"
	"os"
	"regexp"

	"github.com/alexpaden/go-crypto-service/pkg/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ADDRESS  string    `json:"address"`
	BALANCES []Balance `json:"balances"`
}

type Balance struct {
	CHAINID int             `json:"chain_id"`
	BALANCE decimal.Decimal `json:"balance"`
}

type Service struct {
	RetrieveManyBals  func(address string) (*Wallet, error)
	RetrieveSingleBal func(address string, chainId int) (*Wallet, error)
	RetrieveTokenBal  func(address string, chainId int, contract string) (*Wallet, error)
	Test              func()
}

func NewService() *Service {
	return &Service{
		RetrieveManyBals:  RetrieveManyBals,
		RetrieveSingleBal: RetrieveSingleBal,
		RetrieveTokenBal:  RetrieveTokenBal}
}

func retrieveBal(address string, chainId int, ch chan Balance) {
	account := common.HexToAddress(address)
	client, err := ethclient.Dial(infuraStringMaker(chainId))
	if err != nil {
		log.Panic(err)
	}
	wei, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Panic(err)
	}
	client.Close()
	bal := weiToDecimalRounded(wei, 18)
	ch <- Balance{CHAINID: chainId, BALANCE: bal}
}

func RetrieveSingleBal(address string, chainId int) (*Wallet, error) {
	if !isValidAddress(address) {
		err := errors.New("invalid address")
		return nil, err
	}
	if !isValidChainId(chainId) {
		err := errors.New("invalid chainId")
		return nil, err
	}
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	// create channel to receive balance
	ch := make(chan Balance)
	go retrieveBal(address, chainId, ch)
	balance := <-ch
	wallet.BALANCES = append(wallet.BALANCES, balance)

	return &wallet, nil
}

func RetrieveManyBals(address string) (*Wallet, error) {
	if !isValidAddress(address) {
		err := errors.New("invalid address")
		return nil, err
	}
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}

	// loop channels to receive balances
	ch := make(chan Balance)
	chainIds := []int{1, 5, 137}
	for _, chainId := range chainIds {
		go retrieveBal(address, chainId, ch)
	}
	balances := make([]Balance, len(chainIds))
	for i := range balances {
		balances[i] = <-ch
	}
	wallet.BALANCES = balances

	return &wallet, nil
}

func RetrieveTokenBal(address string, chainId int, contract string) (*Wallet, error) {
	if !isValidAddress(address) {
		err := errors.New("invalid address")
		return nil, err
	}
	if !isValidChainId(chainId) {
		err := errors.New("invalid chainId")
		return nil, err
	}
	if !isValidAddress(contract) {
		err := errors.New("invalid token address")
		return nil, err
	}
	account := common.HexToAddress(address)
	wallet := Wallet{
		ADDRESS: account.String(),
	}
	tokenAddress := common.HexToAddress(contract)

	// create client and token to query
	client, err := ethclient.Dial(infuraStringMaker(chainId))
	if err != nil {
		log.Panic(err)
	}
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Panic(err)
	}

	// find wei balance and provided decimals
	wei, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		err := errors.New("error retrieving token balance, check contract address")
		return nil, err
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		err := errors.New("issue creating token instance")
		return nil, err
	}

	weiDecimals := int(decimals)
	client.Close()
	bal := weiToDecimalRounded(wei, weiDecimals)
	wallet.BALANCES = append(wallet.BALANCES, Balance{CHAINID: chainId, BALANCE: bal})

	return &wallet, nil
}

// -----------------------------
/*
 Insert Cool divider here
 This divides things nicely
 I like that
 Divide Away, my friend.
*/
// -----------------------------

func isValidChainId(chainId int) bool {
	switch chainId {
	case 1, 5, 42, 137, 80001:
		return true
	default:
		return false
	}
}

// creates an infura connection string by chainId
func infuraStringMaker(chainId int) string {
	url := "https://"
	switch chainId {
	case 1:
		url = url + "mainnet"
	case 5:
		url = url + "goerli"
	case 42:
		url = url + "kovan"
	case 137:
		url = url + "polygon-mainnet"
	case 80001:
		url = url + "polygon-mumbai"
	default:
		log.Panicf("requested chainId %v is not supported on infura, try {1, 5, 137}", chainId)
	}

	return url + ".infura.io/v3/" + os.Getenv("INFURA_KEY")
}

// IsValidAddress validate hex address
func isValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// ToDecimal wei to decimals
func weiToDecimalRounded(ivalue interface{}, decimals int) decimal.Decimal {
	decimal.MarshalJSONWithoutQuotes = true
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	result = result.RoundDown(3)
	return result
}
