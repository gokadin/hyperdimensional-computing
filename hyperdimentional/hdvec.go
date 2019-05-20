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

func Rotate(h *HdVec) *HdVec {
	var result [SIZE]int

	for index := range result {
		if index == SIZE - 1 {
			result[index] = h.vector[0]
			break
		}

		result[index] = h.vector[index + 1]
	}

	return &HdVec{
		vector: result,
	}
}

func Multiply(a *HdVec, b *HdVec) *HdVec {
	var result [SIZE]int

	for index := range result {
		result[index] = a.vector[index] * b.vector[index]
	}

    return &HdVec{
    	vector: result,
	}
}

func (h *HdVec) Print() {
	for _, value := range h.vector {
		fmt.Println(value)
	}
}
