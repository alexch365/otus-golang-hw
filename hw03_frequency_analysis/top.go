package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

func Top10(str string) []string {
	var topWords []string
	if str == "" {
		return topWords
	}

	wordCount := map[string]int{}
	words := regexp.MustCompile(`\s`).Split(str, -1)
	for _, word := range words {
		word = regexp.MustCompile(`(^['"“\-]|['"”!,.:;\-]+?$)`).ReplaceAllString(word, "")
		if word != "" {
			wordCount[strings.ToLower(word)]++
		}
	}

	type wf struct {
		Word      string
		Frequency int
	}
	var wordFrequencies []wf
	for word, count := range wordCount {
		wordFrequencies = append(wordFrequencies, wf{word, count})
	}
	sort.Slice(wordFrequencies, func(i, j int) bool {
		return wordFrequencies[i].Frequency > wordFrequencies[j].Frequency
	})

	for _, freq := range wordFrequencies[:10] {
		topWords = append(topWords, freq.Word)
	}

	return topWords
}
