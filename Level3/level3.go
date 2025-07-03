package main

import (
	"fmt"
	"sort"
	"strings"
)

func splitText(text string) []string {
	var textSlice []string
	var str []string

	for _, char := range text {
		if char == ' ' || char == '.' || char == ',' || char == '?' || char == '!' {
			if len(str) > 0 {
				textSlice = append(textSlice, strings.Join(str, ""))
				str = str[:0]
			}
			continue
		}
		str = append(str, string(char))
	}

	if len(str) > 0 {
		textSlice = append(textSlice, strings.Join(str, ""))
	}

	return textSlice
}

func AnalyzeText(text string) {

	textSlice := splitText(text)
	countWords := len(textSlice)
	textMap := make(map[string]int)

	for _, word := range textSlice {
		word = strings.ToLower(word)
		textMap[word]++
	}

	countUniqueWords := len(textMap)
	topN := getTopWords(textMap, 5)

	fmt.Println(
		fmt.Sprintf("Количество слов: %d", countWords),
	)

	fmt.Println(
		fmt.Sprintf("Количество уникальных слов: %d", countUniqueWords),
	)
	fmt.Println(
		fmt.Sprintf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)",
			topN[0],
			textMap[topN[0]],
		))
	fmt.Println("Топ-5 самых часто встречающихся слов:")

	for i := 0; i < len(topN); i++ {
		fmt.Println(
			fmt.Sprintf("\"%s\": %d раз",
				topN[i],
				textMap[topN[i]],
			),
		)
	}
}

func getTopWords(wordMap map[string]int, n int) []string {

	maxTimes := make(map[int]string)

	for key, value := range wordMap {
		if _, ok := maxTimes[value]; ok {
			continue
		}
		maxTimes[value] = key
	}

	var topWords []string
	var topWordsInt []int

	for key, _ := range maxTimes {
		topWordsInt = append(topWordsInt, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(topWordsInt)))

	for i := 0; i < len(topWordsInt); i++ {
		topWords = append(topWords, maxTimes[topWordsInt[i]])
	}

	return topWords[:n]
}

func main() {
	//AnalyzeText("one two two three three three four four four four five five five five five")
	//AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	//AnalyzeText("Я так люблю море. Я на море. Я так люблю плавать. Море! Я море!!! ЛЮБЛЮ МОРЕ")
	fmt.Println(getTopWords(map[string]int{
		"one":   5,
		"two":   4,
		"three": 3,
		"four":  2,
		"five":  1,
	}, 1))
	fmt.Println(getTopWords(map[string]int{
		"one":   5,
		"two":   4,
		"three": 3,
		"four":  2,
		"five":  1,
	}, 2))
	fmt.Println(getTopWords(map[string]int{
		"one":   5,
		"two":   4,
		"three": 3,
		"four":  2,
		"five":  1,
	}, 3))
	fmt.Println(getTopWords(map[string]int{
		"one":   5,
		"two":   4,
		"three": 3,
		"four":  2,
		"five":  1,
	}, 4))
	fmt.Println(getTopWords(map[string]int{
		"one":   5,
		"two":   4,
		"three": 3,
		"four":  2,
		"five":  1,
	}, 5))
	fmt.Println(getTopWords(map[string]int{
		"one":   6,
		"two":   5,
		"three": 4,
		"four":  3,
		"five":  2,
		"six":   1,
		"seven": 1,
	}, 5))

}
