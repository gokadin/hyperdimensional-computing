package hyperdimentional

import (
	"fmt"
	"math/rand"
)

const SIZE int = 10000

type HdVec struct {
	vector [SIZE]int
}

func NewHdVec() *HdVec {
    hdVec := &HdVec{}

    for index := range hdVec.vector {
		random := rand.Intn(2)
		if random == 0 {
			random = -1
		}
		hdVec.vector[index] = random
	}

    return hdVec
}

func (h *HdVec) Print() {
	for _, value := range h.vector {
		fmt.Println(value)
	}
}
