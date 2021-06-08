package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestFilterLowercase(t *testing.T) {
	var (
		in  = []string{"Cat", "DOG", "fisH"}
		out = []string{"cat", "dog", "fish"}
	)
	assert.Equal(t, out, filterLowercase(in))
}

func TestFilterStopWord(t *testing.T) {
	var (
		in  = []string{"i", "am", "the", "cat"}
		out = []string{"am", "cat"}
	)
	assert.Equal(t, out, filterStopWord(in))
}

func TestFilterStemmer(t *testing.T) {
	var (
		in  = []string{"cat", "cats", "fish", "fishing", "fished", "airline"}
		out = []string{"cat", "cat", "fish", "fish", "fish", "airlin"}
	)

	assert.Equal(t, out, filterStemmer(in))
}
