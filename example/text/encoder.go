package text

import "github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"

func encodeGram(textIndices *[]int, letters *Letters) *hyperdimensional.VecBinomial {
	gram := hyperdimensional.Rotate(letters[(*textIndices)[0]], len(*textIndices) - 1)
    for i, textIndex := range *textIndices {
		next := hyperdimensional.Rotate(letters[textIndex], len(*textIndices) - i - 2)
		gram = hyperdimensional.Multiply(gram, next)
	}
	
	return  gram
}
