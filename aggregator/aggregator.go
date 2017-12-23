package aggregator

import (
	"sort"
	"strings"
)

// Aggregator computes statistics about arbitrary natural language.
type Aggregator struct {
	words        []string
	wordCounts   map[string]int
	letterCounts map[string]int
}

// Map of a key to a value.
type Map struct {
	key   string
	value int
}

// Mappings implements sort.Interface.
type Mappings []Map

// Len implementation.
func (m Mappings) Len() int {
	return len(m)
}

// Less implementation.
func (m Mappings) Less(i, j int) bool {
	return m[i].value < m[j].value
}

// Swap implementation.
func (m Mappings) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// NewAggregator creates a new Aggregator instance.
func NewAggregator() *Aggregator {
	w := make([]string, 0)
	wc := make(map[string]int)
	lc := make(map[string]int)
	return &Aggregator{
		words:        w,
		wordCounts:   wc,
		letterCounts: lc,
	}
}

// IngestWords processes the given text and updates the occurence of each word seen.
func (a *Aggregator) IngestWords(text string) {
	lc := strings.ToLower(text)
	words := strings.Fields(lc)
	a.updateWordOccurence(words)
	a.updateLetterOccurence(words)
}

// updateWordOccurence updates the number of times a word has been seen.
func (a *Aggregator) updateWordOccurence(words []string) {
	for _, w := range words {
		a.wordCounts[w]++
	}
}

// updateLetterOccurence updates the number of times a letter has been seen.
func (a *Aggregator) updateLetterOccurence(words []string) {
	for _, w := range words {
		letters := strings.Split(w, "")
		for _, l := range letters {
			a.letterCounts[l]++
		}
	}
}

// GetWordCounts returns the word counts.
func (a *Aggregator) GetWordCounts() map[string]int {
	return a.wordCounts
}

// GetLetterCounts returns the letter counts.
func (a *Aggregator) GetLetterCounts() map[string]int {
	return a.letterCounts
}

// TotalWords returns the total number of words seen.
func (a *Aggregator) TotalWords() int {
	total := 0
	for _, c := range a.wordCounts {
		total = total + c
	}

	return total
}

// Top5Words returns up to 5 of the most seen words.
func (a *Aggregator) Top5Words() []string {
	if len(a.wordCounts) == 0 {
		return nil
	}

	maps := make(Mappings, len(a.wordCounts))
	i := 0
	for w, c := range a.wordCounts {
		maps[i] = Map{w, c}
		i++
	}

	sort.Sort(maps)
	takeFrom := 0
	if len(maps) >= 5 {
		takeFrom = len(maps) - 5
	}

	top5 := make([]string, 0)
	for i := takeFrom; i < len(maps); i++ {
		top5 = append(top5, maps[i].key)
	}

	return top5
}

// Top5Letters returns up to 5 of the most seen letters.
func (a *Aggregator) Top5Letters() []string {
	if len(a.letterCounts) == 0 {
		return nil
	}

	maps := make(Mappings, len(a.letterCounts))
	i := 0
	for l, c := range a.letterCounts {
		maps[i] = Map{l, c}
		i++
	}

	sort.Sort(maps)
	takeFrom := 0
	if len(maps) >= 5 {
		takeFrom = len(maps) - 5
	}

	top5 := make([]string, 0)
	for i := takeFrom; i < len(maps); i++ {
		top5 = append(top5, maps[i].key)
	}

	return top5
}
