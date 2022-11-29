## About
This project retrieves Ethereum address balances through a GoLang API.

- localhost:8080 **/balances/:address** 
	-> Retrieve the default balances for (address) on many chains (1, 5, 137)
	
- localhost:8080 **/balances/:address/:chainId** 
	-> Retrieve the default balance of (address) on (chainId)
	
- localhost:8080 **/balances/:address/:chainId/:contract** 
	-> Retrieve the balance of (token) in (address) on (chainId)


## Prepare

This project currently runs on the Infura RPC, to use Infura:
copy the .env.example file to .env and insert your Infura API key.

  
## Run

  T
  

## Test

Tests can be performed with the "go test" command from within the package directory "go-crypto-services/balances"

- Test retrieve single balance using address and chainId
- Test retrieve token balance using address, chainId, and contract address (token)
- Test retrieve many balances by address