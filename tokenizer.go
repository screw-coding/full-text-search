package main

import (
	"strings"
	"unicode"
)

//词法分析

//分词
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})
}

//分词并处理词
func analyze(text string) []string {
	//分词
	tokens := tokenize(text)
	//转小写
	tokens = filterLowercase(tokens)
	//排除语气词
	tokens = filterStopWord(tokens)
	//提取词干
	tokens = filterStemmer(tokens)
	//返回处理好的词
	return tokens
}
