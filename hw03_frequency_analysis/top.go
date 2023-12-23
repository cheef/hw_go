package hw03frequencyanalysis

import (
	"errors"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	Word  string
	Count int
}

var (
	AllWordsRegexp          = regexp.MustCompile(`\S+`)
	SpecialCharactersRegexp = regexp.MustCompile(`(^[^\p{L}]+)|([^\p{L}]+$)`)
)

func Top10(text string) []string {
	wordStats := buildWordStats(text)

	return takeTopWords(wordStats, 10)
}

func buildWordStats(text string) []wordCount {
	stats := make(map[string]int)
	matches := AllWordsRegexp.FindAllString(text, -1)

	for _, match := range matches {
		word, err := cleanWord(match)
		if err != nil {
			continue
		}

		if _, ok := stats[word]; !ok {
			stats[word] = 1
		} else {
			stats[word]++
		}
	}

	return convertMapToWordStats(stats)
}

func cleanWord(rawWord string) (string, error) {
	word := SpecialCharactersRegexp.ReplaceAllString(rawWord, "")

	if word == "" {
		return rawWord, errors.New("not a word")
	}

	word = strings.ToLower(word)

	return word, nil
}

func convertMapToWordStats(stats map[string]int) []wordCount {
	wordCounts := make([]wordCount, 0, len(stats))

	for word, count := range stats {
		wordCounts = append(wordCounts, wordCount{word, count})
	}

	return wordCounts
}

func takeTopWords(wordStats []wordCount, max int) []string {
	top := make([]string, 0, max)

	sort.Slice(wordStats, func(i, j int) bool {
		if wordStats[i].Count == wordStats[j].Count {
			return wordStats[i].Word < wordStats[j].Word
		}

		return wordStats[i].Count > wordStats[j].Count
	})

	for i, kv := range wordStats {
		top = append(top, kv.Word)
		if (i + 1) == max {
			break
		}
	}

	return top
}
