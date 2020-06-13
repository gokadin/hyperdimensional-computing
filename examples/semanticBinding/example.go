package semanticBinding

import (
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
)

type example struct {

}

func NewExample() *example {
	return &example{}
}

func (e *example) Run() {
	fmt.Println("Encoding phrase (I want to run)... (this might take a while)")

	// pick random vectors
	vectors := make(map[string]*hyperdimensional.HdVec)
	vectors["subject"] = hyperdimensional.NewRandBipolar()
	vectors["verb"] = hyperdimensional.NewRandBipolar()
	vectors["object"] = hyperdimensional.NewRandBipolar()
	vectors["I"] = hyperdimensional.NewRandBipolar()
	vectors["want"] = hyperdimensional.NewRandBipolar()
	vectors["run"] = hyperdimensional.NewRandBipolar()

	// bind: "I want to run"
	p := hyperdimensional.Add(
		hyperdimensional.CircularConvolution(vectors["subject"], vectors["I"]),
		hyperdimensional.CircularConvolution(vectors["verb"], vectors["want"]),
		hyperdimensional.CircularConvolution(vectors["object"], vectors["run"]))

	// unbind: query what is the verb?
	answer := hyperdimensional.CircularConvolution(p, vectors["verb"].Involution())

	// compare angle with all vectors to find the closest match
	bestMatch := hyperdimensional.FindClosestCosineInMap(answer, vectors)

	fmt.Printf("\nThe query \"what is the verb?\" returned: %s\n", bestMatch)
}