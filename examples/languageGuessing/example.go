package languageGuessing

import (
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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
	letters   map[string]*hyperdimensional.HdVec
	languages map[string]*hyperdimensional.HdVec
	test      *hyperdimensional.HdVec
}

func NewExample() *example {
	if !fileExists(storageDir) {
		_ = os.Mkdir(storageDir, 0755)
	}
	if !fileExists(fmt.Sprintf("%s/%s", storageDir, lettersDir)) {
		_ = os.Mkdir(fmt.Sprintf("%s/%s", storageDir, lettersDir), 0755)
	}

	letters := make(map[string]*hyperdimensional.HdVec, numLetters)
	return &example{
		letters: letters,
	}
}

func (e *example) Run() {
	t := time.Now()

	fmt.Println("Encoding languages... (this might take a while)")

	e.encodeLetters()
	e.encodeLanguages()
	e.encodeTest()

	diff := time.Now().Sub(t)
	fmt.Println("Finished encoding in ", diff.Seconds(), " seconds.")

	bestMatch := hyperdimensional.FindClosestCosineInMap(e.test, e.languages)
	if bestMatch == "" {
		fmt.Println("could not find any match")
	} else {
		fmt.Println("language is " + bestMatch)
	}

	fmt.Println()
	fmt.Println("Asking (what letter is most commonly used after the set: \"th\"?")

	th := hyperdimensional.Multiply(hyperdimensional.Rotate(e.letters["t"], 2), hyperdimensional.Rotate(e.letters["h"], 1))

	answerVec := hyperdimensional.Multiply(th, e.languages["en"])
	answer := hyperdimensional.FindClosestCosineInMap(answerVec, e.letters)

	fmt.Println(fmt.Sprintf("Answer is %s", answer))
}

func (e *example) encodeLetters() {
	for i := 0; i < 128; i++ {
		fileSuffix := strconv.Itoa(i)
		name := string(i)
		letterFilePath := fmt.Sprintf("%s/%s/%s%s", storageDir, lettersDir, letterFilePrefix, fileSuffix)
		if UseCache {
			if fileExists(letterFilePath) {
				e.letters[name] = VecBipolarFromFile(letterFilePath)
				continue
			}

			e.letters[name] = hyperdimensional.NewRandBipolar()
			writeToCache(letterFilePath, e.letters[name])
			continue
		}

		e.letters[name] = hyperdimensional.NewRandBipolar()
	}
}

func (e *example) encodeLanguages() {
	e.languages = make(map[string]*hyperdimensional.HdVec)
	err := filepath.Walk("data/training/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		languageFilePath := fmt.Sprintf("%s/%s%s", storageDir, languageFilePrefix, info.Name())

		if UseCache {
			if fileExists(languageFilePath) {
				e.languages[info.Name()] = VecBipolarFromFile(languageFilePath)
				return nil
			}
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		text := string(b)

		languageVec := encodeLanguage(text, info.Name(), e.letters)
		e.languages[info.Name()] = languageVec
		if UseCache {
			writeToCache(languageFilePath, languageVec)
		}

		return  nil
	})
	if err != nil {
		panic(err)
	}
}

func (e *example) encodeTest() {
	b, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	text := string(b)

	e.test = encodeLanguage(text, "test", e.letters)
}

func (e *example) compare() {
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
