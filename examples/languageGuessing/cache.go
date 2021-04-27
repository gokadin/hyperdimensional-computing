package languageGuessing

import (
	"encoding/json"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"io/ioutil"
)

func VecFromFile(filename string) *hyperdimensional.HdVec {
	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var vec *hyperdimensional.HdVec
	err = json.Unmarshal(jsonBytes, &vec)
	if err != nil {
		panic(err)
	}

	return vec
}

func writeToCache(filename string, vec *hyperdimensional.HdVec) {
	vecJson, _ := json.Marshal(vec)
	_ = ioutil.WriteFile(filename, vecJson, 0666)
}
