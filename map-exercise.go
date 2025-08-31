package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordMap := make(map[string]int)
	for _, v := range words {
		elem, ok := wordMap[v]
		if ok == true {
			wordMap[v] = elem + 1
		} else {
			wordMap[v] = 1
		}
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
