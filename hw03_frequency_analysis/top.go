package hw03frequencyanalysis

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

var (
	AllWordsRegexp          = regexp.MustCompile(`\S+`)
	SpecialCharactersRegexp = regexp.MustCompile(`(^[^\p{L}]+)|([^\p{L}]+$)`)
)

func Top10(text string) []string {
	stats := buildWordStats(text)

	return takeTopWords(stats, 10)
}

func buildWordStats(text string) map[string]int {
	stats := make(map[string]int)
	matches := AllWordsRegexp.FindAllString(text, -1)

	for _, match := range matches {
		word, err := cleanWord(match)
		if err != nil {
			continue
		}

		stats[word]++
	}

	return stats
}

func cleanWord(rawWord string) (string, error) {
	word := SpecialCharactersRegexp.ReplaceAllString(rawWord, "")

	if word == "" {
		return rawWord, errors.New("not a word")
	}

	word = strings.ToLower(word)

	return word, nil
}

func takeTopWords(stats map[string]int, max int) []string {
	top := make([]string, 0, max)
	allWords := make([]string, 0, len(stats))

	for k := range stats {
		allWords = append(allWords, k)
	}

	sort.SliceStable(allWords, func(i, j int) bool {
		if stats[allWords[i]] == stats[allWords[j]] {
			return allWords[i] < allWords[j]
		}

		return stats[allWords[i]] > stats[allWords[j]]
	})

	for i, k := range allWords {
		top = append(top, k)

		if (i + 1) == max {
			break
		}
	}

	return top
}
