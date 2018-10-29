package main

import (
	"fmt"
)

var keyAnimails = []string{"fly", "spider", "bird", "cat", "dog", "cow"}
var secondLines = []string{
	"That wriggled and wiggled and tickled inside her.",
	"How absurd to swallow a bird.",
	"Fancy that to swallow a cat!",
	"What a hog, to swallow a dog!",
	"I don't know how she swallowed a cow!",
}

func main() {
	getAllLines(keyAnimails, secondLines)
}

func getAllLines(animails []string, lines []string) {
	for i := 0; i < len(animails); i++ {
		getOpeningLine(animails[i], i)
		if i != 0 {
			fmt.Println(lines[i-1])
			for j := i; j >= 1; j-- {
				getLoopingLine(animails[j], animails[j-1], j)
			}
		}
		getEndingLine()
	}
	getTheEnd()
}

func getOpeningLine(animail string, index int) {
	if index == 0 {
		fmt.Println("There was an old lady who swallowed a " + animail + ".")
	} else {
		fmt.Println("There was an old lady who swallowed a " + animail + ";")
	}
}

func getLoopingLine(animail1 string, animail2 string, index int) {
	if index == 1 {
		fmt.Printf("She swallowed the " + animail1 + " to catch the " + animail2 + ";\n")
	} else {
		fmt.Printf("She swallowed the " + animail1 + " to catch the " + animail2 + ",\n")
	}
}

func getEndingLine() {
	fmt.Println("I don't know why she swallowed a fly - perhaps she'll die!")
}

func getTheEnd() {
	fmt.Println("There was an old lady who swallowed a horse...")
	fmt.Println("...She's dead, of course!")
}
