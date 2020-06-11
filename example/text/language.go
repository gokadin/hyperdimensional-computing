package text

import (
    "github.com/gokadin/hyperdimensional-computing/hyperdimensional"
)

type Language struct {
    Name string
    Profile *hyperdimensional.VecBinomial
    letters []*hyperdimensional.VecBinomial
    encoder *Encoder
}

func NewLanguage(name string, letters []*hyperdimensional.VecBinomial) *Language {
    return &Language{
        Name: name,
        letters: letters,
        encoder: NewEncoder(letters),
    }
}

func (l *Language) encodeLanguage(text *string) {
    l.Profile = l.encoder.encodeLanguage(text)
}
