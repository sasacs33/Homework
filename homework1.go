package main

import (
	"fmt"
)

type Score struct {
	Id    int
	Name  string
	Class string
	Score int
}

const (
	givenscore   = 59
	givenclass   = "One"
	givennamelen = 10
)

func main() {
	Input := []Score{
		Score{1, "Tom", "One", 61},
		Score{2, "Sebastianae", "One", 55},
		Score{3, "Woof", "Two", 58},
		Score{4, "Teddy", "Two", 75},
		Score{5, "Maximilians", "Three", 88},
	}
	getLargerScore := SelectLargerScore(Input)
	fmt.Println(getLargerScore)
	getSpecialClass := SelectSpecialClass(Input)
	fmt.Println(getSpecialClass)
	getSpecialName := SelectSpecialName(Input)
	fmt.Println(getSpecialName)
}

func Select(source []Score, key func(s Score) bool) []Score {
	var results []Score
	for _, s := range source {
		if key(s) {
			results = append(results, s)
		}
	}
	return results
}

func SelectLargerScore(source []Score) []Score {
	return Select(source, func(s Score) bool {
		if s.Score > givenscore {
			return true
		}
		return false
	})
}

func SelectSpecialClass(source []Score) []Score {
	return Select(source, func(s Score) bool {
		if s.Class == givenclass {
			return true
		}
		return false
	})
}

func SelectSpecialName(source []Score) []Score {
	return Select(source, func(s Score) bool {
		if len(s.Name) > givennamelen {
			return true
		}
		return false
	})
}
