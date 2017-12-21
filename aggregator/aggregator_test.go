package aggregator_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/PLT875/word-statistics/aggregator"
	"github.com/stretchr/testify/require"
)

func TestIngestWords(t *testing.T) {
	require := require.New(t)
	a := aggregator.NewAggregator()
	// when words are ingested
	a.IngestWords("ipsum dolor sit sit")
	a.IngestWords("dolor sit")

	// then
	wc := a.GetWordCounts()
	ttWc := []struct {
		word  string
		count int
	}{
		{"ipsum", 1}, {"dolor", 2}, {"sit", 3},
	}

	for _, t := range ttWc {
		require.Equal(t.count, wc[t.word], fmt.Sprintf("the word count of %s should be %d", t.word, t.count))
	}

	// and
	lc := a.GetLetterCounts()
	ttLc := []struct {
		letter string
		count  int
	}{
		{"i", 4}, {"p", 1}, {"s", 4}, {"u", 1}, {"m", 1}, {"d", 2}, {"o", 4}, {"l", 2}, {"r", 2},
	}

	for _, t := range ttLc {
		require.Equal(t.count, lc[t.letter], fmt.Sprintf("the letter count of %s should be %d", t.letter, t.count))
	}

	// and
	require.Equal(6, a.TotalWords(), fmt.Sprintf("the total number of words seen should be %d", 6))
}

func TestTop5Words(t *testing.T) {
	require := require.New(t)
	a := aggregator.NewAggregator()
	// initial
	top5 := a.Top5Words()
	require.Nil(top5)

	// given words are ingested
	a.IngestWords("six five four three two one")
	a.IngestWords("six five four three two")
	a.IngestWords("six five four three")
	a.IngestWords("six five four")
	a.IngestWords("six five")
	a.IngestWords("six")

	// when top 5 words are requested
	top5 = a.Top5Words()

	// then
	expected := []string{"two", "three", "four", "five", "six"}
	require.True(reflect.DeepEqual(expected, top5), "the top 5 words should be %v", expected)
}
