package api

import (
	"github.com/PLT875/word-statistics/aggregator"
	"github.com/gin-gonic/gin"
)

// Router returns the API engine.
func Router(agg *aggregator.Aggregator) *gin.Engine {
	router := gin.Default()
	router.GET("/stats", GetStatistics(agg))
	router.HandleMethodNotAllowed = true

	return router
}
