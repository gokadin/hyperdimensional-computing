package languageGuessing

import (
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"runtime"
	"sync"
)

const GramFactor = 3

func encodeLanguage(text, name string, letters map[string]*hyperdimensional.HdVec) *hyperdimensional.HdVec {
	var encoded *hyperdimensional.HdVec
	in := make(chan []string)
	out := make(chan *hyperdimensional.HdVec)

	for i := 0; i < runtime.NumCPU(); i++ {
        go encodeGram(letters, in, out)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("encoding %s...", name)
		numGrams := len(text) - GramFactor + 1
		count := 0
		for gram := range out {
			if encoded == nil {
				encoded = gram
			} else {
				encoded.Add(gram)
			}
			count++
			if count == numGrams {
				close(out)
			}
		}
	}()

    for textIndex := range text {
    	if textIndex > len(text) - GramFactor {
    		break
		}

		gramLetters := make([]string, GramFactor)
        for index := range gramLetters {
            gramLetters[index] = string(text[textIndex + index])
		}

        in <- gramLetters
	}
	close(in)

    wg.Wait()
	encoded.ToBipolar()
    fmt.Print(" done\n")
	return encoded
}

func encodeGram(letters map[string]*hyperdimensional.HdVec, in chan []string, out chan *hyperdimensional.HdVec) {
	for gramLetters := range in {
		var gram *hyperdimensional.HdVec
		for i, letterName := range gramLetters {
			if i == 0 {
				gram = hyperdimensional.Rotate(letters[letterName], len(gramLetters) - 1)
				continue
			}

			next := letters[letterName]
			if len(gramLetters) - i - 1 != 0 {
				next = hyperdimensional.Rotate(next, len(gramLetters) - i - 1)
			}

			gram = hyperdimensional.Multiply(gram, next)
		}

		out <- gram
	}
}
