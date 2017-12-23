package api

import (
	"net/http"

	"github.com/PLT875/word-statistics/aggregator"
	"github.com/gin-gonic/gin"
)

// Statistics represents the metrics returned by the GET /stats endpoint.
type Statistics struct {
	Count       int      `json:"count"`
	Top5Words   []string `json:"top_5_words"`
	Top5Letters []string `json:"top_5_letters"`
}

// GetStatistics is a handler for GET /stats.
func GetStatistics(agg *aggregator.Aggregator) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		count := agg.TotalWords()
		wc := agg.Top5Words()
		lc := agg.Top5Letters()

		stats := Statistics{count, wc, lc}

		c.JSON(http.StatusOK, stats)
		return
	}
}
