package text

import (
	"bufio"
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"os"
	"strconv"
)

func VecBinomialFromFile(filename string) *hyperdimensional.VecBinomial {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("could not open file " + filename, err.Error())
	}
	defer file.Close()

	vec := hyperdimensional.NewRandBinomial()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		parsed, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			panic(err)
		}
		vec.Set(i, float32(parsed))
		i++
	}

	return vec
}

func writeToCache(filename string, vec *hyperdimensional.VecBinomial) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, value := range vec.Values() {
		_, _ = w.WriteString(strconv.FormatFloat(float64(value), 'f', 0, 32) + "\n")
	}
	w.Flush()
}
