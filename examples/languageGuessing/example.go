package languageGuessing

import (
	"awesomeProject/hyperdimensional"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	UseCache           = true
	numLetters         = 127
	storageDir         = "storage"
	lettersDir         = "letters"
	letterFilePrefix   = "computed_letter_"
	languageFilePrefix = "computed_"
	testFilePath       = "data/testing/test_file"
	gramFactor         = 3
)

type example struct {
	letters   map[int]*hyperdimensional.HdVec
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

	letters := make(map[int]*hyperdimensional.HdVec, numLetters)
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

	th := hyperdimensional.Rotate(e.letters[116], 2).Xor(hyperdimensional.Rotate(e.letters[104], 1))

	answerVec := th.Xor(e.languages["en"])
	answer := strconv.Itoa(hyperdimensional.FindClosestCosineInMapInt(answerVec, e.letters))

	fmt.Println(fmt.Sprintf("Answer is %s", answer))
}

func (e *example) encodeLetters() {
	for i := 0; i < 128; i++ {
		fileSuffix := strconv.Itoa(i)
		letterFilePath := fmt.Sprintf("%s/%s/%s%s", storageDir, lettersDir, letterFilePrefix, fileSuffix)
		if UseCache {
			if fileExists(letterFilePath) {
				e.letters[i] = VecFromFile(letterFilePath)
				continue
			}

			e.letters[i] = hyperdimensional.Rand()
			writeToCache(letterFilePath, e.letters[i])
			continue
		}

		e.letters[i] = hyperdimensional.Rand()
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
				e.languages[info.Name()] = VecFromFile(languageFilePath)
				return nil
			}
		}

		b, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}
		text := string(b)

		encoder := hyperdimensional.NewEncoder(e.letters)
		languageVec := encoder.EncodeVec(text, gramFactor)
		e.languages[info.Name()] = languageVec
		if UseCache {
			writeToCache(languageFilePath, languageVec)
		}

		return nil
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

	encoder := hyperdimensional.NewEncoder(e.letters)
	e.test = encoder.EncodeVec(text, gramFactor)
}

func (e *example) compare() {
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
