package server

import (
	"log"
	"strconv"

	"github.com/alexpaden/go-crypto-service/pkg/balances"
	"github.com/gin-gonic/gin"
)

// Server is a struct that contains the router and the balances service
type Server struct {
	Router *gin.Engine
}

// NewServer creates a new server
func NewServer() *Server {
	// Create a new gin router
	router := gin.Default()

	// register "/balances" routes
	registerBalancesRoutes(router)

	return &Server{
		Router: router,
	}
}

func registerBalancesRoutes(router *gin.Engine) {
	balsGroup := router.Group("/balances")
	{
		balancesService := balances.NewService()
		balsGroup.GET("/:address", func(ctx *gin.Context) {
			address := ctx.Param("address")
			wallet, err := balancesService.RetrieveManyBals(address)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.IndentedJSON(200, wallet)
		})

		balsGroup.GET("/:address/:chainId", func(ctx *gin.Context) {
			address := ctx.Param("address")
			chainId := ctx.Param("chainId")
			chainIdInt := stringToInt(chainId)
			wallet, err := balancesService.RetrieveSingleBal(address, chainIdInt)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.IndentedJSON(200, wallet)
		})

		balsGroup.GET("/:address/:chainId/:contract", func(ctx *gin.Context) {
			address := ctx.Param("address")
			chainId := ctx.Param("chainId")
			chainIdInt := stringToInt(chainId)
			contract := ctx.Param("contract")
			wallet, err := balancesService.RetrieveTokenBal(address, chainIdInt, contract)
			if err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.IndentedJSON(200, wallet)
		})
	}
}

func stringToInt(stringT string) int {
	stringI, err := strconv.Atoi(stringT)
	if err != nil {
		log.Default()
	}
	return stringI
}

// Run starts the server
func (s *Server) Run() {
	s.Router.Run("localhost:8080")
}
