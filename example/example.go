package main

import (
	"fmt"
	"github.com/gokadin/hyperdimentional/src/hyperdimentional"
	"io/ioutil"
)

type example struct {
	letters [127]hyperdimentional.Vector
	trigrams []hyperdimentional.Vector
	english hyperdimentional.Vector
}

func Run() {
	eng := newExample("training_en")
	_ = eng

	latin := newExample("training_latin")
	_ = latin

	comp := newExample("comp")
	_ = comp

	x := hyperdimentional.Cosine(eng.english, comp.english)
	y := hyperdimentional.Cosine(latin.english, comp.english)

	fmt.Println("ENG -> comp: ", x)
	fmt.Println("LATIN -> comp: ", y)
}

func newExample(filename string) *example {
	b, err := ioutil.ReadFile("data/" + filename)
	if err != nil {
		fmt.Println("Could not find file ", filename, err.Error())
		return &example{}
	}

	example := &example{}

	example.encodeLetters()
	example.encodeTrigrams(string(b))
	example.encodeEnglish()

	return example
}

func (e *example) encodeLetters() {
	for i := 0; i < len(e.letters); i++ {
		e.letters[i] = hyperdimentional.New(10000)
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
}

func (e *example) encodeEnglish() {
	for index, trigram := range e.trigrams {
		if index == 0 {
			e.english = trigram
			continue
		}

        e.english = hyperdimentional.Add(e.english, trigram)
	}
}
