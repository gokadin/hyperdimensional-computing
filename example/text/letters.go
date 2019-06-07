package text

import "github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"

type Letters [127]*hyperdimensional.VecBinomial

func NewLetters() Letters {
	return *new([127]*hyperdimensional.VecBinomial)
}
