package main

import (
	"fmt"
	"github.com/gokadin/hyperdimentional/src/hyperdimentional"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	v1 := hyperdimentional.NewVecBinomial(10000)
	v2 := hyperdimentional.NewVecBinomial(10000)
	v3 := hyperdimentional.NewVecBinomial(10000)
	v4 := hyperdimentional.NewVecBinomial(10000)

	// Addition => similar
	// Multiplication => dissimilar

	v1.Add(v2)
	v1.Add(v3)
	v1.Add(v4)

	v1.ToBinomial()

	fmt.Println("v1 <-> v2 ", hyperdimentional.Cosine(v1, v2))
	fmt.Println("v1 <-> v3 ", hyperdimentional.Cosine(v1, v3))
	fmt.Println("v1 <-> v4 ", hyperdimentional.Cosine(v1, v4))
}
