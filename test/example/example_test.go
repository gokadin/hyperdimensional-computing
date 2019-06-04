package example

import (
	"github.com/gokadin/hyperdimentional/example/text"
	"testing"
)

func Test_NewExample_works(t *testing.T) {
	// Act
	ex := text.NewExample("abc")

	// Assert
	if (*ex.GetText()) != "abc" {
        t.Fatalf("Did not initialize correctly.")
	}
}

func Test_EncodeLanguageWithOneTrigram(t *testing.T) {
	// Arrange
	ex := text.NewExample("abc")
	ex.EncodeLetters()

	// Act
	ex.EncodeLanguage()

	// Assert
	if ex.GetLanguage().Size() != 10000  {
		t.Fatalf("Language was not encoded successfully.")
	}
}

func Test_EncodeLanguage_withMultipleTrigrams(t *testing.T) {
	// Arrange
	ex := text.NewExample("A sample English phrase.")
	ex.EncodeLetters()

	// Act
	ex.EncodeLanguage()

	// Assert
	if ex.GetLanguage().Size() != 10000  {
		t.Fatalf("Language was not encoded successfully.")
	}
}
