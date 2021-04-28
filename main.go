package main

import (
	"awesomeProject/examples/languageGuessing"
	"awesomeProject/examples/semanticBinding"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())

	//runLanguageGuessingExample()
	runSemanticBindingExample()
}

func runLanguageGuessingExample() {
	example := languageGuessing.NewExample()
	example.Run()
}

func runSemanticBindingExample() {
	example := semanticBinding.NewExample()
	example.Run()
}
