package hyperdimensional

import (
	"math"
	"math/rand"
)

const vectorSize = 10000

type VecBinomial struct {
    values []float32
}

func NewRandBinomial() *VecBinomial {
	vec := NewEmptyBinomial()

    for index := range vec.values {
		random := rand.Intn(2)
		if random == 0 {
            vec.values[index] = -1.0
		} else {
			vec.values[index] = 1.0
		}
	}

    return vec
}

func NewEmptyBinomial() *VecBinomial {
	return &VecBinomial{
		values: make([]float32, vectorSize),
	}
}

func Rotate(v *VecBinomial, count int) *VecBinomial {
	result := NewEmptyBinomial()
	for i := 0; i < result.Size(); i++ {
		if i >= result.Size() - count {
			result.values[i] = v.values[count - (result.Size() - i)]
			continue
		}

		result.values[i] = v.values[i + count]
	}

	return result
}

func Multiply(v1, v2 *VecBinomial) *VecBinomial {
	result := NewEmptyBinomial()
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

func (v *VecBinomial) ToBinomial() {
	for i := 0; i < v.Size(); i++ {
		if v.values[i] > 0 {
			v.values[i] = 1
		} else {
			v.values[i] = -1
		}
	}
}

func Cosine(v1, v2 *VecBinomial) float32 {
    dot := Dot(v1, v2)

    magnitudeProduct := v1.Magnitude() * v2.Magnitude()

    return dot / magnitudeProduct
}

func Dot(v1, v2 *VecBinomial) float32 {
	var result float32
    for i := 0; i < v1.Size(); i++ {
        result += v1.values[i] * v2.values[i]
	}

	return result
}

func (v *VecBinomial) Magnitude() float32 {
	var result float32
	for _, value := range v.values {
        result += value * value
	}

	return float32(math.Sqrt(float64(result)))
}

func (v *VecBinomial) Values() []float32 {
	return v.values
}

func (v *VecBinomial) Size() int {
	return len(v.values)
}

func (v *VecBinomial) At(index int) float32 {
	return v.values[index]
}

func (v *VecBinomial) Set(index int, value float32) {
	v.values[index] = value
}

func (v *VecBinomial) ScaleUp(size int) {
	scaled := make([]float32, size)
	scaleFactor := size / len(v.values)

	for i, value := range v.values {
		for j := 0; j < scaleFactor; j++ {
			scaled[i * scaleFactor + j] = value
		}
	}

	v.values = scaled
}
