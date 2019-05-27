package hyperdimentional

import (
	"fmt"
	"math"
	"math/rand"
)

type HdVecBinomial struct {
	vector []int
}

func NewHdVecBinomial(size int) *HdVecBinomial {
    hdVec := &HdVecBinomial{
    	vector: make([]int, size),
	}

    for index := range hdVec.vector {
		random := rand.Intn(2)
		if random == 0 {
			random = -1
		}
		hdVec.vector[index] = random
	}

    return hdVec
}

func Rotate(h *HdVecBinomial) *HdVecBinomial {
	result := make([]int, h.Size())

	for index := range result {
		if index == h.Size() - 1 {
			result[index] = h.vector[0]
			break
		}

		result[index] = h.vector[index + 1]
	}

	return &HdVecBinomial{
		vector: result,
	}
}

func Multiply(a *HdVecBinomial, b *HdVecBinomial) *HdVecBinomial {
	result := make([]int, a.Size())

	for index := range result {
		result[index] = a.vector[index] * b.vector[index]
	}

    return &HdVecBinomial{
    	vector: result,
	}
}

func Add(a *HdVecBinomial, b *HdVecBinomial) *HdVecBinomial {
	result := make([]int, a.Size())

	for index := range result {
		result[index] = a.vector[index] + b.vector[index]
	}

	return &HdVecBinomial{
		vector: result,
	}
}

func CosineSimilarity(a *HdVecBinomial, b *HdVecBinomial) float64 {
    dot := DotProduct(a, b)

    magnitudeProduct := a.Magnitude() * b.Magnitude()

    return float64(dot) / magnitudeProduct
}

func DotProduct(a *HdVecBinomial, b *HdVecBinomial) int {
	var result int
	for index := range a.vector {
        result += a.vector[index] * b.vector[index]
	}

	return result
}

func (h *HdVecBinomial) Magnitude() float64 {
	var result float64
	for value := range h.vector {
        result += math.Pow(float64(value), 2.0)
	}

	return math.Sqrt(result)
}

func (h *HdVecBinomial) Print() {
	for _, value := range h.vector {
		fmt.Println(value)
	}
}

func (h *HdVecBinomial) Size() int {
	return len(h.vector)
}

func (h *HdVecBinomial) First() int {
	return h.vector[0]
}

func (h *HdVecBinomial) Last() int {
	return h.vector[len(h.vector) - 1]
}

func (h *HdVecBinomial) Get(index int) int {
    return h.vector[index]
}

func (h *HdVecBinomial) Values() *[]int {
	return &h.vector
}
