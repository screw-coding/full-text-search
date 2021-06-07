package main

import "strings"
import snowballeng "github.com/kljensen/snowball/english"

//负责过滤

//全部变小写
func filterLowercase(tokens []string) []string {
	result := make([]string, len(tokens))
	for i, token := range tokens {
		result[i] = strings.ToLower(token)
	}
	return result
}

//去除通用词,语气词
func filterStopWord(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {}, "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	result := make([]string, 0, len(tokens))
	for _, token := range tokens {
		_, ok := stopWords[token]
		if !ok {
			result = append(result, token)
		}
	}
	return result
}

//词干提取 因为英语中语法的不同，文档中可能包含同一个词的不同表现形式，而词干提取就是为了将这些词转化为词干。比如fishing、fished 、 fisher 都会转化为基础词干fish
func filterStemmer(tokens []string) []string {
	result := make([]string, len(tokens))
	for i, token := range tokens {
		result[i] = snowballeng.Stem(token, false)
	}
	return result
}
