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
		balances := balances.GetDefaultBalance("0x71c7656ec7ab88b098defb751b7401b5f6d8976f", 1)
		c.IndentedJSON(http.StatusOK, balances)

	})

	router.Run("localhost:8080")
	//router.GET("/test", balances.Hello)
}
