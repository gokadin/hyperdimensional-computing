package main

import (
	"github.com/gokadin/hyperdimensional-computing/example/text"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	
	example := text.NewExample()
	example.Run()
}
