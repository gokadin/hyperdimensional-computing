package text

import (
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const UseCache = true

type example struct {
	letters Letters
	languages []*Language
	test *Language
}

func NewExample() *example {
	return &example{}
}

func (e *example) Run() {
	t := time.Now()

	e.encodeLetters()
	e.encodeLanguages()
	e.encodeTest()

	diff := time.Now().Sub(t)
	fmt.Println("Finished encoding in ", diff.Seconds(), " seconds.")

	e.compare()
}

func (e *example) encodeLetters() {
	e.letters = NewLetters()

	if UseCache {
		for i := range e.letters {
			e.letters[i] = VecBinomialFromFile("storage/letters/computed_letter_" + strconv.Itoa(i))
		}
	} else {
		for i := range e.letters {
			e.letters[i] = hyperdimensional.NewVecBinomial(10000)
		}
	}
}

func (e *example) encodeLanguages() {
	e.languages = make([]*Language, 0)

	wg := new(sync.WaitGroup)

	err := filepath.Walk("data/training/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		language := NewLanguage(info.Name())

		if UseCache {
			language.Profile = VecBinomialFromFile("storage/computed_" + info.Name() + ".ptrn")
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()

				b, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}
				text := string(b)

				language.Profile = encodeLanguage(&e.letters, &text)
			}()
		}

		e.languages = append(e.languages, language)
		return  nil
	})
	if err != nil {
		panic(err)
	}

	wg.Wait()
}

func (e *example) encodeTest() {
	b, err := ioutil.ReadFile("data/testing/test1")
	if err != nil {
		panic(err)
	}
	text := string(b)

	e.test = NewLanguage("test")
    e.test.Profile = encodeLanguage(&e.letters, &text)
}

func (e *example) compare() {
	smallestAngle := -1.0
	var bestMatch *Language
    for _, language := range e.languages {
    	angle := hyperdimensional.Cosine(e.test.Profile, language.Profile)
        if angle > smallestAngle {
        	smallestAngle = angle
        	bestMatch = language
		}
	}

    if bestMatch == nil {
    	panic("Could not find any match.")
	}

    fmt.Println("Language is " + bestMatch.Name)
}
