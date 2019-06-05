package main

import (
	"github.com/gokadin/hyperdimensional-computing/example/text"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	text.Run()
}
