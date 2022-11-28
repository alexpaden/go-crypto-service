package util

import (
	"log"
	"math/big"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
)

// retrieves env variables from ./.env file
func goGetDotEnv(key string) string {
	println()
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// creates an infura connection string by chainId
func InfuraStringMaker(chainId int) string {
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
		log.Fatalf("requested chainId %v is not supported on infura", chainId)
	}

	return url + ".infura.io/v3/" + goGetDotEnv("INFURA_KEY")
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
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
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
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

	return result
}
