package hyperdimensional

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Encoder struct {
	randomVec *HdVec // used for balancing the ADD operation
	letters   map[int]*HdVec
}

func NewEncoder(letters map[int]*HdVec) *Encoder {
	return &Encoder{
		randomVec: Rand(),
		letters:   letters,
	}
}

func (e *Encoder) EncodeVec(text string, gramFactor int) *HdVec {
	if gramFactor <= 0 {
		panic(fmt.Sprintf("cannot encode with gram factor less than or equal to zero, received %d", gramFactor))
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	text = reg.ReplaceAllString(text, "")
	text = strings.ToLower(text)

	if len(text) == 0 {
		return Ones()
	}

	if len(text) < gramFactor {
		gramFactor = len(text)
	}

	numberOfGrams := len(text) - gramFactor + 1
	if numberOfGrams <= 0 {
		log.Printf("nothing to encode for text %s\n", text)
		return Ones()
	}

	encodedGrams := make([]*HdVec, numberOfGrams)
	for textIndex := range text {
		if textIndex > len(text)-gramFactor {
			break
		}

		gramLettersAscii := make([]int, gramFactor)
		for index := range gramLettersAscii {
			gramLettersAscii[index] = int(text[textIndex+index])
		}

		encodedGrams[textIndex] = encodeGram(e.letters, gramLettersAscii)
	}

	if len(encodedGrams)%2 == 0 {
		encodedGrams = append(encodedGrams, e.randomVec)
	}
	return Add(encodedGrams...)
}

func encodeGram(letters map[int]*HdVec, gramLetters []int) *HdVec {
	var gram *HdVec
	for i, letterAscii := range gramLetters {
		if i == 0 {
			gram = Rotate(letters[letterAscii], len(gramLetters)-1)
			continue
		}

		next := letters[letterAscii]
		if len(gramLetters)-i-1 != 0 {
			next = Rotate(next, len(gramLetters)-i-1)
		}

		gram.Multiply(next)
	}

	return gram
}

//func encodeGramAsync(letters map[string]*HdVec, in chan []string, out chan *HdVec) {
//	for gramLetters := range in {
//		var gram *HdVec
//		for i, letterName := range gramLetters {
//			if i == 0 {
//				gram = Rotate(letters[letterName], len(gramLetters)-1)
//				continue
//			}
//
//			next := letters[letterName]
//			if len(gramLetters)-i-1 != 0 {
//				next = Rotate(next, len(gramLetters)-i-1)
//			}
//
//			gram.Multiply(next)
//		}
//
//		out <- gram
//	}
//}
