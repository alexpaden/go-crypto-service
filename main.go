package main

import (
	"net/http"

	"github.com/alexpaden/go-crypto-service/balances"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	test := "test"
	router.GET("/balances", func(c *gin.Context) {
		balances.Hello(&test)
		balances := balances.GetBalances(c)
		c.IndentedJSON(http.StatusOK, balances)

	})

	router.Run("localhost:8080")
	//router.GET("/test", balances.Hello)
}
