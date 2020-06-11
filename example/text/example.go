package text

import (
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const (
	UseCache = true
	numLetters = 127
	storageDir = "storage"
	lettersDir = "letters"
	letterFilePrefix = "computed_letter_"
	languageFilePrefix = "computed_"
	testFilePath = "data/testing/test_file"
)

type example struct {
	letters []*hyperdimensional.VecBinomial
	languages []*Language
	test *Language
}

func NewExample() *example {
	if !fileExists(storageDir) {
		_ = os.Mkdir(storageDir, 0755)
	}
	if !fileExists(fmt.Sprintf("%s/%s", storageDir, lettersDir)) {
		_ = os.Mkdir(fmt.Sprintf("%s/%s", storageDir, lettersDir), 0755)
	}

	return &example{
		letters: make([]*hyperdimensional.VecBinomial, numLetters),
	}
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
	for i := range e.letters {
		letterFilePath := fmt.Sprintf("%s/%s/%s%s", storageDir, lettersDir, letterFilePrefix, strconv.Itoa(i))
		if UseCache {
			if fileExists(letterFilePath) {
				e.letters[i] = VecBinomialFromFile(letterFilePath)
				continue
			}

			e.letters[i] = hyperdimensional.NewRandBinomial()
			writeToCache(letterFilePath, e.letters[i])
			continue
		}

		e.letters[i] = hyperdimensional.NewRandBinomial()
	}
}

func (e *example) encodeLanguages() {
	e.languages = make([]*Language, 0)

	wg := new(sync.WaitGroup)

	err := filepath.Walk("data/training/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		language := NewLanguage(info.Name(), e.letters)
		languageFilePath := fmt.Sprintf("%s/%s%s", storageDir, languageFilePrefix, language.Name)

		if UseCache {
			if fileExists(languageFilePath) {
				language.Profile = VecBinomialFromFile(languageFilePath)
				e.languages = append(e.languages, language)
				return nil
			}
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			b, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}
			text := string(b)

			language.encodeLanguage(&text)
			if UseCache {
				writeToCache(languageFilePath, language.Profile)
			}
		}()

		e.languages = append(e.languages, language)
		return  nil
	})
	if err != nil {
		panic(err)
	}

	wg.Wait()
}

func (e *example) encodeTest() {
	b, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	text := string(b)

	e.test = NewLanguage("test", e.letters)
	e.test.encodeLanguage(&text)
}

func (e *example) compare() {
	var smallestAngle float32 = -1
	var bestMatch *Language
    for _, language := range e.languages {
    	angle := hyperdimensional.Cosine(e.test.Profile, language.Profile)
		fmt.Println(language.Name + ": ", angle)

        if angle > smallestAngle {
        	smallestAngle = angle
        	bestMatch = language
		}
	}

    if bestMatch == nil {
    	panic("could not find any match.")
	}

    fmt.Println("language is " + bestMatch.Name)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
