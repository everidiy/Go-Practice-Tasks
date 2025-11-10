package main

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	lower := strings.ToLower(text)
	replacer := strings.NewReplacer(".", " ", ",", " ", "!", " ", "?", " ")
	cleanText := replacer.Replace(lower)

	wordsSlice := strings.Fields(cleanText)
	words := len(wordsSlice)

	wordMap := make(map[string]int)
	for _, word := range wordsSlice {
		wordMap[word]++
	}
	unique := len(wordMap)

	maxWord := ""
	maxCount := 0
	for word, count := range wordMap {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}

	topFive := getTopWords(wordMap, 5)

	fmt.Printf("Количество слов: %d\n", words)
	fmt.Printf("Количество уникальных слов: %d\n", unique)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз) \n", maxWord, maxCount)
	fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
	for _, word := range topFive {
		fmt.Printf("\"%s\": %d раз", word, wordMap[word])
	}

}

func getTopWords(wordMap map[string]int, n int) []string {
	slice := []string{}

	for word := range wordMap {
		slice = append(slice, word)
	}

	sort.Slice(slice, func(i, j int) bool { return wordMap[slice[i]] > wordMap[slice[j]] })

	if n > len(slice) {
		n = len(slice)
	}

	return slice[:n]
}
