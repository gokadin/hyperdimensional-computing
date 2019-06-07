package text

import (
	"github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"
)

const GramFactor = 3

func encodeLanguage(letters *Letters, text *string) *hyperdimensional.VecBinomial {
	var profile *hyperdimensional.VecBinomial
	indices := make([]uint8, GramFactor)
	for i := 0; i < len(*text) - GramFactor; i+= GramFactor {
        for index := range indices {
            indices[index] = (*text)[i + index]
		}

		if profile == nil {
            profile = encodeGram(&indices, letters)
		} else {
            profile.Add(encodeGram(&indices, letters))
		}
	}

	profile.ToBinomial()
	return profile
}

func encodeGram(textIndices *[]uint8, letters *Letters) *hyperdimensional.VecBinomial {
	var gram *hyperdimensional.VecBinomial
    for i, textIndex := range *textIndices {
    	if i == 0 {
			gram = hyperdimensional.Rotate(letters[(*textIndices)[0]], len(*textIndices) - 1)
    		continue
		}

    	var next *hyperdimensional.VecBinomial
    	if len(*textIndices) - i - 1 == 0 {
			next = hyperdimensional.Rotate(letters[textIndex], len(*textIndices) - i - 1)
		} else {
            next = letters[textIndex]
		}
		gram = hyperdimensional.Multiply(gram, next)
	}
	
	return  gram
}
