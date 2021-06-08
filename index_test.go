package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndex(t *testing.T) {
	index := make(index)
	assert.Nil(t, index.search("shoe"))
	assert.Nil(t, index.search("xxx"))
	index.add([]document{
		{ID: 1, Text: "A shoe on floor,not any shoes"},
	})
	assert.Nil(t, index.search("a"))
	assert.Equal(t, index.search("shoe"), []int{1})
	assert.Equal(t, index.search("Shoe"), []int{1})
	assert.Equal(t, index.search("floor"), []int{1})
	index.add([]document{
		{ID: 2, Text: "Shoe is a shoe"},
	})
	assert.Nil(t, index.search("a"))
	assert.Equal(t, index.search("shoe"), []int{1, 2})
	assert.Equal(t, index.search("Shoe"), []int{1, 2})
	assert.Equal(t, index.search("floor"), []int{1})

}
