package text

import (
	"fmt"
	"github.com/gokadin/hyperdimentional/src/hyperdimentional"
	"io/ioutil"
	"sync"
)

type example struct {
	letters [127]*hyperdimentional.VecBinomial
	lang *hyperdimentional.VecBinomial
	text string
	mutex *sync.Mutex
}

func Run() {
	b, err := ioutil.ReadFile("data/training_en")
	if err != nil {
		fmt.Println("Could not find file training_en", err.Error())
	}
	eng := NewExample(string(b))
	eng.encodeLetters()
	eng.encodeLanguage()
	_ = eng

	latin := NewExample("training_latin")
	latin.encodeLetters()
	latin.encodeLanguage()
	_ = latin

	comp := NewExample("comp")
	comp.encodeLetters()
	comp.encodeLanguage()
	_ = comp

	x := hyperdimentional.Cosine(eng.lang, comp.lang)
	y := hyperdimentional.Cosine(latin.lang, comp.lang)

	fmt.Println("ENG -> comp: ", x)
	fmt.Println("LATIN -> comp: ", y)
}

func NewExample(text string) *example {
	return &example{
		text: text,
		mutex: &sync.Mutex{},
	}
}

func (e *example) encodeLetters() {
	for i := 0; i < len(e.letters); i++ {
		e.letters[i] = hyperdimentional.NewVecBinomial(10000)
	}
}

func (e *example) getLetter(index uint8) *hyperdimentional.VecBinomial {
	if index < 0 || index > uint8(len(e.letters) - 1) {
		return e.letters[0]
	}

	return e.letters[index]
}

func (e *example) encodeLanguage() {
	fmt.Print("Encoding language... ")

	lCh := make(chan int)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go e.worker(lCh, wg)
	}

	for index := range e.text {
		if index >= len(e.text) - 3 {
			break
		}

		lCh <- index
	}

	close(lCh)
	wg.Wait()

	fmt.Println("DONE")
}

func (e *example) worker(lCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for index := range lCh {
		first := hyperdimentional.Rotate(hyperdimentional.Rotate(e.getLetter(e.text[index])))
		second := hyperdimentional.Rotate(e.getLetter(e.text[index + 1]))
		third := e.getLetter(e.text[index + 2])
		firstMultiply := hyperdimentional.Multiply(first, second)
		secondMultiply := hyperdimentional.Multiply(firstMultiply, third)

		e.mutex.Lock()

		if index == 0 {
			e.lang = secondMultiply
		} else {
			e.lang.Add(secondMultiply)
		}

		e.mutex.Unlock()
	}
}
