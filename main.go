package main

import (
	"net/http"
	"strconv"

	"github.com/alexpaden/go-crypto-service/balances"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/balances", func(c *gin.Context) {
		balances := balances.GetManyBalances("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
		c.IndentedJSON(http.StatusOK, balances)
	})

	router.GET("/balances/:address", func(c *gin.Context) {
		address := c.Param("address")
		balances := balances.GetManyBalances(address)
		c.IndentedJSON(http.StatusOK, balances)
	})

	router.GET("/balances/:address/:chainId", func(c *gin.Context) {
		address := c.Param("address")
		chainId, err := strconv.Atoi(c.Param("chainId"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ch := make(chan balances.Balance)
		go balances.GoGetSingleBal(address, chainId, ch)
		balance := <-ch

		//balance := balances.GetSingleBalance(address, chainId)
		c.IndentedJSON(http.StatusOK, balance)
	})

	router.GET("/balances/:address/:chainId/:token", func(c *gin.Context) {
		// 0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0 polygon contract
		// 0x7D38aE457a3E24E5aF60a637638e134c97e2a1d5 wallet address
		address := c.Param("address")
		chainId, err := strconv.Atoi(c.Param("chainId"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		token := c.Param("token")
		balance := balances.GetTokenBalance(address, chainId, token)
		c.IndentedJSON(http.StatusOK, balance)
	})

	router.Run("localhost:8080")
	//router.GET("/test", balances.Hello)
}
