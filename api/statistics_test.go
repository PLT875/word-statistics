package api_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PLT875/word-statistics/aggregator"
	"github.com/PLT875/word-statistics/api"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type StatisticsTestConfig struct {
	agg    *aggregator.Aggregator
	router *gin.Engine
}

func (stc *StatisticsTestConfig) Init() {
	stc.agg = aggregator.NewAggregator()
	stc.agg.IngestWords("aa aa aa aa aa aa bb bb bb bb bb")
	stc.agg.IngestWords("cc cc cc cc dd dd dd ee ee ff")
	stc.router = api.Router(stc.agg)
}

func TestGetStatistics(t *testing.T) {
	require := require.New(t)
	stc := new(StatisticsTestConfig)
	stc.Init()
	// given the API is up and running
	req, _ := http.NewRequest("GET", "/stats", bytes.NewReader([]byte{}))
	res := httptest.NewRecorder()

	// when a GET /stats request is made
	stc.router.ServeHTTP(res, req)

	// then
	require.Equal(http.StatusOK, res.Code, fmt.Sprintf("the response code should be %d", http.StatusOK))

	// and
	expectedJSON := `{"count":21,"top_5_words":["ee","dd","cc","bb","aa"],"top_5_letters":["e","d","c","b","a"]}`
	require.JSONEq(expectedJSON, res.Body.String(), fmt.Sprintf("the response body should be %s", expectedJSON))
}
