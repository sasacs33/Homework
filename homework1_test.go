package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SelectLargerScore(t *testing.T) {
	Input := []Score{
		Score{1, "Tom", "One", 61},
		Score{2, "Sebastianae", "One", 55},
		Score{3, "Woof", "Two", 58},
		Score{4, "Teddy", "Two", 75},
		Score{5, "Maximilians", "Three", 88},
	}
	getLargerScore := SelectLargerScore(Input)
	assert.NotNil(t, getLargerScore)
	assert.Equal(t, 3, len(getLargerScore))
	for _, key := range getLargerScore {
		if key.Score <= 59 || key.Score > 100 || key.Score < 0 {
			t.Error()
		}
	}
}

func Test_SelectSpecialClass(t *testing.T) {
	Input := []Score{
		Score{1, "Tom", "One", 61},
		Score{2, "Sebastianae", "One", 55},
		Score{3, "Woof", "Two", 58},
		Score{4, "Teddy", "Two", 75},
		Score{5, "Maximilians", "Three", 88},
	}
	getSpecialClass := SelectSpecialClass(Input)
	assert.NotNil(t, getSpecialClass)
	assert.Equal(t, 2, len(getSpecialClass))
	for _, key := range getSpecialClass {
		if key.Class != givenclass {
			t.Error()
		}
	}
}

func Test_SelectSpecialName(t *testing.T) {
	Input := []Score{
		Score{1, "Tom", "One", 61},
		Score{2, "Sebastianae", "One", 55},
		Score{3, "Woof", "Two", 58},
		Score{4, "Teddy", "Two", 75},
		Score{5, "Maximilians", "Three", 88},
	}
	getSpecialName := SelectSpecialName(Input)
	assert.NotNil(t, getSpecialName)
	assert.Equal(t, 2, len(getSpecialName))
	for _, key := range getSpecialName {
		if len(key.Name) <= 10 {
			t.Error()
		}
	}
}
