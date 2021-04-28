package semanticBinding

import (
	"awesomeProject/hyperdimensional"
	"fmt"
)

type example struct {
}

func NewExample() *example {
	return &example{}
}

func (e *example) Run() {
	fmt.Println("Encoding phrase \"I want to run\"... (this might take a while)")

	// pick random vectors
	vectors := make(map[string]*hyperdimensional.HdVec)
	vectors["subject"] = hyperdimensional.Rand()
	vectors["verb"] = hyperdimensional.Rand()
	vectors["object"] = hyperdimensional.Rand()
	vectors["I"] = hyperdimensional.Rand()
	vectors["want"] = hyperdimensional.Rand()
	vectors["run"] = hyperdimensional.Rand()

	// bind: "I want to run"
	p := hyperdimensional.Add(
		hyperdimensional.Xor(vectors["subject"], vectors["I"]),
		hyperdimensional.Xor(vectors["verb"], vectors["want"]),
		hyperdimensional.Xor(vectors["object"], vectors["run"]))

	// unbind: query what is the verb?
	answer := hyperdimensional.Xor(p, vectors["verb"])

	// compare angle with all vectors to find the closest match
	bestMatch := hyperdimensional.FindClosestCosineInMap(answer, vectors)

	fmt.Printf("\nThe query \"what is the verb?\" returned: %s\n", bestMatch)
}
