package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/alexpaden/go-crypto-service/balances"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/balances/:address/:chainId/:token", func(c *gin.Context) {
		address := c.Param("address")
		chainId, err := strconv.Atoi(c.Param("chainId"))
		if err != nil {
			log.Default().Println(err)
		}
		token := c.Param("token")
		wallet, err := balances.RetrieveTokenBal(address, chainId, token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, wallet)
		}
	})

	router.GET("/balances/:address/:chainId", func(c *gin.Context) {
		address := c.Param("address")
		chainId, err := strconv.Atoi(c.Param("chainId"))
		if err != nil {
			log.Default().Println(err)
		}
		wallet, err := balances.RetrieveSingleBal(address, chainId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, wallet)
		}
	})

	router.GET("/balances/:address", func(c *gin.Context) {
		address := c.Param("address")
		wallet, err := balances.RetrieveManyBalances(address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.IndentedJSON(http.StatusOK, wallet)
		}
	})

	router.Run("localhost:8080")
}
