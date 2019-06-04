package main

import (
	"github.com/gokadin/hyperdimentional/example/text"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	text.Run()
}
