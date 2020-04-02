package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	if str == "" {
		return []string{}
	}

	wordCount := map[string]int{}
	words := regexp.MustCompile(`\s`).Split(str, -1)
	regexpCompiled := regexp.MustCompile(`(^['"“\-]|['"”!,.:;\-]+?$)`)
	for _, word := range words {
		word = regexpCompiled.ReplaceAllString(word, "")
		if word != "" {
			wordCount[strings.ToLower(word)]++
		}
	}

	type wf struct {
		Word      string
		Frequency int
	}
	wordFrequencies := make([]wf, 0, len(wordCount))
	for word, count := range wordCount {
		wordFrequencies = append(wordFrequencies, wf{word, count})
	}
	sort.Slice(wordFrequencies, func(i, j int) bool {
		return wordFrequencies[i].Frequency > wordFrequencies[j].Frequency
	})

	topWords := make([]string, 0, 10)
	for _, freq := range wordFrequencies[:10] {
		topWords = append(topWords, freq.Word)
	}

	return topWords
}
