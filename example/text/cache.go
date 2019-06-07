package text

import (
	"bufio"
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/src/hyperdimensional"
	"os"
	"strconv"
)

func VecBinomialFromFile(filename string) *hyperdimensional.VecBinomial {
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
