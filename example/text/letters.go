package text

import "github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"

type Letters [127]*hyperdimensional.VecBinomial

func NewLetters() Letters {
	letters := new([127]*hyperdimensional.VecBinomial)
	for i := range letters {
		letters[i] = hyperdimensional.NewVecBinomial(10000)
	}
	
	return  *letters
}
