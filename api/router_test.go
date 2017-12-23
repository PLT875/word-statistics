package api_test

import (
	"fmt"
	"testing"

	"github.com/PLT875/word-statistics/aggregator"
	"github.com/PLT875/word-statistics/api"
	"github.com/stretchr/testify/require"
)

func TestRoutes(t *testing.T) {
	require := require.New(t)
	// when a new router is created
	agg := aggregator.NewAggregator()
	router := api.Router(agg)
	routes := router.Routes()

	// then
	require.Equal(1, len(routes))
	require.Equal("GET", routes[0].Method, fmt.Sprintf("the method at index %d of the routes should be %s", 0, "GET"))
}
