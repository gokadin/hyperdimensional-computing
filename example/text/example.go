package text

import (
	"bufio"
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type letters struct {
	letters [127]*hyperdimensional.VecBinomial
}

type example struct {
	letters *letters
	lang *hyperdimensional.VecBinomial
	text string
	mutex *sync.Mutex
}

func Run2() {
	letters := NewLetters()
	_ = letters
	
	err := filepath.Walk("data/training", func(path string, info os.FileInfo, err error) error {
		
		
		return  nil
	})
	if err != nil {
        panic(err)
	}
}

func Run() {
	fromCache := true

	letters := &letters{}
	letters.EncodeLetters(fromCache)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	t := time.Now()

	var eng *example
	go func() {
		eng = runLanguage(letters, "en", wg, fromCache)
	}()

	wg.Add(1)
	var fr *example
	go func() {
		fr = runLanguage(letters, "fr", wg, fromCache)
	}()

	wg.Add(1)
	var test *example
	go func() {
		test = runLanguage(letters, "test", wg, false)
	}()

	wg.Wait()

	diff := time.Now().Sub(t)
	fmt.Println("Finished encoding in ", diff.Seconds(), " seconds.")

	x := hyperdimensional.Cosine(eng.lang, test.lang)
	y := hyperdimensional.Cosine(fr.lang, test.lang)

	fmt.Println("ENG -> ", x)
	fmt.Println("FR -> ", y)

	if x > y {
		fmt.Println("LANGUAGE IS ENGLISH")
	} else {
		fmt.Println("LANGUAGE IS FRENCH")
	}
}

func runLanguage(letters *letters, lang string, wg *sync.WaitGroup, useCache bool) *example {
	defer wg.Done()

	if useCache {
		return runLanguageFromCache(letters, lang)
	}

	b, err := ioutil.ReadFile("data/training_" + lang)
	if err != nil {
		fmt.Println("Could not find file training_" + lang, err.Error())
	}
	ex := NewExample(string(b), letters)
	ex.EncodeLanguage()

	writePattern("storage/computed_" + lang + ".ptrn", ex.lang)

	return ex
}

func runLanguageFromCache(letters *letters, lang string) *example {
	ex := NewExample("", letters)
	ex.lang = vecFromFile("storage/computed_" + lang + ".ptrn")
	return ex
}

func vecFromFile(filename string) *hyperdimensional.VecBinomial {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Could not find file " + filename, err.Error())
	}
	defer file.Close()

	vec := hyperdimensional.NewVecBinomial(10000)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		parsed, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			panic(err)
		}
		vec.Set(i, parsed)
		i++
	}

	return vec
}

func writePattern(filename string, vec *hyperdimensional.VecBinomial) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	for _, value := range *vec.Values() {
		w.WriteString(strconv.FormatFloat(value, 'f', 0, 64) + "\n")
	}

	w.Flush()
}

func NewExample(text string, letters *letters) *example {
	return &example{
		letters: letters,
		text: text,
		mutex: &sync.Mutex{},
	}
}

func (e *example) GetText() *string {
	return &e.text
}

func (l *letters) EncodeLetters(fromCache bool) {
	if fromCache {
		for i := 0; i < len(l.letters); i++ {
			l.letters[i] = vecFromFile("storage/letters/computed_letter_" + strconv.Itoa(i))
		}

		return
	}

	for i := 0; i < len(l.letters); i++ {
		l.letters[i] = hyperdimensional.NewVecBinomial(10000)

		writePattern("storage/letters/computed_letter_" + strconv.Itoa(i), l.letters[i])
	}
}

func (l *letters) GetLetter(index uint8) *hyperdimensional.VecBinomial {
	if index < 0 || index > uint8(len(l.letters) - 1) {
		return l.letters[0]
	}

	return l.letters[index]
}

func (e *example) GetLanguage() *hyperdimensional.VecBinomial {
	return e.lang
}

func (e *example) EncodeLanguage() {
	lCh := make(chan int)
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go e.worker(lCh, wg)
	}

	for index := range e.text {
		if index > len(e.text) - 3 {
			break
		}

		lCh <- index
	}

	close(lCh)
	wg.Wait()

	e.lang.ToBinomial()
}

func (e *example) worker(lCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for index := range lCh {
		first := hyperdimensional.Rotate(hyperdimensional.Rotate(e.letters.GetLetter(e.text[index])))
		second := hyperdimensional.Rotate(e.letters.GetLetter(e.text[index + 1]))
		third := e.letters.GetLetter(e.text[index + 2])
		firstMultiply := hyperdimensional.Multiply(first, second)
		secondMultiply := hyperdimensional.Multiply(firstMultiply, third)

		e.mutex.Lock()

		if e.lang == nil {
			e.lang = secondMultiply
		} else {
			e.lang.Add(secondMultiply)
		}

		e.mutex.Unlock()
	}
}
