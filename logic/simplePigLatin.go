package logic

import (
	"fmt"
	"strings"
)

//Description:
//Move the first letter of each word to the end of it, then add 'ay' to the end of the word.
//pig_it('Pig latin is cool') # igPay atinlay siay oolcay
func SimplePigLatin() {

	var inputsample = []string{
		"Pig latin is cool",
		"This is my string",
		"hey this is my string !",
	}

	var resultsample = []string{
		"igPay atinlay siay oolcay",
		"hisTay siay ymay tringsay",
		"eyhay histay siay ymay tringsay !",
	}

	currectcnt := 0
	for _, input := range inputsample {
		output := PigLatinProcess(input)
		for _, result := range resultsample {
			if output == result {
				currectcnt++
				break
			}
		}
	}
	if currectcnt == len(resultsample) {
		fmt.Println("good job!")
	}
}

func PigLatinProcess(text string) string {
	result := ""
	words := strings.Split(text, " ")
	for _, word := range words {
		if word == "?" || word == "!" {
			result += word + " "
		} else {
			result += word[1:] + word[:1] + "ay "
		}
	}
	return result[:len(result)-1]
}
