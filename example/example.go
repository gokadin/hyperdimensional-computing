package example

import (
	"../hyperdimentional"
	"fmt"
)

type example struct {
	letters [127]*hyperdimentional.HdVec
	trigrams []*hyperdimentional.HdVec
	english *hyperdimentional.HdVec
}

func Run() {
	example := newExample()

	_ = example
}

func newExample() *example {
	example := &example{}

	example.encodeLetters()
	example.encodeTrigrams("the quick brown fox jumped over the green turtle and ate an ant sometime around noon yesterday. This was the saddest moment in history.")
	example.encodeEnglish()

	return example
}

func (e *example) encodeLetters() {
	for i := 0; i < len(e.letters); i++ {
		e.letters[i] = hyperdimentional.NewHdVec()
	}
}

func (e *example) encodeTrigrams(text string) {
    for index := range text {
    	if index >= len(text) - 3 {
            break
		}

		first := hyperdimentional.Rotate(hyperdimentional.Rotate(e.letters[text[index]]))
		second := hyperdimentional.Rotate(e.letters[text[index + 1]])
		third := e.letters[text[index + 2]]
		firstMultiply := hyperdimentional.Multiply(first, second)
		secondMultiply := hyperdimentional.Multiply(firstMultiply, third)

		e.trigrams = append(e.trigrams, secondMultiply)
	}

    fmt.Println(len(e.trigrams))
}

func (e *example) encodeEnglish() {

}
