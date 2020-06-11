package text

import (
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"sync"
)

const GramFactor = 3

type Encoder struct {
	letters []*hyperdimensional.VecBinomial
	profile *hyperdimensional.VecBinomial
	mutex *sync.Mutex
	totalCount int
	counter int
}

func NewEncoder(letters []*hyperdimensional.VecBinomial) *Encoder {
	return &Encoder{
		letters: letters,
		mutex: new(sync.Mutex),
	}
}

func (e *Encoder) encodeLanguage(text *string) *hyperdimensional.VecBinomial {
	e.profile = nil
	e.counter = 0
	gramChannel := make(chan []uint8)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
        wg.Add(1)
        go e.encodeGram(gramChannel, wg)
	}
	
	e.totalCount = len(*text) - GramFactor
    for textIndex := range *text {
    	if textIndex > len(*text) - GramFactor {
    		break
		}

		asciiLetters := make([]uint8, GramFactor)
        for index := range asciiLetters {
            asciiLetters[index] = (*text)[textIndex + index]
		}

        gramChannel <- asciiLetters
	}

	close(gramChannel)
	wg.Wait()

	e.profile.ToBinomial()
	return e.profile
}

func (e *Encoder) encodeGram(AsciiLettersChannel chan []uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	for asciiLetters := range AsciiLettersChannel {
		var gram *hyperdimensional.VecBinomial
		for i, textIndex := range asciiLetters {
			if i == 0 {
				gram = hyperdimensional.Rotate(e.letters[textIndex], len(asciiLetters) - 1)
				continue
			}

			next := e.letters[textIndex]
			if len(asciiLetters) - i - 1 != 0 {
				next = hyperdimensional.Rotate(next, len(asciiLetters) - i - 1)
			}

			gram = hyperdimensional.Multiply(gram, next)
		}

		e.mutex.Lock()

		if e.profile == nil {
			e.profile = gram
		} else {
			e.profile.Add(gram)
		}
		
		e.counter++
		
		e.mutex.Unlock()
	}
}
