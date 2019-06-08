package text

import (
	"github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"
	"sync"
)

const GramFactor = 3

type Encoder struct {
	letters *Letters
	profile *hyperdimensional.VecBinomial
	mutex *sync.Mutex
	totalCount int
	counter int
}

func NewEncoder(letters *Letters) *Encoder {
	return &Encoder{
		letters: letters,
		mutex: new(sync.Mutex),
	}
}

func (e *Encoder) encodeLanguage(text *string) *hyperdimensional.VecBinomial {
	e.profile = nil
	e.counter = 0
	gramChannel := make(chan *[]uint8)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
        wg.Add(1)
        go e.encodeGram(gramChannel, wg)
	}
	
	indices := make([]uint8, GramFactor)
	e.totalCount = len(*text) - GramFactor
    for textIndex := range *text {
    	if textIndex > len(*text) - GramFactor {
    		break
		}

        for index := range indices {
            indices[index] = (*text)[textIndex + index]
		}

        gramChannel <- &indices
	}

	close(gramChannel)
	wg.Wait()

	e.profile.ToBinomial()
	return e.profile
}

func (e *Encoder) encodeGram(textIndicesChannel chan *[]uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	for textIndices := range textIndicesChannel {
		var gram *hyperdimensional.VecBinomial
		for i, textIndex := range *textIndices {
			if i == 0 {
				gram = hyperdimensional.Rotate(e.letters[textIndex], len(*textIndices) - 1)
				continue
			}

			next := e.letters[textIndex]
			if len(*textIndices) - i - 1 != 0 {
				next = hyperdimensional.Rotate(next, len(*textIndices) - i - 1)
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
