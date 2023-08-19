package main

import (
	"fmt"
	"sort"
	"strings"
	
	// "github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
)

var topics = map[int]string{
	1:  "Personal Pronouns",
	2:  "Demonstrative Pronouns",
	3:  "Question Words",
	4:  "Present Verbs",
	5:  "Adjectives",
	6:  "Prepositions",
	7:  "Possessive Determiners",
	8:  "Adverbs",
	9:  "Colors",
	10: "Dates",
	11: "Conjunctions",
	12: "Part of the House",
	0:  "All",
}

func main() {
	fmt.Println("What topic you want to train?")
	showTopics()
	words := chooseTopic()
	point := float64(0)

	processGame(words, point)
	getResult(point, words)
}

func showTopics() {
	keys := make([]int, 0)
	for k, _ := range topics {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, "] ", topics[k])
	}
}

func chooseTopic() map[string]string {
	var topicChosen int
	fmt.Scan(&topicChosen)
	words := map[string]string{}
	
	switch topicChosen {
	case 1:
		words = getPersonalPronouns()
	case 2:
		words = getDemonstrativePronouns()
	case 3:
		words = getQuestionWords()
	case 4:
		words = getPresentVerbs()
	case 5:
		words = getAdjectives()
	case 6:
		words = getPreposition()
	case 7:
		words = getPossessiveDeterminers()
	case 8:
		words = getAdverb()
	case 9:
		words = getColors()
	case 10:
		words = getDates()
	case 11:
		words = getConjunctions()
	case 12:
		words = getPartOfHouse()
	case 0:
		words = getPersonalPronouns()
		maps.Copy(words, getDemonstrativePronouns())
		maps.Copy(words, getQuestionWords())
		maps.Copy(words, getPresentVerbs())
		maps.Copy(words, getAdjectives())
		maps.Copy(words, getPreposition())
		maps.Copy(words, getPossessiveDeterminers())
		maps.Copy(words, getAdverb())
		maps.Copy(words, getColors())
		maps.Copy(words, getDates())
		maps.Copy(words, getPartOfHouse())
	}

	return words
}

func processGame(words map[string]string, point float64) {
	for ukr, eng := range words {
		var userInput string
		fmt.Printf("\n Guess the meaning of this word: %v \n", strings.ToUpper(ukr))
		fmt.Scan(&userInput)

		if strings.ToLower(userInput) == strings.ToLower(eng) {
			point++
			fmt.Println("Great Job!")
			fmt.Printf("Current point %v \n", point)
		} else {
			fmt.Printf("Error, the real meaning is : %v \n", strings.ToUpper(eng))
			fmt.Printf("Current point %v \n", point)
		}
	}
}

func getResult(point float64, words map[string]string) {
	fmt.Println("\nGame over!")
	fmt.Printf("You made %v point on %v elements \n", point, len(words))

	var result = point / float64(len(words)) * 100
	fmt.Printf("You was right %v%% of the time \n", int(result))
}
