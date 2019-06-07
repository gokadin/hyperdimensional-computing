package text

import "github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"

type Language struct {
    Name string
    Profile *hyperdimensional.VecBinomial
}

func NewLanguage(name string) *Language {
    return &Language{
        Name: name,
    }
}
