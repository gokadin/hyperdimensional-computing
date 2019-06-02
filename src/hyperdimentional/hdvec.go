package hyperdimentional

import (
	"fmt"
	"math"
	"math/rand"
)

type VecBinomial struct {
    values []float64
}

func NewVecBinomial(size int) *VecBinomial {
	vec := &VecBinomial{
		values: make([]float64, size),
	}

    for index := range vec.values {
		random := rand.Intn(100)
		if random >= 50 {
            vec.values[index] = 1.0
		} else {
			vec.values[index] = -1.0
		}
	}

    return vec
}

func newEmpty(size int) *VecBinomial {
	return &VecBinomial{
		values: make([]float64, size),
	}
}

func Rotate(v *VecBinomial) *VecBinomial {
	result := newEmpty(v.Size())

	for i := 0; i < result.Size(); i++ {
		if i == result.Size() - 1 {
			result.values[i] = v.values[0]
			break
		}

		result.values[i] = v.values[i + 1]
	}

	return result
}

func Multiply(v1, v2 *VecBinomial) *VecBinomial {
	result := newEmpty(v1.Size())

    for i := 0; i < result.Size(); i++ {
		result.values[i] = v1.values[i] * v2.values[i]
	}

    return result
}

func (v *VecBinomial) Add(v2 *VecBinomial) {
    for i := 0; i < v.Size(); i++ {
		v.values[i] += v2.values[i]
	}
}

func Cosine(v1, v2 *VecBinomial) float64 {
    dot := Dot(v1, v2)

    magnitudeProduct := v1.Magnitude() * v2.Magnitude()

    return dot / magnitudeProduct
}

func Dot(v1, v2 *VecBinomial) float64 {
	result := 0.0
    for i := 0; i < v1.Size(); i++ {
        result += v1.values[i] * v2.values[i]
	}

	return result
}

func (v *VecBinomial) Magnitude() float64 {
	result := 0.0

	for _, value := range v.values {
        result += value * value
	}

	return math.Sqrt(result)
}

func (v *VecBinomial) Print() {
	for _, value := range v.values {
		fmt.Println(value)
	}
}

func (v *VecBinomial) Values() *[]float64 {
	return &v.values
}

func (v *VecBinomial) Size() int {
	return len(v.values)
}

func (v *VecBinomial) At(index int) float64 {
	return v.values[index]
}
