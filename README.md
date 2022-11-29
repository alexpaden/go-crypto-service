## About
This project retrieves Ethereum address balances through a GoLang API.

- localhost:8080 **/balances/:address** 
	-> Retrieve the default balances for (address) on many chains (1, 5, 137)
	
- localhost:8080 **/balances/:address/:chainId** 
	-> Retrieve the default balance of (address) on (chainId)
	
- localhost:8080 **/balances/:address/:chainId/:contract** 
	-> Retrieve the balance of (token) in (address) on (chainId)


## Prepare

Clone project repository

This project currently runs on the Infura RPC, to use Infura:
copy the **.env.example** file to **.env** and insert your Infura API key.

  
## Run

  From inside the **go-crypto-service** run **go run .**
  
  **Sample links**
  - 404 fail: http://localhost:8080/balances
  - 200 succ: http://localhost:8080/balances/0x71c7656ec7ab88b098defb751b7401b5f6d8976f
  - 200 succ: http://localhost:8080/balances/0x71c7656ec7ab88b098defb751b7401b5f6d8976f/1
  - 400 fail: http://localhost:8080/balances/0x71c7656ec7ab88b098defb751b7401b5f6d8976f/11
  - 200 succ: http://localhost:8080/balances/0x71c7656ec7ab88b098defb751b7401b5f6d8976f/1/0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0
  - 400 fail: http://localhost:8080/balances/0x71c7656ec7ab88b098defb751b7401b5f6d8976f/11/0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0


## Test

Tests can be performed with the "go test" command from within the package directory "go-crypto-services/balances"

- Test retrieve single balance using address and chainId
- Test retrieve token balance using address, chainId, and contract address (token)
- Test retrieve many balances by address